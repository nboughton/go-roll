package main

import (
	"fmt"

	"github.com/nboughton/go-roll"
)

func main() {
	for i := 0; i < 100; i++ {
		fmt.Printf("%+v\n", roll.Dice(6, roll.D10).Explode(9, 10))
	}
}
