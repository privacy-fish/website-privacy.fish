package refgen

import (
	"crypto/rand"
	"errors"
	"strings"
)

const alphabet = "0123456789ABCDEFGHJKMNPQRSTVWXYZ"

func New() (string, error) {
	raw := make([]byte, 20)
	if _, err := rand.Read(raw); err != nil {
		return "", err
	}
	chars := make([]byte, 0, 24)
	for i, b := range raw {
		if i > 0 && i%4 == 0 {
			chars = append(chars, '-')
		}
		chars = append(chars, alphabet[int(b)&31])
	}
	return string(chars), nil
}

func Valid(s string) bool {
	if len(s) != 24 {
		return false
	}
	for i, c := range s {
		if i == 4 || i == 9 || i == 14 || i == 19 {
			if c != '-' {
				return false
			}
			continue
		}
		if !strings.ContainsRune(alphabet, c) {
			return false
		}
	}
	return true
}

var ErrInvalid = errors.New("invalid_reference")
