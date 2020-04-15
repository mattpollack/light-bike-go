package controller

import (
	"../game"
)

type MK_1 struct{}

func NewMK_1() *MK_1 {
	return &MK_1{}
}

func (c *MK_1) Reset() {}

func (c *MK_1) Update(state game.State) game.PlayerMove {
	return game.MoveNorth
}
