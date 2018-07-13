package bingo

import (
	"testing"
)

func Test_NewPlayer_Input_Aey_Should_Be_Player_Aey(t *testing.T) {
	playerName := "Aey"
	playerNameStruct := Player{Name: playerName}
	expected := playerNameStruct

	actual := NewPlayer(playerName)

	if expected != actual {
		t.Errorf("Should be %v but got %v", expected, actual)
	}
}
func Test_NewBoard_Input_Player_Aey_Should_Be_Board_Player_Aey(t *testing.T) {
	player := Player{Name: "Aey"}

	expected := Board{
		Player: Player{Name: "Aey"},
		Size:   Size{X: 5, Y: 5},
	}

	actual := NewBoard(player)

	if actual.Player != expected.Player || expected.Size != actual.Size {
		t.Errorf("expected %v but got %v ", actual, expected)
	}
}

func Test_CheckBingo_Should_Be_True(t *testing.T) {
	boardNumber := []int{7, 3, 12, 9, 1}
	slots := make([][]State, 5)
	for index := 0; index < 5; index++ {
		slots[index] = make([]State, 5)
		slots[index][0] = State{
			Symbol: "O",
			Number: boardNumber[index],
		}
	}
	board := Board{
		Slots:  slots,
		Player: Player{},
		Size:   Size{X: 5, Y: 5},
	}
	game := Game{Board: board}
	expected := true

	actual := game.CheckBingo()

	if expected != actual {
		t.Errorf("Should be %t but got %t", expected, actual)
	}
}
