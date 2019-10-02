# go-roll
go-roll is a dice rolling library written in Go. It will eventually include Roll Tables and possibly some other bits and pieces useful for writing 
code relevant to table top RPGs.

Roll is a successor to go-dice. It's designed to facilitate chaining roll results and currently includes a simple lexxer for rolling simple dice
strings. It supports keeping high/low results as preferred and exploding dice.

There are a few example applications in the cmd/ folder.

  - dnd-stats
    - Rolls dnd character stats using the 4d6 drop lowest method
  - dprob
    - Calculates probability of rolling a set of results
  - fate
    - Rolls a standard set of 4 Fate dice
  - pgraph
    - Renders a plot of dice sets as a png