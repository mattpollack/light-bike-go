package main

import (
	"fmt"
	"time"

	"./controller"
	"./game"
	"./match"
)

func main() {
	fmt.Println("## light-bike-go")
	startTime := time.Now()
	defer func() { fmt.Println(time.Now().Sub(startTime)) }()

	width := 60
	height := 30

	players := []game.PlayerController{
		controller.NewAlwaysTurnLeftController(),
		controller.NewLazyRandomController(),
		controller.NewEagerRandomController(),
		controller.NewMK_1(),
	}

	// ---------------------------------------

	if false {
		g := game.New(width, height, players)

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

		fmt.Println("\n------------------------\n")
	}

	t := match.New(width, height, players)

	fmt.Println(t)

}
