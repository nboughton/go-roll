// Package roll defines a number of types for creating individual die and sets. Most
// common use cases will just require the Roll() function to create a Result set
// can chain Keep and Explode methods
package roll

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
