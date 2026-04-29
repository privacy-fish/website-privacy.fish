package captcha

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strings"
	"time"
)

var ErrFailed = errors.New("captcha_failed")

type Verifier interface {
	Verify(ctx context.Context, token string) error
}

type Cap struct {
	BaseURL string
	SiteKey string
	Secret  string
	Client  *http.Client
}

func (c Cap) Verify(ctx context.Context, token string) error {
	if strings.TrimSpace(token) == "" {
		return ErrFailed
	}
	body, err := json.Marshal(map[string]string{
		"secret":   c.Secret,
		"response": token,
	})
	if err != nil {
		return err
	}
	base := strings.TrimRight(c.BaseURL, "/")
	url := fmt.Sprintf("%s/%s/siteverify", base, c.SiteKey)
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, url, bytes.NewReader(body))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")
	client := c.Client
	if client == nil {
		client = &http.Client{Timeout: 5 * time.Second}
	}
	res, err := client.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()
	var parsed struct {
		Success bool `json:"success"`
	}
	if err := json.NewDecoder(res.Body).Decode(&parsed); err != nil {
		return err
	}
	if res.StatusCode < 200 || res.StatusCode > 299 || !parsed.Success {
		return ErrFailed
	}
	return nil
}
