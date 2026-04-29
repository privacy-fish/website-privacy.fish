package store

import (
	"crypto/subtle"
	"encoding/hex"
	"errors"
	"fmt"
	"io"
	"os"
)

const HMACSize = 32

type Used struct {
	Path        string
	Size        int64
	OverlayPath string
}

func NewUsed(path string, size int64) *Used {
	return &Used{Path: path, Size: size, OverlayPath: path + ".overlay"}
}

func (u *Used) Validate() error {
	st, err := os.Stat(u.Path)
	if err != nil {
		return err
	}
	want := u.Size * HMACSize
	if st.Size() != want {
		return fmt.Errorf("used.bin size mismatch: got %d bytes, want %d", st.Size(), want)
	}
	if st, err := os.Stat(u.OverlayPath); err == nil {
		if st.Size()%HMACSize != 0 {
			return fmt.Errorf("used overlay size mismatch: got %d bytes, not divisible by %d", st.Size(), HMACSize)
		}
	} else if !errors.Is(err, os.ErrNotExist) {
		return err
	}
	return nil
}

func (u *Used) Scan(target []byte) (bool, int64, error) {
	if len(target) != HMACSize {
		return false, 0, errors.New("target hmac must be 32 bytes")
	}
	foundBase, scannedBase, err := scanFixedEntries(u.Path, u.Size, target)
	if err != nil {
		return false, scannedBase, err
	}
	foundOverlay, scannedOverlay, err := scanOverlayEntries(u.OverlayPath, target)
	if err != nil {
		return false, scannedBase + scannedOverlay, err
	}
	return foundBase || foundOverlay, scannedBase + scannedOverlay, nil
}

func (u *Used) Consume(hexHMAC string) (int64, error) {
	h, err := hex.DecodeString(hexHMAC)
	if err != nil || len(h) != HMACSize {
		return 0, errors.New("username_hmac must be 64 hex chars")
	}
	if err := u.Validate(); err != nil {
		return 0, err
	}
	if found, idx, err := findFixedEntry(u.Path, u.Size, h); err != nil {
		return 0, err
	} else if found {
		return idx, nil
	}
	if found, idx, err := findOverlayEntry(u.OverlayPath, h); err != nil {
		return 0, err
	} else if found {
		return u.Size + idx, nil
	}
	f, err := os.OpenFile(u.OverlayPath, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0o600)
	if err != nil {
		return 0, err
	}
	defer f.Close()
	st, err := f.Stat()
	if err != nil {
		return 0, err
	}
	if st.Size()%HMACSize != 0 {
		return 0, fmt.Errorf("used overlay size mismatch: got %d bytes, not divisible by %d", st.Size(), HMACSize)
	}
	idx := st.Size() / HMACSize
	if _, err := f.Write(h); err != nil {
		return 0, err
	}
	return u.Size + idx, nil
}

func scanFixedEntries(path string, entries int64, target []byte) (bool, int64, error) {
	f, err := os.Open(path)
	if err != nil {
		return false, 0, err
	}
	defer f.Close()
	return scanEntries(f, entries, target)
}

func scanOverlayEntries(path string, target []byte) (bool, int64, error) {
	f, err := os.Open(path)
	if errors.Is(err, os.ErrNotExist) {
		return false, 0, nil
	}
	if err != nil {
		return false, 0, err
	}
	defer f.Close()
	st, err := f.Stat()
	if err != nil {
		return false, 0, err
	}
	if st.Size()%HMACSize != 0 {
		return false, 0, fmt.Errorf("used overlay size mismatch: got %d bytes, not divisible by %d", st.Size(), HMACSize)
	}
	return scanEntries(f, st.Size()/HMACSize, target)
}

func scanEntries(r io.Reader, entries int64, target []byte) (bool, int64, error) {
	buf := make([]byte, HMACSize)
	var found int
	var scanned int64
	for scanned < entries {
		_, err := io.ReadFull(r, buf)
		if err != nil {
			return false, scanned, err
		}
		found |= subtle.ConstantTimeCompare(buf, target)
		scanned++
	}
	return found == 1, scanned, nil
}

func findFixedEntry(path string, entries int64, target []byte) (bool, int64, error) {
	f, err := os.Open(path)
	if err != nil {
		return false, 0, err
	}
	defer f.Close()
	return findEntry(f, entries, target)
}

func findOverlayEntry(path string, target []byte) (bool, int64, error) {
	f, err := os.Open(path)
	if errors.Is(err, os.ErrNotExist) {
		return false, 0, nil
	}
	if err != nil {
		return false, 0, err
	}
	defer f.Close()
	st, err := f.Stat()
	if err != nil {
		return false, 0, err
	}
	if st.Size()%HMACSize != 0 {
		return false, 0, fmt.Errorf("used overlay size mismatch: got %d bytes, not divisible by %d", st.Size(), HMACSize)
	}
	return findEntry(f, st.Size()/HMACSize, target)
}

func findEntry(r io.Reader, entries int64, target []byte) (bool, int64, error) {
	buf := make([]byte, HMACSize)
	for i := int64(0); i < entries; i++ {
		if _, err := io.ReadFull(r, buf); err != nil {
			return false, 0, err
		}
		if subtle.ConstantTimeCompare(buf, target) == 1 {
			return true, i, nil
		}
	}
	return false, 0, nil
}
