package roll

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

var (
	dice    = regexp.MustCompile(`\d+d\d+`)
	keep    = regexp.MustCompile(`K(l|h)\d+`)
	exp     = regexp.MustCompile(`X[\d,]+`)
	matcher = regexp.MustCompile(fmt.Sprintf("(%s|%s|%s)", dice, keep, exp))
)

// FromString reads a dice string like 3d6X6Kh2: roll 3 6 sided dice, exploding 6s, and keep the lowest 2, and return a Result struct
func FromString(s string) (Result, error) {
	var roll Result

	// Get indices
	for i, op := range matcher.FindAllString(s, -1) {
		if i == 0 && !dice.MatchString(op) {
			return roll, fmt.Errorf("First argument must be a dice string (3d6 etc)")
		}

		switch {
		case dice.MatchString(op):
			roll = Dice(dieFromString(op))

		case keep.MatchString(op):
			roll = roll.Keep(keepFromString(op))

		case exp.MatchString(op):
			roll = roll.Explode(expFromString(op)...)

		default:
			return roll, fmt.Errorf("Invalid operation: %s", op)
		}

	}

	return roll, nil
}

func dieFromString(s string) (int, Die) {
	n, f := 0, 0
	fmt.Sscanf(s, "%dd%d", &n, &f)
	return n, Die{faces: makeFaces(f)}
}

func keepFromString(s string) (int, Range) {
	r, n := LOW, 0

	if strings.Contains(s, "l") {
		fmt.Sscanf(s, "Kl%d", &n)
	} else {
		r = HIGH
		fmt.Sscanf(s, "Kh%d", &n)
	}

	return n, r
}

func expFromString(s string) []int {
	m := []int{}

	for _, char := range s {
		n, err := strconv.Atoi(string(char))
		if err != nil {
			continue
		}

		m = append(m, n)
	}

	return m
}
