package roll

import (
	"math/rand"
	"time"
)

func init() {
	rand.Seed((time.Now().UnixNano()))
}

// Dice rolls n Die and returns a result set
func Dice(n int, d Die) Result {
	var r Result

	for i := 0; i < n; i++ {
		r = append(r, d.Roll())
	}

	return r
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
	Max   int
}

// Roll returns a random face of d Die
func (d Die) Roll() Face {
	return d.faces[rand.Intn(len(d.faces))]
}
