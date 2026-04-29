package server

import (
	"context"
	"crypto/ed25519"
	"crypto/rand"
	"encoding/json"
	"io"
	"log/slog"
	"net/http"
	"net/http/httptest"
	"net/url"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"testing"
	"time"

	"filippo.io/age"
	"golang.org/x/crypto/ssh"
)

type mockCaptcha struct {
	err error
}

func (m mockCaptcha) Verify(context.Context, string) error {
	return m.err
}

func TestSignupFlow(t *testing.T) {
	svc := testServer(t, nil)

	lookup := postForm(t, svc, "/api/lookup-session", url.Values{"cap-token": {"ok"}})
	if lookup.Code != http.StatusOK {
		t.Fatalf("lookup status = %d body=%s", lookup.Code, lookup.Body.String())
	}
	var lookupBody struct {
		Credential string `json:"credential"`
		TriesLeft  int    `json:"tries_left"`
	}
	decodeBody(t, lookup, &lookupBody)
	if lookupBody.Credential == "" || lookupBody.TriesLeft != 5 {
		t.Fatalf("lookup body = %+v", lookupBody)
	}

	check := postForm(t, svc, "/api/check-username", url.Values{
		"credential": {lookupBody.Credential},
		"username":   {"Alice"},
	})
	if check.Code != http.StatusOK {
		t.Fatalf("check status = %d body=%s", check.Code, check.Body.String())
	}
	var checkBody struct {
		Available bool `json:"available"`
		TriesLeft int  `json:"tries_left"`
	}
	decodeBody(t, check, &checkBody)
	if !checkBody.Available || checkBody.TriesLeft != 4 {
		t.Fatalf("check body = %+v", checkBody)
	}

	signup := postForm(t, svc, "/api/signup", url.Values{
		"credential": {lookupBody.Credential},
		"username":   {"alice"},
		"ssh_keys":   {testAuthorizedKey(t)},
	})
	if signup.Code != http.StatusOK {
		t.Fatalf("signup status = %d body=%s", signup.Code, signup.Body.String())
	}
	var signupBody struct {
		OK               bool   `json:"ok"`
		PaymentReference string `json:"payment_reference"`
		TriesLeft        int    `json:"tries_left"`
	}
	decodeBody(t, signup, &signupBody)
	if !signupBody.OK || signupBody.PaymentReference == "" || signupBody.TriesLeft != 3 {
		t.Fatalf("signup body = %+v", signupBody)
	}
	items, err := svc.pending.List()
	if err != nil {
		t.Fatal(err)
	}
	if len(items) != 1 {
		t.Fatalf("pending count = %d, want 1", len(items))
	}
}

func TestSignupRejectsBadOriginAndTakenUsername(t *testing.T) {
	svc := testServer(t, []string{"taken"})
	badOriginReq := httptest.NewRequest(http.MethodPost, "/api/lookup-session", strings.NewReader("cap-token=ok"))
	badOriginReq.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	badOriginReq.Header.Set("Origin", "https://evil.example")
	rec := httptest.NewRecorder()
	svc.ServeHTTP(rec, badOriginReq)
	if rec.Code != http.StatusForbidden {
		t.Fatalf("bad origin status = %d", rec.Code)
	}

	token, _, err := svc.credentials.Mint(time.Now())
	if err != nil {
		t.Fatal(err)
	}
	res := postForm(t, svc, "/api/signup", url.Values{
		"credential": {token},
		"username":   {"taken"},
		"ssh_keys":   {testAuthorizedKey(t)},
	})
	if res.Code != http.StatusConflict {
		t.Fatalf("taken status = %d body=%s", res.Code, res.Body.String())
	}
	var body struct {
		Code string `json:"code"`
	}
	decodeBody(t, res, &body)
	if body.Code != "username_taken" {
		t.Fatalf("code = %s", body.Code)
	}
}

func TestConcurrentSignupSameUsernameCreatesOnePending(t *testing.T) {
	svc := testServer(t, nil)
	tokenA, _, err := svc.credentials.Mint(time.Now())
	if err != nil {
		t.Fatal(err)
	}
	tokenB, _, err := svc.credentials.Mint(time.Now())
	if err != nil {
		t.Fatal(err)
	}
	key := testAuthorizedKey(t)

	var wg sync.WaitGroup
	results := make([]int, 2)
	for i, token := range []string{tokenA, tokenB} {
		wg.Add(1)
		go func(i int, token string) {
			defer wg.Done()
			res := postForm(t, svc, "/api/signup", url.Values{
				"credential": {token},
				"username":   {"race"},
				"ssh_keys":   {key},
			})
			results[i] = res.Code
		}(i, token)
	}
	wg.Wait()

	var ok, conflict int
	for _, code := range results {
		if code == http.StatusOK {
			ok++
		}
		if code == http.StatusConflict {
			conflict++
		}
	}
	if ok != 1 || conflict != 1 {
		t.Fatalf("status counts ok=%d conflict=%d results=%v", ok, conflict, results)
	}
	items, err := svc.pending.List()
	if err != nil {
		t.Fatal(err)
	}
	if len(items) != 1 {
		t.Fatalf("pending count = %d, want 1", len(items))
	}
}

func testServer(t *testing.T, usedUsernames []string) *Server {
	t.Helper()
	dir := t.TempDir()
	identity, err := age.GenerateX25519Identity()
	if err != nil {
		t.Fatal(err)
	}
	pepper := []byte("0123456789abcdef0123456789abcdef")
	salt := []byte("0123456789abcdef")
	entries := make([]byte, 8*32)
	if _, err := rand.Read(entries); err != nil {
		t.Fatal(err)
	}
	for i, username := range usedUsernames {
		copy(entries[i*32:(i+1)*32], UsernameHMAC(pepper, salt, username))
	}
	if err := os.WriteFile(filepath.Join(dir, "used.bin"), entries, 0o600); err != nil {
		t.Fatal(err)
	}
	svc, err := New(Config{
		DataDir:           dir,
		SiteOrigin:        "https://privacy.fish",
		AdminAgePublicKey: identity.Recipient().String(),
		UsernamePepper:    pepper,
		UsernameSalt:      salt,
		UsedStoreSize:     8,
		PendingTTL:        time.Hour,
		CredentialTries:   5,
		CredentialTTL:     time.Minute,
	}, mockCaptcha{}, slog.New(slog.NewTextHandler(io.Discard, nil)))
	if err != nil {
		t.Fatal(err)
	}
	return svc
}

func postForm(t *testing.T, handler http.Handler, path string, values url.Values) *httptest.ResponseRecorder {
	t.Helper()
	req := httptest.NewRequest(http.MethodPost, path, strings.NewReader(values.Encode()))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("Origin", "https://privacy.fish")
	rec := httptest.NewRecorder()
	handler.ServeHTTP(rec, req)
	return rec
}

func decodeBody(t *testing.T, rec *httptest.ResponseRecorder, target any) {
	t.Helper()
	if err := json.NewDecoder(rec.Body).Decode(target); err != nil {
		t.Fatal(err)
	}
}

func testAuthorizedKey(t *testing.T) string {
	t.Helper()
	pub, _, err := ed25519.GenerateKey(rand.Reader)
	if err != nil {
		t.Fatal(err)
	}
	sshPub, err := ssh.NewPublicKey(pub)
	if err != nil {
		t.Fatal(err)
	}
	return strings.TrimSpace(string(ssh.MarshalAuthorizedKey(sshPub)))
}
