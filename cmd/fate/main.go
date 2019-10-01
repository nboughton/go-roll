package main

import (
	"fmt"
	"strings"

	"github.com/nboughton/go-roll"
)

func main() {
	r := roll.Roll(4, roll.Fate)
	fmt.Printf("Total:\t%d\nValues:\t%s\n", r.Sum(), strings.Join(r.Values(), ", "))
}
