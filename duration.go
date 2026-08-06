// Package duration parses and formats compact durations like 1h30m, dependency-free.
package duration

import (
	"fmt"
	"regexp"
	"strconv"
)

var unit = map[string]int{"s": 1, "m": 60, "h": 3600, "d": 86400}
var token = regexp.MustCompile(`(\d+)([smhd])`)
var duration = regexp.MustCompile(`^(?:\d+[smhd])+$`)

func Parse(s string) (int, error) {
	if !duration.MatchString(s) {
		return 0, fmt.Errorf("invalid duration")
	}
	m := token.FindAllStringSubmatch(s, -1)
	total := 0
	for _, g := range m {
		n, _ := strconv.Atoi(g[1])
		total += n * unit[g[2]]
	}
	return total, nil
}

// Format returns seconds as a compact duration using the largest units first.
func Format(seconds int) string {
	if seconds <= 0 {
		return "0s"
	}

	result := ""
	for _, part := range []struct {
		suffix  string
		seconds int
	}{
		{"d", unit["d"]},
		{"h", unit["h"]},
		{"m", unit["m"]},
		{"s", unit["s"]},
	} {
		if n := seconds / part.seconds; n > 0 {
			result += strconv.Itoa(n) + part.suffix
			seconds %= part.seconds
		}
	}
	return result
}
