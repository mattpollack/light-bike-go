# Light Bike Go
Program a light bike driver to compete with others!

## Setup

Only need go and terminal

```bash
go run main.go
```

## Rules

Write a program that pilots a light-bike. Light-bikes are like regular bikes except they also leave a trail behind them. 

- You lose if you collide into a trail
- You lose if you collide into the arena
- You win if you're the last player standing

Please don't modify game state using unsafe and/or reflect

## Player Controller

Implement this interface:

```go
type PlayerController interface {
	Update(State) PlayerMove
}
```

