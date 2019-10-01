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
