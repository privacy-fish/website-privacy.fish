package store

import (
	"bytes"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"os"
	"path/filepath"
	"testing"
	"time"

	"filippo.io/age"
)

func TestPendingWriteReadScanAndSweep(t *testing.T) {
	identity, err := age.GenerateX25519Identity()
	if err != nil {
		t.Fatal(err)
	}
	dir := filepath.Join(t.TempDir(), "pending")
	pending := NewPending(dir, time.Hour, identity.Recipient())
	now := time.Date(2026, 4, 28, 12, 0, 0, 0, time.UTC)
	hmac := bytes.Repeat([]byte{7}, HMACSize)
	payload := SignupPayload{
		Username:         "alice",
		Keys:             []string{"ssh-ed25519 AAAA"},
		PaymentReference: "0000-0000-0000-0000-0000",
		CreatedAt:        now.Format(time.RFC3339),
	}
	id, expiresAt, err := pending.Write(hmac, payload, now)
	if err != nil {
		t.Fatal(err)
	}
	record, err := pending.Read(id)
	if err != nil {
		t.Fatal(err)
	}
	if record.UsernameHMAC != hex.EncodeToString(hmac) {
		t.Fatalf("hmac = %s", record.UsernameHMAC)
	}
	if !record.ExpiresAt.Equal(expiresAt.UTC()) {
		t.Fatalf("expires = %s, want %s", record.ExpiresAt, expiresAt.UTC())
	}
	ciphertext, err := base64.StdEncoding.DecodeString(record.CiphertextAgeB64)
	if err != nil {
		t.Fatal(err)
	}
	reader, err := age.Decrypt(bytes.NewReader(ciphertext), identity)
	if err != nil {
		t.Fatal(err)
	}
	var decoded SignupPayload
	if err := json.NewDecoder(reader).Decode(&decoded); err != nil {
		t.Fatal(err)
	}
	if decoded.Username != "alice" || decoded.PaymentReference == "" {
		t.Fatalf("decoded payload = %+v", decoded)
	}
	found, err := pending.ScanHMAC(hmac)
	if err != nil {
		t.Fatal(err)
	}
	if !found {
		t.Fatal("expected pending hmac match")
	}
	if n, err := pending.Sweep(now.Add(30 * time.Minute)); err != nil || n != 0 {
		t.Fatalf("early sweep n=%d err=%v", n, err)
	}
	if n, err := pending.Sweep(now.Add(2 * time.Hour)); err != nil || n != 1 {
		t.Fatalf("expired sweep n=%d err=%v", n, err)
	}
}

func TestPendingListSkipsInvalidFilenames(t *testing.T) {
	dir := filepath.Join(t.TempDir(), "pending")
	if err := os.MkdirAll(dir, 0o700); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(filepath.Join(dir, "not-an-id.json"), []byte("{}"), 0o600); err != nil {
		t.Fatal(err)
	}
	items, err := NewPending(dir, time.Hour, nil).List()
	if err != nil {
		t.Fatal(err)
	}
	if len(items) != 0 {
		t.Fatalf("items = %d, want 0", len(items))
	}
}
