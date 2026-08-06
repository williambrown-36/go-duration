package duration

import "testing"

func TestParse(t *testing.T) {
	cases := map[string]int{
		"90s":    90,
		"1h30m":  5400,
		"2d":     172800,
		"1h1m1s": 3661,
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
	invalid := []string{"later", "1h later", "later 1h", "1h later 30m", "1h30"}
	for _, in := range invalid {
		if _, err := Parse(in); err == nil {
			t.Errorf("Parse(%q) should reject malformed input", in)
		}
	}
}

func TestFormat(t *testing.T) {
	cases := map[int]string{
		0:      "0s",
		90:     "1m30s",
		5400:   "1h30m",
		172800: "2d",
		90061:  "1d1h1m1s",
	}
	for in, want := range cases {
		if got := Format(in); got != want {
			t.Errorf("Format(%d) = %q, want %q", in, got, want)
		}
	}
}

func TestFormatRoundTrip(t *testing.T) {
	for _, seconds := range []int{0, 1, 59, 60, 90, 3600, 5400, 86400, 90061} {
		formatted := Format(seconds)
		got, err := Parse(formatted)
		if err != nil {
			t.Fatalf("Parse(Format(%d)) errored: %v", seconds, err)
		}
		if got != seconds {
			t.Errorf("Parse(Format(%d)) = %d", seconds, got)
		}
	}
}
