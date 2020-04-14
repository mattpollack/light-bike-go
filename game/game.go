package game

import (
	"fmt"
	"math/rand"
	"time"
)

type Game struct {
	arena      Arena
	step       int
	players    []*player
	seededRand *rand.Rand
}

func New(width int, height int, playerCtrls []PlayerController) *Game {
	seededRand := rand.New(rand.NewSource(time.Now().UnixNano()))
	arena := Arena{
		width:  width,
		height: height,
		blocks: make([]int, width*height),
	}

	players := []*player{}

	for i, playerCtrl := range playerCtrls {
		p := &player{controller: playerCtrl}
		p.Reset()

		// Non-deterministic player placement
		for placed := false; !placed; {
			p.x = seededRand.Int() % width
			p.y = seededRand.Int() % height
			placed = true

			for _, p2 := range players {
				if p.x == p2.x && p.y == p2.y {
					placed = false
					break
				}
			}
		}

		arena.set(i+1, p.x, p.y)

		players = append(players, p)
	}

	return &Game{
		arena:      arena,
		step:       0,
		players:    players,
		seededRand: seededRand,
	}
}

type PlayerMove int

const (
	MoveNorth = PlayerMove(iota)
	MoveEast
	MoveSouth
	MoveWest
	MoveNone
)

type PlayerController interface {
	Update(State) PlayerMove
}

type player struct {
	controller PlayerController
	move       PlayerMove
	x          int
	y          int
	dead       bool
}

func (p *player) Reset() {
	p.move = MoveNone
	p.x = 0
	p.y = 0
	p.dead = false
}

type Arena struct {
	width  int
	height int
	blocks []int
}

func (a *Arena) Width() int {
	return a.width
}

func (a *Arena) Height() int {
	return a.height
}

func (a *Arena) set(v int, x int, y int) bool {
	if x < 0 || y < 0 || x >= a.width || y >= a.height {
		return false
	}

	a.blocks[x+y*a.width] = v

	return true
}

func (a *Arena) Get(x int, y int) int {
	if x < 0 || y < 0 || x >= a.width || y >= a.height {
		return -1
	}

	return a.blocks[x+y*a.width]
}

func (a *Arena) draw() {
	for x := 0; x < a.width+2; x++ {
		fmt.Print("#")
	}

	fmt.Println()

	for i, block := range a.blocks {
		x := i % a.width

		if x == 0 {
			fmt.Print("#")
		}

		switch block {
		case 0:
			fmt.Print(" ")
		default:
			fmt.Print(block - 1)
		}

		if x == a.width-1 {
			fmt.Println("#")
		}
	}

	for x := 0; x < a.width+2; x++ {
		fmt.Print("#")
	}

	fmt.Println()
}

type State struct {
	Arena        Arena
	X            int
	Y            int
	PreviousMove PlayerMove
}

func (g *Game) state(p int) State {
	return State{
		Arena:        g.arena,
		X:            g.players[p].x,
		Y:            g.players[p].y,
		PreviousMove: g.players[p].move,
	}
}

func (g *Game) Step() ([]int, error) {
	defer func() { g.step++ }()

	alivePlayers := []int{}

	for i, p := range g.players {
		if !p.dead {
			alivePlayers = append(alivePlayers, i)
		}
	}

	// Sanity check
	if len(alivePlayers) == 0 {
		return []int{}, fmt.Errorf("Somehow there isn't a winner?")
	}

	// Update players
	for _, i := range alivePlayers {
		p := g.players[i]

		switch move := p.controller.Update(g.state(i)); move {
		case MoveNorth:
			if p.move != MoveSouth {
				p.y++
				p.move = move
			}

		case MoveSouth:
			if p.move != MoveNorth {
				p.y--
				p.move = move
			}

		case MoveEast:
			if p.move != MoveWest {
				p.x++
				p.move = move
			}

		case MoveWest:
			if p.move != MoveEast {
				p.x--
				p.move = move
			}

		default:
			// Kill the player if it tries an invalid move
			p.dead = true
		}
	}

	// Move players and collide with obstacles
	for _, i := range alivePlayers {
		p := g.players[i]

		if v := g.arena.Get(p.x, p.y); v != 0 || !g.arena.set(i+1, p.x, p.y) {
			p.dead = true
		}
	}

	// Check for ties
	deadPlayers := 0

	for _, i := range alivePlayers {
		p1 := g.players[i]

		for _, j := range alivePlayers {
			p2 := g.players[j]

			if i == j {
				continue
			}

			if p1.x == p2.x && p1.y == p2.y {
				p1.dead = true
				p2.dead = true

				break
			}
		}

		if p1.dead {
			deadPlayers++
		}
	}

	// Its a tie
	if deadPlayers == len(alivePlayers) {
		return alivePlayers, nil
	}

	// There's a winner
	if deadPlayers == len(alivePlayers)-1 {
		for i, p := range g.players {
			if !p.dead {
				return []int{i}, nil
			}
		}
	}

	return []int{}, nil
}

func (g *Game) Draw() {
	fmt.Printf("Step %d\n", g.step)

	g.arena.draw()
}
