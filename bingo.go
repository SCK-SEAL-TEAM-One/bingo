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
	Player Player
	Board  Board
}

func (game *Game) CheckBoard(number int) *State {
	for i := 0; i < game.Board.Size.X; i++ {
		for j := 0; j < game.Board.Size.Y; j++ {
			if game.Board.Slots[i][j].Number == number {
				return &game.Board.Slots[i][j]
			}

		}
	}

	return nil
}

func NewBoard(player Player) Board {
	slots := make([][]State, 5)
	slots[0] = []State{}
	for index := 0; index < 5; index++ {
		slots[index] = make([]State, 5)
	}

	return Board{
		Slots:  slots,
		Player: player,
		Size:   Size{X: 5, Y: 5},
	}
}
