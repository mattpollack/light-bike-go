package controller

import (
	"math/rand"
	"time"

	"../game"
)

// Picks its previous move unless it would kill it
type LazyRandomController struct {
	seededRand *rand.Rand
}

func NewLazyRandomController() *LazyRandomController {
	return &LazyRandomController{
		seededRand: rand.New(rand.NewSource(time.Now().UnixNano())),
	}
}

func (c *LazyRandomController) Reset() {}
func (c *LazyRandomController) Update(state game.State) game.PlayerMove {
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

	for _, m := range moves {
		if m == state.PreviousMove {
			return state.PreviousMove
		}
	}

	if len(moves) == 0 {
		return state.PreviousMove
	}

	return moves[c.seededRand.Int()%len(moves)]
}
