package roll

// Result can be a result from any main type within Roll. It just needs to have
// a good String method to concisely display its output
type Result interface {
	String()
}

type DiceRoll struct {
}

type Opt struct {
}
