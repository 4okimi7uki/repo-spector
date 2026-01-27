package aggregate

import "testing"

func TestToExcludeSet(t *testing.T) {
	input := []string{" Go ", "", "Python", "  ", "go"}
	set := ToExcludeSet(input)

	if len(set) != 2 {
		t.Fatalf("expected 2 items, got %d", len(set))
	}
	if _, ok := set["go"]; !ok {
		t.Fatalf("expected set to contain go")
	}
	if _, ok := set["python"]; !ok {
		t.Fatalf("expected set to contain python")
	}
}

func TestIsExcludeLang(t *testing.T) {
	set := map[string]struct{}{"go": {}}

	if !IsExcludeLang("Go", set) {
		t.Fatalf("expected Go to be excluded")
	}
	if IsExcludeLang("Rust", set) {
		t.Fatalf("did not expect Rust to be excluded")
	}
}
