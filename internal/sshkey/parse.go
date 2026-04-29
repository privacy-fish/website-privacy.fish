package sshkey

import (
	"encoding/base64"
	"errors"
	"strings"

	"golang.org/x/crypto/ssh"
)

var (
	ErrMissing   = errors.New("keys_missing")
	ErrTooMany   = errors.New("keys_too_many")
	ErrInvalid   = errors.New("key_invalid")
	ErrWrongType = errors.New("key_wrong_type")
)

func ParseAuthorized(input string) ([]string, error) {
	lines := strings.FieldsFunc(input, func(r rune) bool { return r == '\n' || r == '\r' })
	out := make([]string, 0, len(lines))
	seen := make(map[string]struct{}, len(lines))
	for _, raw := range lines {
		line := strings.TrimSpace(raw)
		if line == "" {
			continue
		}
		key, _, _, _, err := ssh.ParseAuthorizedKey([]byte(line))
		if err != nil {
			return nil, ErrInvalid
		}
		if key.Type() != "ssh-ed25519" {
			return nil, ErrWrongType
		}
		canonical := "ssh-ed25519 " + base64.StdEncoding.EncodeToString(key.Marshal())
		if _, ok := seen[canonical]; ok {
			continue
		}
		seen[canonical] = struct{}{}
		out = append(out, canonical)
		if len(out) > 10 {
			return nil, ErrTooMany
		}
	}
	if len(out) == 0 {
		return nil, ErrMissing
	}
	return out, nil
}
