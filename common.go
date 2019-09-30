package roll

import (
	"strconv"
)

var (
	// D2 is a two sided die, effectively a coin flip
	D2 = Die{faces: makeFaces(2)}
	// D3 is a 3 sided die, typically done by rolling a d6 and dividing the result by 2
	D3 = Die{faces: makeFaces(3)}
	// D4 also known as the caltrop is a 4 sided pyramid
	D4 = Die{faces: makeFaces(4)}
	// D6 is a 6 sided die
	D6 = Die{faces: makeFaces(6)}
	// D8 is an 8 sided die
	D8 = Die{faces: makeFaces(8)}
	// D10 is a 10 sided die
	D10 = Die{faces: makeFaces(10)}
	// D12 is the 12 sided dodecahedron that doesn't get enough love
	D12 = Die{faces: makeFaces(12)}
	// D20 is the screeching diva that hogs all the limelight
	D20 = Die{faces: makeFaces(20)}
	// D100 aka D%, typically two d10s reading one as tens and the other as units
	D100 = Die{faces: makeFaces(100)}
)

func makeFaces(n int) []Face {
	var f []Face

	for i := 1; i <= n; i++ {
		f = append(f, Face{i, strconv.Itoa(i)})
	}

	return f
}