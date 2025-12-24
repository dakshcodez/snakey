package main

import (
	"math/rand"
	//"image/color"

	// "github.com/hajimehoshi/ebiten/v2"
	// "github.com/hajimehoshi/ebiten/v2/ebitenutil"
	// "github.com/hajimehoshi/ebiten/v2/inpututil"
)

func SpawnFood(snake []Point) Point {
	for {
		p := Point{
			X : rand.Intn(sWidth / cellSize),
			Y : rand.Intn(sHeight / cellSize),
		}

		valid := true
		for _, s := range snake {
			if s == p {
				valid = false
				break
			}
		}

		if valid{
			return p
		}
	}
}