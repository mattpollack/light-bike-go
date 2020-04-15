package controller

import (
	"../game"
)

// Always tries to turn left
type AlwaysTurnLeftController struct{}

func NewAlwaysTurnLeftController() *AlwaysTurnLeftController {
	return &AlwaysTurnLeftController{}
}

func (c *AlwaysTurnLeftController) Reset() {}

func (c *AlwaysTurnLeftController) Update(state game.State) game.PlayerMove {
	move := state.PreviousMove

	if state.PreviousMove == game.MoveNone {
		move = game.MoveNorth
	}

	for i := 0; i < 4; i++ {
		move = TurnLeft(move)

		if ValidMove(move, state) {
			break
		}
	}

	return move
}
