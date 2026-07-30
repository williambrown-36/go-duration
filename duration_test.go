package duration

import "testing"

func TestParse(t *testing.T) {
	cases := map[string]int{
		"90s":     90,
		"1h30m":   5400,
		"2d":      172800,
		"1h1m1s":  3661,
	}
	for in, want := range cases {
		got, err := Parse(in)
		if err != nil {
			t.Fatalf("Parse(%q) errored: %v", in, err)
		}
		if got != want {
			t.Errorf("Parse(%q) = %d, want %d", in, got, want)
		}
	}
	if _, err := Parse("later"); err == nil {
		t.Error("Parse should reject input with no duration tokens")
	}
}
