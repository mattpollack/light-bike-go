package match

import (
	"fmt"

	"../game"
)

type Result struct {
	winRates []float64
}

const (
	TotalGames = 10000
)

func New(width int, height int, players []game.PlayerController) Result {
	wins := map[int]int{}
	games := 0
	var longestGame *game.Game = nil
	var shortestGame *game.Game = nil

	for ; games < TotalGames; games++ {
		for _, p := range players {
			p.Reset()
		}

		g := game.New(width, height, players)

		for {
			winners, err := g.Step()

			if err != nil {
				fmt.Printf("Error: %s\n", err)
				g.Draw()
				return Result{}
			}

			if len(winners) > 0 {
				for _, i := range winners {
					if _, ok := wins[i]; !ok {
						wins[i] = 0
					}

					wins[i]++
				}

				break
			}
		}

		if longestGame == nil {
			longestGame = g
		}

		if shortestGame == nil {
			shortestGame = g
		}

		if longestGame.StepsTaken() < g.StepsTaken() {
			longestGame = g
		}

		if shortestGame.StepsTaken() > g.StepsTaken() {
			shortestGame = g
		}
	}

	fmt.Println("--- LONGEST GAME ---")
	longestGame.Draw()
	fmt.Println("--- SHORTEST GAME ---")
	shortestGame.Draw()

	fmt.Println(wins)

	return Result{}
}
