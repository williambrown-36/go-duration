// Package duration parses compact durations like 1h30m, dependency-free.
package duration

import (
	"fmt"
	"regexp"
	"strconv"
)

var unit = map[string]int{"s": 1, "m": 60, "h": 3600, "d": 86400}
var token = regexp.MustCompile(`(\d+)([smhd])`)

func Parse(s string) (int, error) {
	m := token.FindAllStringSubmatch(s, -1)
	if len(m) == 0 {
		return 0, fmt.Errorf("no duration tokens")
	}
	total := 0
	for _, g := range m {
		n, _ := strconv.Atoi(g[1])
		total += n * unit[g[2]]
	}
	return total, nil
}
