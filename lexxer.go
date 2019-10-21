package roll

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

var (
	lexDice    = regexp.MustCompile(`\d+d\d+`)
	lexKeep    = regexp.MustCompile(`K(l|h)\d+`)
	lexKeepN   = regexp.MustCompile(`Kn[\d,]+`)
	lexDrop    = regexp.MustCompile(`D(l|h)\d+`)
	lexDropN   = regexp.MustCompile(`Dn[\d,]+`)
	lexExp     = regexp.MustCompile(`X[\d,]+`)
	lexNum     = regexp.MustCompile(`\d+`)
	lexMatcher = regexp.MustCompile(fmt.Sprintf("(%s|%s|%s|%s|%s|%s)", lexDice, lexKeep, lexKeepN, lexDrop, lexDropN, lexExp))
)

// FromString reads a dice string like 3d6X6Kh2: roll 3 6 sided dice, exploding 6s, and keep the lowest 2, and returns a Result struct
func FromString(s string) (Result, error) {
	var roll Result

	// Get indices
	for i, op := range lexMatcher.FindAllString(s, -1) {
		if i == 0 && !lexDice.MatchString(op) {
			return roll, fmt.Errorf("first argument must be a dice string (3d6 etc)")
		}

		switch {
		case lexDice.MatchString(op):
			if string(op[0]) == "0" {
				return roll, fmt.Errorf("cannot roll 0 dice")
			}
			roll = Roll(parseDie(op))

		case lexKeep.MatchString(op):
			roll = roll.Keep(parseKeep(op))

		case lexKeepN.MatchString(op):
			roll = roll.KeepN(parseComSepN(op)...)

		case lexDrop.MatchString(op):
			roll = roll.Drop(parseDrop(op))

		case lexDropN.MatchString(op):
			roll = roll.DropN(parseComSepN(op)...)

		case lexExp.MatchString(op):
			roll = roll.Explode(parseComSepN(op)...)

		default:
			return roll, fmt.Errorf("invalid operation: %s", op)
		}

	}

	return roll, nil
}

func parseDie(s string) (int, Die) {
	n, f := 0, 0
	fmt.Sscanf(s, "%dd%d", &n, &f)
	return n, Die{faces: makeFaces(f)}
}

func parseKeep(s string) (int, MatchType) {
	m, n := LOW, 0

	if strings.Contains(s, "l") {
		fmt.Sscanf(s, "Kl%d", &n)
	} else {
		m = HIGH
		fmt.Sscanf(s, "Kh%d", &n)
	}

	return n, m
}

func parseDrop(s string) (int, MatchType) {
	m, n := LOW, 0

	if strings.Contains(s, "l") {
		fmt.Sscanf(s, "Dl%d", &n)
	} else {
		m = HIGH
		fmt.Sscanf(s, "Dh%d", &n)
	}

	return n, m
}

func parseComSepN(s string) []int {
	m := []int{}

	for _, tok := range lexNum.FindAllString(s, -1) {
		n, err := strconv.Atoi(tok)
		if err != nil {
			continue
		}

		m = append(m, n)
	}

	return m
}
