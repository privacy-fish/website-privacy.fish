package refgen

import "testing"

func TestNewFormat(t *testing.T) {
	ref, err := New()
	if err != nil {
		t.Fatal(err)
	}
	if !Valid(ref) {
		t.Fatalf("invalid reference format: %q", ref)
	}
}

func TestValidRejectsAmbiguousAlphabet(t *testing.T) {
	for _, ref := range []string{
		"ILOU-0000-0000-0000-0000",
		"00000000000000000000",
		"0000-0000-0000-0000-000",
	} {
		if Valid(ref) {
			t.Fatalf("expected invalid reference: %q", ref)
		}
	}
}
