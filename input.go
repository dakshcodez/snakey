package main

import (
	"github.com/hajimehoshi/ebiten/v2"
)

func HandleInput(g * Game) {
	if (ebiten.IsKeyPressed(ebiten.KeyArrowUp) || ebiten.IsKeyPressed(ebiten.KeyW)) && g.direction != Down {
		g.direction = Up
	}
	if (ebiten.IsKeyPressed(ebiten.KeyArrowDown) || ebiten.IsKeyPressed(ebiten.KeyS)) && g.direction != Up {
		g.direction = Down
	}
	if (ebiten.IsKeyPressed(ebiten.KeyArrowLeft) || ebiten.IsKeyPressed(ebiten.KeyA)) && g.direction != Right {
		g.direction = Left
	}
	if (ebiten.IsKeyPressed(ebiten.KeyArrowRight) || ebiten.IsKeyPressed(ebiten.KeyD)) && g.direction != Left {
		g.direction = Right
	}
}