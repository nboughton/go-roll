# go-roll
go-roll is a dice rolling library written in Go. It will eventually include Roll Tables and possibly some other bits and pieces useful for writing 
code relevant to table top RPGs.

Roll is a successor to go-dice. It's designed to facilitate chaining roll results and currently includes a simple lexxer for rolling dice
strings. It supports keeping and dropping high/low results as preferred, and exploding dice.

Dice string syntax support is:

  - ndx: 3d6, 4d10 etc. A Dice string must begin with this.
  - K(h|l)x: Kh1, Kl2 etc. Keep highest or lowest n dice.
  - Kn1,2,3...: Only keep rolls matching 1,2,3...
  - D(h|l)x: Dl1, Dh2 etc. Drop the highest or lowest n dice.
  - Dn1,2,3...: Drop all rolls matching 1,2,3...
  - Xn,n...: X9,10 etc. Explode any dice in the set
  
These can be chained with a string like 4d10Kh3X10Dl1 to produce an end result.

Quickstart:
```Go
package main

import (
  "fmt"
  
  "github.com/nboughton/go-roll"
)

func main() {
  fmt.Println(roll.Roll(4, roll.D10).Keep(3, roll.HIGH).Explode(10).Drop(1, roll.Low))
}
```

There are a few example applications in the cmd/ folder.

  - dnd-stats
    - Rolls dnd character stats using the 4d6 drop lowest method
  - dprob
    - Calculates probability of rolling a set of results
  - fate
    - Rolls a standard set of 4 Fate dice
  - pgraph
    - Renders a plot of dice sets as a png
  - roll
    - Rolls a dice string and prints the result