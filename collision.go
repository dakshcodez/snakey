package main

import "fmt"

func CheckCollision(g *Game) {
	head := g.snake[0]

	if head.X < 0 || head.Y < 0 ||
		head.X >= sWidth/cellSize ||
		head.Y >= sHeight/cellSize {
			g.gameover = true
			return
	}
	
	for _, s := range g.snake[1:] {
		if s == head {
			g.gameover = true
			return
		}
	}

	if head == g.food {
		GrowSnake(g)
		g.food = SpawnFood(g.snake)
		g.point++
		fmt.Println("Points:", g.point)
	}
}