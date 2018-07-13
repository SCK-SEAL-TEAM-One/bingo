package bingo

import "testing"

func Test_CheckBoard_Input_RandomNumber_7_Shuold_Be_State_Symbol_o_And_Number_7(t *testing.T) {
	player := Player{Name: "Aey"}
	states := State{Symbol: "-", Number: 7}
	board := NewBoard(player)
	board.Slots[0][0] = states
	game := Game{Player: player, Board: board}
	randomNumber := 7
	expected := State{Symbol: "-", Number: 7}

	state := game.CheckBoard(randomNumber)

	actual := *state
	if expected != actual {
		t.Errorf("Expectd %v but got %v", expected, actual)
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
