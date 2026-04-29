package server

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/json"
	"errors"
	"log/slog"
	"mime"
	"net"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"sync"
	"time"

	"filippo.io/age"
	"github.com/privacyfish/pfish-signup/internal/captcha"
	"github.com/privacyfish/pfish-signup/internal/limit"
	"github.com/privacyfish/pfish-signup/internal/refgen"
	"github.com/privacyfish/pfish-signup/internal/sshkey"
	"github.com/privacyfish/pfish-signup/internal/store"
)

const maxBodyBytes = 16 << 10

var usernamePattern = regexp.MustCompile(`^[a-z0-9](?:[a-z0-9.-]{0,29}[a-z0-9])?$`)

type Config struct {
	DataDir           string
	SiteOrigin        string
	AdminAgePublicKey string
	UsernamePepper    []byte
	UsernameSalt      []byte
	UsedStoreSize     int64
	PendingTTL        time.Duration
	CredentialTries   int
	CredentialTTL     time.Duration
}

type Server struct {
	cfg         Config
	log         *slog.Logger
	captcha     captcha.Verifier
	credentials *credentialStore
	attempts    *limit.Limiter
	signups     *limit.Limiter
	pending     *store.Pending
	used        *store.Used
	reserveMu   sync.Mutex
}

func New(cfg Config, verifier captcha.Verifier, logger *slog.Logger) (*Server, error) {
	recipient, err := age.ParseX25519Recipient(cfg.AdminAgePublicKey)
	if err != nil {
		return nil, err
	}
	if logger == nil {
		logger = slog.Default()
	}
	if err := os.MkdirAll(cfg.DataDir, 0o700); err != nil {
		return nil, err
	}
	if err := os.Chmod(cfg.DataDir, 0o700); err != nil {
		return nil, err
	}
	used := store.NewUsed(filepath.Join(cfg.DataDir, "used.bin"), cfg.UsedStoreSize)
	if err := used.Validate(); err != nil {
		return nil, err
	}
	pending := store.NewPending(filepath.Join(cfg.DataDir, "pending"), cfg.PendingTTL, recipient)
	if err := pending.Ensure(); err != nil {
		return nil, err
	}
	return &Server{
		cfg:         cfg,
		log:         logger,
		captcha:     verifier,
		credentials: newCredentialStore(cfg.CredentialTries, cfg.CredentialTTL),
		attempts:    limit.New(30, time.Hour),
		signups:     limit.New(5, time.Hour),
		pending:     pending,
		used:        used,
	}, nil
}

func (s *Server) StartSweepers(stop <-chan struct{}) {
	go func() {
		credTicker := time.NewTicker(time.Minute)
		pendingTicker := time.NewTicker(time.Hour)
		defer credTicker.Stop()
		defer pendingTicker.Stop()
		s.sweep(time.Now())
		for {
			select {
			case now := <-credTicker.C:
				n := s.credentials.Sweep(now)
				if n > 0 {
					s.log.Info("credential_sweep", "deleted", n)
				}
			case now := <-pendingTicker.C:
				s.sweep(now)
			case <-stop:
				return
			}
		}
	}()
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/healthz":
		s.handleHealthz(w, r)
	case "/api/lookup-session":
		s.handleLookupSession(w, r)
	case "/api/check-username":
		s.handleCheckUsername(w, r)
	case "/api/signup":
		s.handleSignup(w, r)
	default:
		http.NotFound(w, r)
	}
}

func (s *Server) handleLookupSession(w http.ResponseWriter, r *http.Request) {
	form, ok := s.parseForm(w, r)
	if !ok {
		return
	}
	now := time.Now()
	if !s.allowAttempt(w, r, now) {
		return
	}
	if err := s.captcha.Verify(r.Context(), form.Get("cap-token")); err != nil {
		s.fail(w, http.StatusBadRequest, "captcha_failed")
		return
	}
	token, c, err := s.credentials.Mint(now)
	if err != nil {
		s.fail(w, http.StatusInternalServerError, "server_error")
		return
	}
	writeJSON(w, http.StatusOK, map[string]any{
		"ok":         true,
		"credential": token,
		"tries_left": c.TriesLeft,
		"expires_at": c.ExpiresAt.UTC().Format(time.RFC3339),
	})
}

func (s *Server) handleCheckUsername(w http.ResponseWriter, r *http.Request) {
	form, ok := s.parseForm(w, r)
	if !ok {
		return
	}
	now := time.Now()
	if !s.allowAttempt(w, r, now) {
		return
	}
	triesLeft, err := s.credentials.Use(form.Get("credential"), now)
	if err != nil {
		s.credentialFailure(w, err)
		return
	}
	username, ok := normalizeUsername(form.Get("username"))
	if !ok {
		s.fail(w, http.StatusBadRequest, "username_invalid")
		return
	}
	taken, err := s.usernameTaken(username)
	if err != nil {
		s.fail(w, http.StatusInternalServerError, "server_error")
		return
	}
	writeJSON(w, http.StatusOK, map[string]any{
		"ok":         true,
		"available":  !taken,
		"tries_left": triesLeft,
	})
}

func (s *Server) handleSignup(w http.ResponseWriter, r *http.Request) {
	form, ok := s.parseForm(w, r)
	if !ok {
		return
	}
	now := time.Now()
	if !s.allowAttempt(w, r, now) {
		return
	}
	triesLeft, err := s.credentials.Use(form.Get("credential"), now)
	if err != nil {
		s.credentialFailure(w, err)
		return
	}
	username, ok := normalizeUsername(form.Get("username"))
	if !ok {
		s.fail(w, http.StatusBadRequest, "username_invalid")
		return
	}
	usernameHMAC := UsernameHMAC(s.cfg.UsernamePepper, s.cfg.UsernameSalt, username)
	keys, err := sshkey.ParseAuthorized(form.Get("ssh_keys"))
	if err != nil {
		s.keyFailure(w, err)
		return
	}
	reference, err := refgen.New()
	if err != nil {
		s.fail(w, http.StatusInternalServerError, "server_error")
		return
	}
	createdAt := now.UTC().Format(time.RFC3339)
	payload := store.SignupPayload{
		Username:         username,
		Keys:             keys,
		PaymentReference: reference,
		CreatedAt:        createdAt,
	}

	s.reserveMu.Lock()
	defer s.reserveMu.Unlock()

	count, err := s.pending.Count()
	if err != nil {
		s.fail(w, http.StatusInternalServerError, "server_error")
		return
	}
	if count >= 1000 {
		s.fail(w, http.StatusTooManyRequests, "rate_limited")
		return
	}
	taken, err := s.hmacTaken(usernameHMAC)
	if err != nil {
		s.fail(w, http.StatusInternalServerError, "server_error")
		return
	}
	if taken {
		s.fail(w, http.StatusConflict, "username_taken")
		return
	}
	if !s.allowSignup(w, r, now) {
		return
	}
	_, expiresAt, err := s.pending.Write(usernameHMAC, payload, now)
	if err != nil {
		s.fail(w, http.StatusInternalServerError, "server_error")
		return
	}
	writeJSON(w, http.StatusOK, map[string]any{
		"ok":                true,
		"payment_reference": reference,
		"expires_at":        expiresAt.UTC().Format(time.RFC3339),
		"tries_left":        triesLeft,
	})
}

func (s *Server) handleHealthz(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		s.fail(w, http.StatusMethodNotAllowed, "bad_method")
		return
	}
	if err := s.health(); err != nil {
		w.WriteHeader(http.StatusServiceUnavailable)
		_, _ = w.Write([]byte("not ready\n"))
		return
	}
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write([]byte("ok\n"))
}

func (s *Server) health() error {
	if err := s.used.Validate(); err != nil {
		return err
	}
	if _, err := age.ParseX25519Recipient(s.cfg.AdminAgePublicKey); err != nil {
		return err
	}
	tmp, err := os.CreateTemp(s.cfg.DataDir, ".healthz-*")
	if err != nil {
		return err
	}
	name := tmp.Name()
	if err := tmp.Close(); err != nil {
		return err
	}
	return os.Remove(name)
}

func (s *Server) parseForm(w http.ResponseWriter, r *http.Request) (url.Values, bool) {
	if r.Method != http.MethodPost {
		s.fail(w, http.StatusMethodNotAllowed, "bad_method")
		return nil, false
	}
	if !s.goodOrigin(r) {
		s.fail(w, http.StatusForbidden, "bad_origin")
		return nil, false
	}
	mediaType, _, err := mime.ParseMediaType(r.Header.Get("Content-Type"))
	if err != nil || mediaType != "application/x-www-form-urlencoded" {
		s.fail(w, http.StatusBadRequest, "bad_request")
		return nil, false
	}
	r.Body = http.MaxBytesReader(w, r.Body, maxBodyBytes)
	if err := r.ParseForm(); err != nil {
		if strings.Contains(err.Error(), "request body too large") {
			s.fail(w, http.StatusRequestEntityTooLarge, "body_too_large")
			return nil, false
		}
		s.fail(w, http.StatusBadRequest, "bad_request")
		return nil, false
	}
	return r.PostForm, true
}

func (s *Server) goodOrigin(r *http.Request) bool {
	if s.cfg.SiteOrigin == "" {
		return false
	}
	if origin := r.Header.Get("Origin"); origin != "" {
		return origin == s.cfg.SiteOrigin
	}
	referer := r.Header.Get("Referer")
	if referer == "" {
		return false
	}
	u, err := url.Parse(referer)
	if err != nil {
		return false
	}
	return u.Scheme+"://"+u.Host == s.cfg.SiteOrigin
}

func (s *Server) allowAttempt(w http.ResponseWriter, r *http.Request, now time.Time) bool {
	if s.attempts.Allow(sourceIP(r), now) {
		return true
	}
	s.log.Warn("rate_limited", "bucket", "api_attempts")
	s.fail(w, http.StatusTooManyRequests, "rate_limited")
	return false
}

func (s *Server) allowSignup(w http.ResponseWriter, r *http.Request, now time.Time) bool {
	if s.signups.Allow(sourceIP(r), now) {
		return true
	}
	s.log.Warn("rate_limited", "bucket", "successful_signups")
	s.fail(w, http.StatusTooManyRequests, "rate_limited")
	return false
}

func (s *Server) usernameTaken(username string) (bool, error) {
	return s.hmacTaken(UsernameHMAC(s.cfg.UsernamePepper, s.cfg.UsernameSalt, username))
}

func (s *Server) hmacTaken(h []byte) (bool, error) {
	usedTaken, _, err := s.used.Scan(h)
	if err != nil {
		return false, err
	}
	pendingTaken, err := s.pending.ScanHMAC(h)
	if err != nil {
		return false, err
	}
	return usedTaken || pendingTaken, nil
}

func UsernameHMAC(pepper, salt []byte, username string) []byte {
	mac := hmac.New(sha256.New, pepper)
	mac.Write(salt)
	mac.Write([]byte(username))
	return mac.Sum(nil)
}

func normalizeUsername(raw string) (string, bool) {
	username := strings.ToLower(strings.TrimSpace(raw))
	return username, usernamePattern.MatchString(username)
}

func sourceIP(r *http.Request) string {
	host, _, err := net.SplitHostPort(r.RemoteAddr)
	if err != nil {
		host = r.RemoteAddr
	}
	if ip := net.ParseIP(host); ip != nil && ip.IsLoopback() {
		if xff := r.Header.Get("X-Forwarded-For"); xff != "" {
			return strings.TrimSpace(strings.Split(xff, ",")[0])
		}
	}
	return host
}

func (s *Server) credentialFailure(w http.ResponseWriter, err error) {
	switch {
	case errors.Is(err, errCredentialExhausted):
		s.fail(w, http.StatusUnauthorized, "credential_exhausted")
	default:
		s.fail(w, http.StatusUnauthorized, "credential_invalid")
	}
}

func (s *Server) keyFailure(w http.ResponseWriter, err error) {
	switch {
	case errors.Is(err, sshkey.ErrMissing):
		s.fail(w, http.StatusBadRequest, "keys_missing")
	case errors.Is(err, sshkey.ErrTooMany):
		s.fail(w, http.StatusBadRequest, "keys_too_many")
	case errors.Is(err, sshkey.ErrWrongType):
		s.fail(w, http.StatusBadRequest, "key_wrong_type")
	default:
		s.fail(w, http.StatusBadRequest, "key_invalid")
	}
}

func (s *Server) fail(w http.ResponseWriter, status int, code string) {
	s.log.Warn("request_failed", "code", code)
	writeJSON(w, status, map[string]any{"ok": false, "code": code})
}

func (s *Server) sweep(now time.Time) {
	n, err := s.pending.Sweep(now)
	if err != nil {
		s.log.Error("pending_sweep_failed", "error", err)
		return
	}
	if n > 0 {
		s.log.Info("pending_sweep", "deleted", n)
	}
}

func writeJSON(w http.ResponseWriter, status int, v any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(v)
}
