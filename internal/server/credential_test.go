package server

import (
	"errors"
	"sync"
	"testing"
	"time"
)

func TestCredentialUseExpiryAndExhaustion(t *testing.T) {
	now := time.Date(2026, 4, 28, 12, 0, 0, 0, time.UTC)
	store := newCredentialStore(2, time.Minute)
	token, c, err := store.Mint(now)
	if err != nil {
		t.Fatal(err)
	}
	if c.TriesLeft != 2 {
		t.Fatalf("tries = %d", c.TriesLeft)
	}
	if left, err := store.Use(token, now); err != nil || left != 1 {
		t.Fatalf("first use left=%d err=%v", left, err)
	}
	if left, err := store.Use(token, now); err != nil || left != 0 {
		t.Fatalf("second use left=%d err=%v", left, err)
	}
	if _, err := store.Use(token, now); !errors.Is(err, errCredentialExhausted) {
		t.Fatalf("expected exhausted, got %v", err)
	}
	expired, _, err := store.Mint(now)
	if err != nil {
		t.Fatal(err)
	}
	if _, err := store.Use(expired, now.Add(2*time.Minute)); !errors.Is(err, errCredentialInvalid) {
		t.Fatalf("expected invalid expired credential, got %v", err)
	}
}

func TestCredentialConcurrentUse(t *testing.T) {
	now := time.Now()
	store := newCredentialStore(100, time.Minute)
	token, _, err := store.Mint(now)
	if err != nil {
		t.Fatal(err)
	}
	var wg sync.WaitGroup
	for i := 0; i < 20; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			_, _ = store.Use(token, now)
		}()
	}
	wg.Wait()
	left, err := store.Use(token, now)
	if err != nil {
		t.Fatal(err)
	}
	if left != 79 {
		t.Fatalf("left = %d, want 79", left)
	}
}
