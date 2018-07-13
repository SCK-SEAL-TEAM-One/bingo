package bingo

import "testing"

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
