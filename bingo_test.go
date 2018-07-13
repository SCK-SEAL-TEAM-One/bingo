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

func Test_NewGame_Input_Aey_Should_Be_Board_Player_Aey(t *testing.T) {
	playerName := Player{Name: "Aey"}
	expected := Game{
		Player: Player{
			Name: "Aey",
		}, 
		Board: Board{
			Player: Player{Name: "Aey"},
		},
	}

	actual := NewGame(playerName)

	if actual.Board.Player != expected.Board.Player {
		t.Errorf("Expected is %v but got %v", expected, actual)
	}
}
