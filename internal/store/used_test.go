package store

import (
	"bytes"
	"crypto/rand"
	"os"
	"path/filepath"
	"testing"
)

func TestUsedScanAndConsume(t *testing.T) {
	dir := t.TempDir()
	path := filepath.Join(dir, "used.bin")
	entries := make([]byte, 4*HMACSize)
	if _, err := rand.Read(entries); err != nil {
		t.Fatal(err)
	}
	target := bytes.Repeat([]byte{9}, HMACSize)
	copy(entries[2*HMACSize:3*HMACSize], target)
	if err := os.WriteFile(path, entries, 0o600); err != nil {
		t.Fatal(err)
	}
	used := NewUsed(path, 4)
	found, scanned, err := used.Scan(target)
	if err != nil {
		t.Fatal(err)
	}
	if !found || scanned != 4 {
		t.Fatalf("found=%v scanned=%d", found, scanned)
	}
	slot, err := used.Consume("aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa")
	if err != nil {
		t.Fatal(err)
	}
	if slot != 4 {
		t.Fatalf("slot out of range: %d", slot)
	}
	st, err := os.Stat(path)
	if err != nil {
		t.Fatal(err)
	}
	if st.Size() != int64(4*HMACSize) {
		t.Fatalf("size changed to %d", st.Size())
	}
	overlay, err := os.Stat(path + ".overlay")
	if err != nil {
		t.Fatal(err)
	}
	if overlay.Size() != HMACSize {
		t.Fatalf("overlay size = %d, want %d", overlay.Size(), HMACSize)
	}
	found, scanned, err = used.Scan(bytes.Repeat([]byte{0xaa}, HMACSize))
	if err != nil {
		t.Fatal(err)
	}
	if !found || scanned != 5 {
		t.Fatalf("overlay found=%v scanned=%d", found, scanned)
	}
}

func TestUsedValidateSize(t *testing.T) {
	path := filepath.Join(t.TempDir(), "used.bin")
	if err := os.WriteFile(path, []byte{1, 2, 3}, 0o600); err != nil {
		t.Fatal(err)
	}
	if err := NewUsed(path, 1).Validate(); err == nil {
		t.Fatal("expected size mismatch")
	}
}
