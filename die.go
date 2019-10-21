package roll

import (
	"math/rand"
	"sort"
	"time"
)

func init() {
	rand.Seed((time.Now().UnixNano()))
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

func (f Faces) contains(m int) bool {
	for _, n := range f {
		if n.N == m {
			return true
		}
	}

	return false
}

// NewDie returns a unique Die useful for custom dice systems like FFG/Genesys
func NewDie(faces Faces) Die {
	return Die{faces: faces}
}

// Die represents a single rollable die. Die mthods always return a Face that has both
// a numerical value and symbol represented as a string.
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
