package sshkey

import (
	"crypto/ed25519"
	"crypto/rand"
	"crypto/rsa"
	"errors"
	"strings"
	"testing"

	"golang.org/x/crypto/ssh"
)

func TestParseAuthorizedCanonicalizesAndDedupes(t *testing.T) {
	key := testAuthorizedKey(t, "ed25519")
	withComments := key + " laptop\n" + key + " phone\n"
	keys, err := ParseAuthorized(withComments)
	if err != nil {
		t.Fatal(err)
	}
	if len(keys) != 1 {
		t.Fatalf("len(keys) = %d, want 1", len(keys))
	}
	if strings.Count(keys[0], " ") != 1 {
		t.Fatalf("canonical key kept comment: %q", keys[0])
	}
}

func TestParseAuthorizedRejectsBadInputs(t *testing.T) {
	tests := []struct {
		name string
		in   string
		err  error
	}{
		{name: "missing", in: "\n", err: ErrMissing},
		{name: "malformed", in: "ssh-ed25519 not-base64", err: ErrInvalid},
		{name: "wrong type", in: testAuthorizedKey(t, "rsa"), err: ErrWrongType},
		{name: "too many", in: manyEd25519Keys(t, 11), err: ErrTooMany},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := ParseAuthorized(tt.in)
			if !errors.Is(err, tt.err) {
				t.Fatalf("err = %v, want %v", err, tt.err)
			}
		})
	}
}

func manyEd25519Keys(t *testing.T, n int) string {
	t.Helper()
	var b strings.Builder
	for i := 0; i < n; i++ {
		b.WriteString(testAuthorizedKey(t, "ed25519"))
		b.WriteByte('\n')
	}
	return b.String()
}

func testAuthorizedKey(t *testing.T, kind string) string {
	t.Helper()
	var pub any
	switch kind {
	case "ed25519":
		p, _, err := ed25519.GenerateKey(rand.Reader)
		if err != nil {
			t.Fatal(err)
		}
		pub = p
	case "rsa":
		priv, err := rsa.GenerateKey(rand.Reader, 1024)
		if err != nil {
			t.Fatal(err)
		}
		pub = &priv.PublicKey
	default:
		t.Fatalf("unknown key kind %s", kind)
	}
	sshPub, err := ssh.NewPublicKey(pub)
	if err != nil {
		t.Fatal(err)
	}
	return strings.TrimSpace(string(ssh.MarshalAuthorizedKey(sshPub)))
}
