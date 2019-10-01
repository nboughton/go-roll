package roll

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
