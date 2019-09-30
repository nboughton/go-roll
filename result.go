package roll

import (
	"sort"
	"strings"
)

// Result represents a set of dice rolls that satsifies the sort interface
type Result []Face

func (r Result) Len() int           { return len(r) }
func (r Result) Less(i, j int) bool { return r[i].N < r[j].N }
func (r Result) Swap(i, j int)      { r[i], r[j] = r[j], r[i] }

func (r Result) String() string {
	var out []string

	sort.Sort(r)
	for _, f := range r {
		out = append(out, f.Value)
	}

	return strings.Join(out, ", ")
}

// Range provides readable identifiers for selecting HIGH/LOW values as keep or drop
type Range int

// Ranges for Result.Keep/Drop
const (
	HIGH Range = iota
	LOW
)

// Keep returns a new result struct containing the highest or lowest n results
func (r Result) Keep(n int, hl Range) Result {
	var out Result

	sort.Sort(r)
	switch hl {
	case HIGH:
		out = r[len(r)-n:]
	case LOW:
		out = r[:n]
	}

	return out
}

// Ints returns just the number values (useful for running totals)
func (r Result) Ints() []int {
	var out []int

	for _, n := range r {
		out = append(out, n.N)
	}

	return out
}

// Sum returns the total numerical value of a result set
func (r Result) Sum() int {
	var s int

	for _, n := range r {
		s += n.N
	}

	return s
}

// Explode recursively rerolls d Die for any results included in match and returns a completed
// Result set with all exploded items
func (r Result) Explode(match []int, d Die) Result {
	var out Result

	return out
}
