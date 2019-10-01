package roll

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
