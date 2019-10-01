package roll

import (
	"sort"
	"strings"
)

// Result represents a set of dice rolls of a single die type
type Result struct {
	die   Die
	rolls []Face
}

// Satisfy the Sort interface
func (r Result) Len() int           { return len(r.rolls) }
func (r Result) Less(i, j int) bool { return r.rolls[i].N < r.rolls[j].N }
func (r Result) Swap(i, j int)      { r.rolls[i], r.rolls[j] = r.rolls[j], r.rolls[i] }

// Satisfy the String interface
func (r Result) String() string {
	var out []string

	sort.Sort(r)
	for _, f := range r.rolls {
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
	out := Result{die: r.die}

	sort.Sort(r)
	switch hl {
	case HIGH:
		out.rolls = r.rolls[len(r.rolls)-n:]
	case LOW:
		out.rolls = r.rolls[:n]
	}

	return out
}

// Explode recursively rerolls d Die for any results included in match and returns a completed
// Result set with all exploded items
func (r Result) Explode(match ...int) Result {
	var x func(store, results []Face) ([]Face, []Face)

	x = func(store, results []Face) ([]Face, []Face) {
		for i, result := range results {
			for _, m := range match {
				if result.N == m {
					// Roll the exploded die
					results = append(results, r.die.Roll())
					// Store the current match
					store = append(store, result)
					// Remove it from results so it doesn't get exploded again
					results = append(results[:i], results[i+1:]...)
					// Explode the resultant set
					store, results = x(store, results)
				}
			}
		}

		return store, results
	}

	var store []Face
	store, out := x(store, r.rolls)

	return Result{die: r.die, rolls: append(out, store...)}
}

// Ints returns just the number values (useful for running totals)
func (r Result) Ints() []int {
	var out []int

	for _, n := range r.rolls {
		out = append(out, n.N)
	}

	return out
}

// Sum returns the total numerical value of a result set
func (r Result) Sum() int {
	var s int

	for _, n := range r.rolls {
		s += n.N
	}

	return s
}
