package bingo

type Player struct {
	Name string
}

func NewPlayer(name string) Player {
	playerStruct := Player{Name: name}
	return playerStruct
}
