package controller

import (
	"../game"
)

func TurnLeft(move game.PlayerMove) game.PlayerMove {
	switch move {
	case game.MoveNorth:
		return game.MoveWest
	case game.MoveWest:
		return game.MoveSouth
	case game.MoveSouth:
		return game.MoveEast
	case game.MoveEast:
		return game.MoveNorth
	}

	return game.MoveNone
}

func TurnRight(move game.PlayerMove) game.PlayerMove {
	switch move {
	case game.MoveNorth:
		return game.MoveEast
	case game.MoveWest:
		return game.MoveNorth
	case game.MoveSouth:
		return game.MoveWest
	case game.MoveEast:
		return game.MoveSouth
	}

	return game.MoveNone
}

func ValidMove(move game.PlayerMove, state game.State) bool {
	switch move {
	case game.MoveNorth:
		return state.Arena.Get(state.X, state.Y+1) == 0 && state.PreviousMove != game.MoveSouth
	case game.MoveWest:
		return state.Arena.Get(state.X-1, state.Y) == 0 && state.PreviousMove != game.MoveEast
	case game.MoveSouth:
		return state.Arena.Get(state.X, state.Y-1) == 0 && state.PreviousMove != game.MoveNorth
	case game.MoveEast:
		return state.Arena.Get(state.X+1, state.Y) == 0 && state.PreviousMove != game.MoveWest
	}

	return false
}
