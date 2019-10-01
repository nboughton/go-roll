package main

import (
	"fmt"
	"os"

	"github.com/nboughton/go-roll"
)

func main() {
	r, err := roll.FromString(os.Args[1])
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("Total:\t%d\nRolls:\t%s\n", r.Sum(), r)
}
