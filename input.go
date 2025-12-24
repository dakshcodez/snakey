package main

import (
	"github.com/hajimehoshi/ebiten/v2"
)

func HandleInput(g * Game) {
	if ebiten.IsKeyPressed(ebiten.KeyArrowUp) && g.direction != Down {
		g.direction = Up
	}
	if ebiten.IsKeyPressed(ebiten.KeyArrowDown) && g.direction != Up {
		g.direction = Down
	}
	if ebiten.IsKeyPressed(ebiten.KeyArrowLeft) && g.direction != Right {
		g.direction = Left
	}
	if ebiten.IsKeyPressed(ebiten.KeyArrowRight) && g.direction != Left {
		g.direction = Right
	}
}