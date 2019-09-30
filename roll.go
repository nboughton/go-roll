package roll

import (
	"math/rand"
	"sort"
	"strings"
	"time"
)

func init() {
	rand.Seed((time.Now().UnixNano()))
}

// NewDie returns a uniq Die useful for custom dice systems like FFG/Genesys
func NewDie(faces []Face) Die {
	return Die{faces: faces}
}

// Die represents a single rollable die
type Die struct {
	faces []Face
}

// Face represents a single face of a die and can have both a number and textual
// value for custom dice
type Face struct {
	N     int
	Value string
}

// Roll returns a random face of d Die
func (d Die) Roll() Face {
	return d.faces[rand.Intn(len(d.faces))]
}

// Roll is the core function of this package
func Roll(n int, d Die) Result {
	var r Result

	for i := 0; i < n; i++ {
		r = append(r, d.Roll())
	}

	return r
}

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
