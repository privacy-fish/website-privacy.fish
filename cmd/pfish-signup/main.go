package main

import (
	"encoding/hex"
	"errors"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/privacyfish/pfish-signup/internal/admin"
	"github.com/privacyfish/pfish-signup/internal/captcha"
	"github.com/privacyfish/pfish-signup/internal/server"
)

func main() {
	logger := slog.New(slog.NewJSONHandler(os.Stderr, nil))
	if len(os.Args) > 1 && os.Args[1] == "admin" {
		cfg, err := loadAdminConfig()
		if err != nil {
			logger.Error("config_error", "error", err)
			os.Exit(1)
		}
		if err := admin.Run(os.Args[2:], cfg.DataDir, cfg.UsedStoreSize, os.Stdout); err != nil {
			logger.Error("admin_error", "error", err)
			os.Exit(1)
		}
		return
	}

	listenAddr := envDefault("LISTEN_ADDR", "127.0.0.1:8080")
	cfg, capCfg, err := loadServerConfig()
	if err != nil {
		logger.Error("config_error", "error", err)
		os.Exit(1)
	}
	svc, err := server.New(cfg, captcha.Cap{
		BaseURL: capCfg.baseURL,
		SiteKey: capCfg.siteKey,
		Secret:  capCfg.secret,
	}, logger)
	if err != nil {
		logger.Error("startup_error", "error", err)
		os.Exit(1)
	}
	stop := make(chan struct{})
	defer close(stop)
	svc.StartSweepers(stop)
	logger.Info("startup", "listen_addr", listenAddr)
	httpServer := &http.Server{
		Addr:              listenAddr,
		Handler:           svc,
		ReadHeaderTimeout: 5 * time.Second,
		ReadTimeout:       10 * time.Second,
		WriteTimeout:      30 * time.Second,
		IdleTimeout:       60 * time.Second,
	}
	if err := httpServer.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
		logger.Error("server_error", "error", err)
		os.Exit(1)
	}
}

type capConfig struct {
	baseURL string
	siteKey string
	secret  string
}

func loadServerConfig() (server.Config, capConfig, error) {
	adminCfg, err := loadAdminConfig()
	if err != nil {
		return server.Config{}, capConfig{}, err
	}
	siteOrigin, err := required("SITE_ORIGIN")
	if err != nil {
		return server.Config{}, capConfig{}, err
	}
	capSiteKey, err := required("CAP_SITE_KEY")
	if err != nil {
		return server.Config{}, capConfig{}, err
	}
	capSecret, err := required("CAP_SECRET")
	if err != nil {
		return server.Config{}, capConfig{}, err
	}
	agePublic, err := required("ADMIN_AGE_PUBLIC_KEY")
	if err != nil {
		return server.Config{}, capConfig{}, err
	}
	pepper, err := requiredHex("USERNAME_HMAC_PEPPER", 32)
	if err != nil {
		return server.Config{}, capConfig{}, err
	}
	salt, err := requiredHex("USERNAME_HMAC_SALT", 16)
	if err != nil {
		return server.Config{}, capConfig{}, err
	}
	pendingTTL, err := durationEnv("PENDING_TTL", 720*time.Hour)
	if err != nil {
		return server.Config{}, capConfig{}, err
	}
	if pendingTTL <= 0 {
		return server.Config{}, capConfig{}, fmt.Errorf("PENDING_TTL must be positive")
	}
	credentialTTL, err := durationEnv("CREDENTIAL_TTL", 10*time.Minute)
	if err != nil {
		return server.Config{}, capConfig{}, err
	}
	if credentialTTL <= 0 {
		return server.Config{}, capConfig{}, fmt.Errorf("CREDENTIAL_TTL must be positive")
	}
	credentialTries, err := intEnv("CREDENTIAL_TRIES", 5)
	if err != nil {
		return server.Config{}, capConfig{}, err
	}
	if credentialTries <= 0 {
		return server.Config{}, capConfig{}, fmt.Errorf("CREDENTIAL_TRIES must be positive")
	}
	return server.Config{
			DataDir:           adminCfg.DataDir,
			SiteOrigin:        siteOrigin,
			AdminAgePublicKey: agePublic,
			UsernamePepper:    pepper,
			UsernameSalt:      salt,
			UsedStoreSize:     adminCfg.UsedStoreSize,
			PendingTTL:        pendingTTL,
			CredentialTries:   credentialTries,
			CredentialTTL:     credentialTTL,
		}, capConfig{
			baseURL: envDefault("CAP_API_BASE", "http://127.0.0.1:3000"),
			siteKey: capSiteKey,
			secret:  capSecret,
		}, nil
}

type adminConfig struct {
	DataDir       string
	UsedStoreSize int64
}

func loadAdminConfig() (adminConfig, error) {
	size, err := intEnv("USED_STORE_SIZE", 1000000)
	if err != nil {
		return adminConfig{}, err
	}
	if size <= 0 {
		return adminConfig{}, fmt.Errorf("USED_STORE_SIZE must be positive")
	}
	return adminConfig{
		DataDir:       envDefault("DATA_DIR", "/var/lib/pfish-signup"),
		UsedStoreSize: int64(size),
	}, nil
}

func required(name string) (string, error) {
	value := os.Getenv(name)
	if value == "" {
		return "", fmt.Errorf("%s is required", name)
	}
	return value, nil
}

func requiredHex(name string, minBytes int) ([]byte, error) {
	value, err := required(name)
	if err != nil {
		return nil, err
	}
	decoded, err := hex.DecodeString(value)
	if err != nil {
		return nil, fmt.Errorf("%s must be hex", name)
	}
	if len(decoded) < minBytes {
		return nil, fmt.Errorf("%s must decode to at least %d bytes", name, minBytes)
	}
	return decoded, nil
}

func envDefault(name, fallback string) string {
	if value := os.Getenv(name); value != "" {
		return value
	}
	return fallback
}

func intEnv(name string, fallback int) (int, error) {
	value := os.Getenv(name)
	if value == "" {
		return fallback, nil
	}
	parsed, err := strconv.Atoi(value)
	if err != nil {
		return 0, fmt.Errorf("%s must be an integer", name)
	}
	return parsed, nil
}

func durationEnv(name string, fallback time.Duration) (time.Duration, error) {
	value := os.Getenv(name)
	if value == "" {
		return fallback, nil
	}
	parsed, err := time.ParseDuration(value)
	if err != nil {
		return 0, fmt.Errorf("%s must be a duration", name)
	}
	return parsed, nil
}
