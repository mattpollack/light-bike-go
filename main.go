package main

import (
	"fmt"
	"math/rand"
	"time"

	"./game"
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

func (c *EagerRandomController) Update(state game.State) game.PlayerMove {
	moves := []game.PlayerMove{}

	if state.Arena.Get(state.X+1, state.Y) == 0 {
		moves = append(moves, game.MoveEast)
	}

	if state.Arena.Get(state.X-1, state.Y) == 0 {
		moves = append(moves, game.MoveWest)
	}

	if state.Arena.Get(state.X, state.Y+1) == 0 {
		moves = append(moves, game.MoveNorth)
	}

	if state.Arena.Get(state.X, state.Y-1) == 0 {
		moves = append(moves, game.MoveSouth)
	}

	if len(moves) == 0 {
		return state.PreviousMove
	}

	return moves[c.seededRand.Int()%len(moves)]
}

// Picks its previous move unless it would kill it
type LazyRandomController struct {
	seededRand *rand.Rand
}

func NewLazyRandomController() *LazyRandomController {
	return &LazyRandomController{
		seededRand: rand.New(rand.NewSource(time.Now().UnixNano())),
	}
}

func (c *LazyRandomController) Update(state game.State) game.PlayerMove {
	moves := []game.PlayerMove{}

	if state.Arena.Get(state.X+1, state.Y) == 0 {
		moves = append(moves, game.MoveEast)
	}

	if state.Arena.Get(state.X-1, state.Y) == 0 {
		moves = append(moves, game.MoveWest)
	}

	if state.Arena.Get(state.X, state.Y+1) == 0 {
		moves = append(moves, game.MoveNorth)
	}

	if state.Arena.Get(state.X, state.Y-1) == 0 {
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

func main() {
	fmt.Println("## light-bike-go")

	g := game.New(40, 20, []game.PlayerController{
		//NewEagerRandomController(),
		NewLazyRandomController(),
		NewLazyRandomController(),
		NewLazyRandomController(),
	})
	startTime := time.Now()

	g.Draw()

	for {
		time.Sleep(200 * time.Millisecond)

		winners, err := g.Step()

		if err != nil {
			fmt.Printf("Error: %s\n", err)
			return
		}

		g.Draw()

		if len(winners) > 0 {
			fmt.Println("The winners are: ")
			fmt.Println(winners)

			break
		}
	}

	fmt.Println(time.Now().Sub(startTime))
}
