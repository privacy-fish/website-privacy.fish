package store

import (
	"bytes"
	"crypto/rand"
	"crypto/subtle"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"sort"
	"time"

	"filippo.io/age"
)

type Pending struct {
	Dir       string
	TTL       time.Duration
	Recipient age.Recipient
}

type PendingRecord struct {
	UsernameHMAC     string    `json:"username_hmac"`
	CreatedAt        time.Time `json:"created_at"`
	ExpiresAt        time.Time `json:"expires_at"`
	CiphertextAgeB64 string    `json:"ciphertext_age_b64"`
}

type SignupPayload struct {
	Username         string   `json:"username"`
	Keys             []string `json:"keys"`
	PaymentReference string   `json:"payment_reference"`
	CreatedAt        string   `json:"created_at"`
}

var pendingIDPattern = regexp.MustCompile(`^[a-f0-9]{64}$`)

func NewPending(dir string, ttl time.Duration, recipient age.Recipient) *Pending {
	return &Pending{Dir: dir, TTL: ttl, Recipient: recipient}
}

func (p *Pending) Ensure() error {
	if err := os.MkdirAll(p.Dir, 0o700); err != nil {
		return err
	}
	return os.Chmod(p.Dir, 0o700)
}

func (p *Pending) Count() (int, error) {
	records, err := p.List()
	return len(records), err
}

func (p *Pending) Write(usernameHMAC []byte, payload SignupPayload, now time.Time) (string, time.Time, error) {
	if len(usernameHMAC) != HMACSize {
		return "", time.Time{}, errors.New("username hmac must be 32 bytes")
	}
	if err := p.Ensure(); err != nil {
		return "", time.Time{}, err
	}
	plain, err := json.Marshal(payload)
	if err != nil {
		return "", time.Time{}, err
	}
	var ciphertext bytes.Buffer
	w, err := age.Encrypt(&ciphertext, p.Recipient)
	if err != nil {
		return "", time.Time{}, err
	}
	if _, err := w.Write(plain); err != nil {
		return "", time.Time{}, err
	}
	if err := w.Close(); err != nil {
		return "", time.Time{}, err
	}

	id, err := randomHex(32)
	if err != nil {
		return "", time.Time{}, err
	}
	expiresAt := now.Add(p.TTL)
	record := PendingRecord{
		UsernameHMAC:     hex.EncodeToString(usernameHMAC),
		CreatedAt:        now.UTC(),
		ExpiresAt:        expiresAt.UTC(),
		CiphertextAgeB64: base64.StdEncoding.EncodeToString(ciphertext.Bytes()),
	}
	data, err := json.MarshalIndent(record, "", "  ")
	if err != nil {
		return "", time.Time{}, err
	}
	data = append(data, '\n')
	tmp := filepath.Join(p.Dir, id+".tmp")
	dst := filepath.Join(p.Dir, id+".json")
	if err := os.WriteFile(tmp, data, 0o600); err != nil {
		return "", time.Time{}, err
	}
	if err := os.Rename(tmp, dst); err != nil {
		_ = os.Remove(tmp)
		return "", time.Time{}, err
	}
	return id, expiresAt, nil
}

func (p *Pending) ScanHMAC(target []byte) (bool, error) {
	if len(target) != HMACSize {
		return false, errors.New("target hmac must be 32 bytes")
	}
	records, err := p.List()
	if err != nil {
		return false, err
	}
	var found int
	for _, item := range records {
		h, err := hex.DecodeString(item.Record.UsernameHMAC)
		if err != nil || len(h) != HMACSize {
			return false, fmt.Errorf("invalid pending hmac in %s", item.ID)
		}
		found |= subtle.ConstantTimeCompare(h, target)
	}
	return found == 1, nil
}

type PendingItem struct {
	ID     string
	Record PendingRecord
}

func (p *Pending) List() ([]PendingItem, error) {
	if err := p.Ensure(); err != nil {
		return nil, err
	}
	entries, err := os.ReadDir(p.Dir)
	if err != nil {
		return nil, err
	}
	out := make([]PendingItem, 0, len(entries))
	for _, entry := range entries {
		if entry.IsDir() || filepath.Ext(entry.Name()) != ".json" {
			continue
		}
		id := entry.Name()[:len(entry.Name())-len(".json")]
		if !pendingIDPattern.MatchString(id) {
			continue
		}
		record, err := p.Read(id)
		if err != nil {
			return nil, err
		}
		out = append(out, PendingItem{ID: id, Record: record})
	}
	sort.Slice(out, func(i, j int) bool { return out[i].ID < out[j].ID })
	return out, nil
}

func (p *Pending) Read(id string) (PendingRecord, error) {
	if !pendingIDPattern.MatchString(id) {
		return PendingRecord{}, errors.New("invalid pending id")
	}
	data, err := os.ReadFile(filepath.Join(p.Dir, id+".json"))
	if err != nil {
		return PendingRecord{}, err
	}
	var record PendingRecord
	if err := json.Unmarshal(data, &record); err != nil {
		return PendingRecord{}, err
	}
	return record, nil
}

func (p *Pending) Delete(id string) error {
	if !pendingIDPattern.MatchString(id) {
		return errors.New("invalid pending id")
	}
	return os.Remove(filepath.Join(p.Dir, id+".json"))
}

func (p *Pending) Sweep(now time.Time) (int, error) {
	items, err := p.List()
	if err != nil {
		return 0, err
	}
	n := 0
	for _, item := range items {
		if item.Record.ExpiresAt.Before(now) {
			if err := p.Delete(item.ID); err != nil {
				return n, err
			}
			n++
		}
	}
	return n, nil
}

func randomHex(n int) (string, error) {
	b := make([]byte, n)
	if _, err := rand.Read(b); err != nil {
		return "", err
	}
	return hex.EncodeToString(b), nil
}
