package controller

import (
	"math/rand"
	"time"

	"../game"
)

// Picks a random move that won't kill it
type EagerRandomController struct {
	seededRand *rand.Rand
}

func NewEagerRandomController() *EagerRandomController {
	return &EagerRandomController{
		seededRand: rand.New(rand.NewSource(time.Now().UnixNano())),
	}
}

func (c *EagerRandomController) Reset() {}
func (c *EagerRandomController) Update(state game.State) game.PlayerMove {
	moves := []game.PlayerMove{}

	if ValidMove(game.MoveEast, state) {
		moves = append(moves, game.MoveEast)
	}

	if ValidMove(game.MoveWest, state) {
		moves = append(moves, game.MoveWest)
	}

	if ValidMove(game.MoveNorth, state) {
		moves = append(moves, game.MoveNorth)
	}

	if ValidMove(game.MoveSouth, state) {
		moves = append(moves, game.MoveSouth)
	}

	if len(moves) == 0 {
		return state.PreviousMove
	}

	return moves[c.seededRand.Int()%len(moves)]
}
