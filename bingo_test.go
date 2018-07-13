package bingo

import "testing"

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

func Test_Draw_Random_Between_1_To_75_Should_Be_57(t *testing.T) {
	expected := 57

	game := Game{}
	actual := game.Draw()

	if actual != expected {
		t.Errorf("expected %v but got %v ", expected, actual)
	}

}

func Test_Draw_Random_Between_1_To_75_Should_Be_13(t *testing.T) {
	expected := 13

	game := Game{}
	actual := game.Draw()

	if actual != expected {
		t.Errorf("expected %v but got %v ", expected, actual)
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

func Test_FillBoard_Input_Pointer_State_Should_Be_True(t *testing.T){
	state := State{Symbol:"", Number: 15}
	addressState := &state 
	expected := true
	actual := FillBoard(addressState)

	if actual != expected || state.Symbol != "O" {
		t.Errorf("Expected is %v but got %v", expected, actual)
	}
}