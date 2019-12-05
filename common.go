package roll

import (
	"strconv"
)

var (
	// D2 is a two sided die, effectively a coin flip
	D2 = NewDie(makeFaces(2))
	// D3 is a 3 sided die, typically done by rolling a d6 and dividing the result by 2
	D3 = NewDie(makeFaces(3))
	// D4 also known as the caltrop is a 4 sided pyramid
	D4 = NewDie(makeFaces(4))
	// D6 is a 6 sided die
	D6 = NewDie(makeFaces(6))
	// D8 is an 8 sided die
	D8 = NewDie(makeFaces(8))
	// D10 is a 10 sided die
	D10 = NewDie(makeFaces(10))
	// D12 is the 12 sided dodecahedron that doesn't get enough love
	D12 = NewDie(makeFaces(12))
	// D20 is the screeching diva that hogs all the limelight
	D20 = NewDie(makeFaces(20))
	// D100 aka D%, typically two d10s reading one as tens and the other as units
	D100 = NewDie(makeFaces(100))
	// Fate aka the Fate die
	Fate = NewDie(Faces{{-1, "[-]"}, {-1, "[-]"}, {0, "[ ]"}, {0, "[ ]"}, {1, "[+]"}, {1, "[+]"}})
	// D66 as used in Mutant: Year Zero
	D66 = makeD66()
	// D666 as used in Mutant: Year Zero
	D666 = makeD666()
)

func makeFaces(n int) Faces {
	var f Faces

	for i := 1; i <= n; i++ {
		f = append(f, Face{i, strconv.Itoa(i)})
	}

	return f
}

func makeD66() Faces {
	var f Faces

	for _, tens := range []string{"1", "2", "3", "4", "5", "6"} {
		for _, digits := range []string{"1", "2", "3", "4", "5", "6"} {
			n, _ := strconv.Atoi(tens + digits)
			f = append(f, Face{N: n, Value: tens + digits})
		}
	}

	return f
}

func makeD666() Faces {
	var f Faces

	for _, hundreds := range []string{"1", "2", "3", "4", "5", "6"} {
		for _, tens := range []string{"1", "2", "3", "4", "5", "6"} {
			for _, digits := range []string{"1", "2", "3", "4", "5", "6"} {
				n, _ := strconv.Atoi(hundreds + tens + digits)
				f = append(f, Face{N: n, Value: hundreds + tens + digits})
			}
		}
	}

	return f
}
