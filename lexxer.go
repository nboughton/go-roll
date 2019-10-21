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

/*FromString reads a dice string like 3d6X6Kh2: roll 3 6 sided dice, exploding 6s, and keep the lowest 2, and returns a Result struct
FromString will return an error containing any unparsed characters which can be used to check syntax and troubleshoot dice strings.
FromString does some minimal checking of input:
	* Comma separated number lists (explode, dropN, keepN etc) are filtered to remove duplicate numbers and an error will be raised if the number of arguments exceeds or equals the faces of the die as this likely means that it will match all items.
	* Keep/Drop operations will return an error if the
*/
func FromString(s string) (Result, error) {
	var roll Result

	// Remove leading/trailing spaces
	s = strings.TrimSpace(s)

	// Get indices
	for i, op := range lexMatcher.FindAllString(s, -1) {
		if i == 0 && !lexDice.MatchString(op) {
			return roll, fmt.Errorf("%s: first operation must be a dice string (3d6 etc)", op)
		}

		switch {
		case lexDice.MatchString(op):
			n, die := parseDie(op)
			if n == 0 || die.faces.Len() == 0 || die.faces.Len() == 1 {
				return roll, fmt.Errorf("non-euclidean die: %s", op)
			}

			roll = Roll(n, die)

		case lexKeep.MatchString(op):
			k, m := parseKeep(op)
			if k < 1 {
				return roll, fmt.Errorf("cannot keep a negative quantity of dice: %s", op)
			}
			roll = roll.Keep(k, m)

		case lexKeepN.MatchString(op):
			n := parseComSepN(op)
			if len(n) >= len(roll.die.faces) {
				return roll, fmt.Errorf("numbers kept equals or exceeds faces of die")
			}
			roll = roll.KeepN(n...)

		case lexDrop.MatchString(op):
			d, m := parseDrop(op)
			if d > roll.rolls.Len() {
				return roll, fmt.Errorf("cannot drop more dice than rolled: %s", op)
			}
			roll = roll.Drop(d, m)

		case lexDropN.MatchString(op):
			n := parseComSepN(op)
			if len(n) >= len(roll.die.faces) {
				return roll, fmt.Errorf("numbers dropped equals or exceeds faces of die")
			}
			roll = roll.DropN(parseComSepN(op)...)

		case lexExp.MatchString(op):
			n := parseComSepN(op)
			if len(n) >= len(roll.die.faces) {
				return roll, fmt.Errorf("numbers exploded equals or exceeds faces of die")
			}
			roll = roll.Explode(n...)

		default:
			return roll, fmt.Errorf("invalid operation: %s", op)
		}

		// Remmove parsed op from s
		s = strings.Replace(s, op, "", 1)
	}

	if len(s) > 0 {
		return roll, fmt.Errorf("unparsed characters: %s", s)
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
	m, check := []int{}, make(map[int]int)

	for _, tok := range lexNum.FindAllString(s, -1) {
		n, err := strconv.Atoi(tok)
		if err != nil {
			continue
		}

		check[n]++ // Only count each number once
		if check[n] == 1 {
			m = append(m, n)
		}
	}

	return m
}
