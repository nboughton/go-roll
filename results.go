package roll

// Results is used for a collection of Result structs representing difference Die types
type Results []Result

// Dice returns the number and die types of all results in the current set
func (r Results) Dice() Set {
	var s Set

	for _, rs := range r {
		s = append(s, Dice{N: len(rs.rolls), Die: rs.die})
	}

	return s
}

// Min returns the minimum possible roll of any result in the set
func (r Results) Min() int {
	min := r.Max()

	for _, result := range r {
		m := result.Min()
		if m < min {
			min = m
		}
	}

	return min
}

// Max returns the maximum possible roll of any result in the set
func (r Results) Max() int {
	max := 0

	for _, result := range r {
		m := result.Max()
		if m > max {
			max = m
		}
	}

	return max
}
