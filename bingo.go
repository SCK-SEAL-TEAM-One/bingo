package bingo

type State struct {
	Symbol string
	Number int
}

func CheckBoard(number int) *State {

	return &State{Symbol: "o", Number: number}
}
