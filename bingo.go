package bingo

import "math/rand"

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

func NewPlayer(name string) Player {
	playerStruct := Player{Name: name}
	return playerStruct
}

func NewBoard(player Player) Board {
	slots := make([][]State, 5)
	for index := 0; index < 5; index++ {
		slots[index] = make([]State, 5)
	}
	slots[3][3].Symbol = "O"
	return Board{
		Slots:  slots,
		Player: player,
		Size:   Size{X: 5, Y: 5},
	}
}

func NewGame(player Player) Game {
	return Game{Player: player, Board: NewBoard(player)}
}

func (g Game) FillBoard(state *State) bool {
	if state.Symbol == "" {
		state.Symbol = "O"
		return true
	}
	return false
}

func (g *Game) CheckBoard(number int) *State {
	for i := 0; i < g.Board.Size.X; i++ {
		for j := 0; j < g.Board.Size.Y; j++ {
			if g.Board.Slots[i][j].Number == number {
				return &g.Board.Slots[i][j]
			}
		}
	}
	return nil
}

func (g Game) Draw() int {
	return rand.Intn(75) + 1
}

func (g Game) CheckBingo() bool {
	for i := 0; i < g.Board.Size.Y; i++ {
		if g.Board.Slots[i][0].Symbol == "-" {
			return false
		}
	}
	return true
}

func (g Game) GetWinner() string {
	return g.Board.Player.Name
}
