package main

import (
	"math/rand"
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

func GrowSnake(g *Game) {
	tail := g.snake[len(g.snake) - 1]
	g.snake = append(g.snake, tail)
}