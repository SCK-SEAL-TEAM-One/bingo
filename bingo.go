package bingo

type Player struct {
	Name string
}

type State struct {
	Symbol string
	Number int
}

type Size struct {
	X int
	Y int
}

type Board struct {
	Slots  [][]State
	Player Player
	Size   Size
}
type Game struct {
	Player
	Board
}

func NewBoard(player Player) Board {
	slots := make([][]State, 5)
	for index := 0; index < 5; index++ {
		slots[index] = make([]State, 5)
	}

	return Board{
		Slots:  slots,
		Player: player,
		Size:   Size{X: 5, Y: 5},
	}
}

func NewGame(player Player) Game {

	return Game{Player: player, Board: NewBoard(player)}
}
