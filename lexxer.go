package roll

import (
	"fmt"
	"regexp"
	"strings"
)

var (
	dice    = regexp.MustCompile(`\d+d\d+`)
	keep    = regexp.MustCompile(`K(l|h)\d+`)
	exp     = regexp.MustCompile(`X[\d,]+`)
	matcher = regexp.MustCompile(fmt.Sprintf("(%s|%s|%s)", dice, keep, exp))
)

// FromString reads a dice string like 3d6X6Kh2: roll 3 6 sided dice, exploding 6s, and keep the lowest 2, and return a Result struct
func FromString(s string) string {
	// Get indices
	for _, t := range matcher.FindAllString(s, -1) {
		switch {
		case dice.MatchString(t):

		}
		fmt.Println(t)
	}

	return ""
	/*
		nD, d, dErr := dieFromString(s)
		if dErr != nil {
			return Result{}, dErr
		}

		nK, r, kErr := keepFromString(s)

		roll := Dice(nD, d)
		fmt.Println(roll)
		if kErr == nil {
			roll = roll.Keep(nK, r)
		}

		return roll, nil
	*/
}

func dieFromString(s string) (int, Die, error) {
	m := dice.FindString(s)
	if m == "" {
		return 0, Die{}, fmt.Errorf("No die string found")
	}

	n, f := 0, 0
	fmt.Sscanf(m, "%dd%d", &n, &f)
	return n, Die{faces: makeFaces(f)}, nil
}

func keepFromString(s string) (int, Range, error) {
	k := keep.FindString(s)
	if k == "" {
		return 0, -1, fmt.Errorf("No keep string found")
	}

	var (
		r Range
		n int
	)

	if strings.Contains(k, "l") {
		r = LOW
		fmt.Sscanf(k, "Kl%d", &n)
	} else {
		r = HIGH
		fmt.Sscanf(k, "Kh%d", &n)
	}

	return n, r, nil
}
