package roll

import (
	"math/rand"
	"sort"
	"time"
)

func init() {
	rand.Seed((time.Now().UnixNano()))
}

// Roller interface for any type that can Roll and return a single Result set
type Roller interface {
	Roll() Result
	Min() int
	Max() int
}

// Roll rolls n Die and returns a result set
func Roll(n int, d Die) Result {
	r := Result{die: d}

	for i := 0; i < n; i++ {
		r.rolls = append(r.rolls, d.Roll())
	}

	return r
}

// Face represents a single face of a die and can have both a number and textual
// value for custom dice
type Face struct {
	N     int
	Value string
}

// Faces is a set of Face structs that satisfies the sort interface
type Faces []Face

func (f Faces) Len() int           { return len(f) }
func (f Faces) Less(i, j int) bool { return f[i].N < f[j].N }
func (f Faces) Swap(i, j int)      { f[i], f[j] = f[j], f[i] }

// NewDie returns a uniq Die useful for custom dice systems like FFG/Genesys
func NewDie(faces []Face) Die {
	return Die{faces: faces}
}

// Die represents a single rollable die
type Die struct {
	faces Faces
}

// Roll returns a random face of d Die
func (d Die) Roll() Face {
	return d.faces[rand.Intn(len(d.faces))]
}

// Min returns the lowest value face of Die
func (d Die) Min() Face {
	sort.Sort(d.faces)
	return d.faces[0]
}

// Max returns the highest value face of Die
func (d Die) Max() Face {
	sort.Sort(d.faces)
	return d.faces[len(d.faces)-1]
}

// Dice represents a number of Die
type Dice struct {
	N   int
	Die Die
}

// Roll dice and return the Result
func (d Dice) Roll() Result {
	return Roll(d.N, d.Die)
}

// Min returns the minimum possible roll for a Dice struct
func (d Dice) Min() int {
	t := 0

	for i := 0; i < d.N; i++ {
		t += d.Die.Min().N
	}

	return t
}

// Max returns the maximum possible roll for a Dice struct
func (d Dice) Max() int {
	t := 0

	for i := 0; i < d.N; i++ {
		t += d.Die.Max().N
	}

	return t
}

// Set represents a collection of different Die types and numbers thereof
type Set []Dice

// Roll all items in a set and return the Results
func (s Set) Roll() Results {
	var r Results

	for _, d := range s {
		r = append(r, d.Roll())
	}

	return r
}

// Min returns the minimum possile numerical result for Set s
func (s Set) Min() int {
	t := 0

	for _, d := range s {
		t += d.Min()
	}

	return t
}

// Max returns the maximum possile numerical result for Set s
func (s Set) Max() int {
	t := 0

	for _, d := range s {
		t += d.Max()
	}

	return t
}
