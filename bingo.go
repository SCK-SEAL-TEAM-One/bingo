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

func (g Game) Draw() int {

	return rand.Intn(75) + 1
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

func FillBoard(state *State) bool{
	if state.Symbol ==  "" {
		state.Symbol = "O" 
		return true 
	} 
	return false
}
func (g Game) GetWinner() string {

	return g.Board.Player.Name
}
