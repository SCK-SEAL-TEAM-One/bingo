package bingo

import "testing"

func Test_CheckBoard_Input_RandomNumber_7_Shuold_Be_State_Symbol_o_And_Number_7(t *testing.T) {
	randomNumber := 7
	expected := State{Symbol: "o", Number: 7}

	state := CheckBoard(randomNumber)
	actual := *state
	if expected != actual {
		t.Errorf("Expectd %v but got %v", expected, actual)
	}
}
