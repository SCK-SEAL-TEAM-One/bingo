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
