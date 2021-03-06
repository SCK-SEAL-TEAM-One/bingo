package bingo

import (
	"testing"
)

func Test_CheckBoard_Input_RandomNumber_7_Shuold_Be_State_Symbol_o_And_Number_7(t *testing.T) {
	player := Player{Name: "Aey"}
	states := State{Symbol: "-", Number: 7}
	board := NewBoard(player)
	board.Slots[0][0] = states
	game := Game{Player: player, Board: board}
	randomNumber := 7
	expected := State{Symbol: "-", Number: 7}

	actual := *game.CheckBoard(randomNumber)

	if expected != actual {
		t.Errorf("Expectd %v but got %v", expected, actual)
	}
}
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
	player := Player{Name: "Aey"}
	expected := Game{
		Player: Player{
			Name: "Aey",
		},
		Board: Board{
			Player: Player{Name: "Aey"},
		},
	}
	actual := NewGame(player)

	if actual.Board.Player != expected.Board.Player {
		t.Errorf("Expected is %v but got %v", expected, actual)
	}
}

func Test_FillBoard_Input_Pointer_State_Should_Be_True(t *testing.T){
	state := State{Symbol:"", Number: 15}
	addressState := &state 
	player :=NewPlayer("Aoi")
	game := NewGame(player)
	
	expected := true
	actual := game.FillBoard(addressState)

	if actual != expected || state.Symbol != "O" {
		t.Errorf("Expected is %v but got %v", expected, actual)
	}
} 

func Test_GetWinner_Should_Be_Aey(t *testing.T) {
	expected := "Aey"
	game := Game{
		Board: Board{
			Player: Player{Name: "Aey"},
		},
	}

	actual := game.GetWinner()

	if actual != expected {
		t.Errorf("expected %v but got %v ", expected, actual)
	}
}

func Test_GetWinner_Should_Be_Aoi(t *testing.T) {
	expected := "Aoi"
	game := Game{
		Board: Board{
			Player: Player{Name: "Aoi"},
		},
	}

	actual := game.GetWinner()

	if actual != expected {
		t.Errorf("expected %v but got %v ", expected, actual)
	}
}
