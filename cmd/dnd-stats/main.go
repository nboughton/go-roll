package main

// Example program using go-roll to output D&D character stats. This example
// uses the 4D6 drop lowest method

import (
	"fmt"

	"github.com/nboughton/go-roll"
)

func main() {
	// Roll 6 times
	for i := 0; i < 6; i++ {
		fmt.Println(roll.Dice(4, roll.D6).Keep(3, roll.HIGH).Sum())
	}
}
