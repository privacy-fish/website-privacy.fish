package server

import (
	"crypto/rand"
	"encoding/hex"
	"errors"
	"sync"
	"time"
)

var (
	errCredentialInvalid   = errors.New("credential_invalid")
	errCredentialExhausted = errors.New("credential_exhausted")
)

type credential struct {
	TriesLeft int
	ExpiresAt time.Time
}

type credentialStore struct {
	mu      sync.Mutex
	tries   int
	ttl     time.Duration
	entries map[string]credential
}

func newCredentialStore(tries int, ttl time.Duration) *credentialStore {
	return &credentialStore{tries: tries, ttl: ttl, entries: make(map[string]credential)}
}

func (s *credentialStore) Mint(now time.Time) (string, credential, error) {
	var b [32]byte
	if _, err := rand.Read(b[:]); err != nil {
		return "", credential{}, err
	}
	token := hex.EncodeToString(b[:])
	c := credential{TriesLeft: s.tries, ExpiresAt: now.Add(s.ttl)}
	s.mu.Lock()
	s.entries[token] = c
	s.mu.Unlock()
	return token, c, nil
}

func (s *credentialStore) Use(token string, now time.Time) (int, error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	c, ok := s.entries[token]
	if !ok || now.After(c.ExpiresAt) {
		delete(s.entries, token)
		return 0, errCredentialInvalid
	}
	if c.TriesLeft <= 0 {
		return 0, errCredentialExhausted
	}
	c.TriesLeft--
	s.entries[token] = c
	return c.TriesLeft, nil
}

func (s *credentialStore) Sweep(now time.Time) int {
	s.mu.Lock()
	defer s.mu.Unlock()
	n := 0
	for token, c := range s.entries {
		if now.After(c.ExpiresAt) {
			delete(s.entries, token)
			n++
		}
	}
	return n
}
