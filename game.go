package main

import (
	"fmt"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

const (
	sWidth = 640
	sHeight = 480
	cellSize = 20
)

type Direction int

const (
	Up Direction = iota
	Down
	Left
	Right
)

type Point struct {
	X int
	Y int
}

type Game struct {
	snake []Point
	direction Direction
	food Point
	gameover bool
	tick int
}

func NewGame() *Game {
	g := &Game{
		snake : []Point{{5,5}, {4,5}, {3,5}},
		direction: Right,
	}
	g.food = SpawnFood(g.snake)
	return g
}

func (g *Game) moveSnake() {
	head := g.snake[0]

	switch g.direction {
	case Up:
		head.Y--
	case Down:
		head.Y++
	case Left:
		head.X--
	case Right:
		head.X++
	}

	g.snake = append([]Point{head}, g.snake...)
	g.snake = g.snake[:len(g.snake)-1]
}

func(g *Game) Update() error {
	if g.gameover {
		return nil
	}

	HandleInput(g)

	g.tick++
	if g.tick%8 != 0 {
		return nil
	}

	g.moveSnake()
	CheckCollision(g)

	if inpututil.IsKeyJustPressed(ebiten.KeyQ) || inpututil.IsKeyJustPressed(ebiten.KeyEscape) {
		return ebiten.Termination
	}
	fmt.Println("Tick:", g.tick, "Head:", g.snake[0], "GameOver:", g.gameover)

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(color.Black)

	// snake
	for _, s := range g.snake {
		ebitenutil.DrawRect(
			screen,
			float64(s.X * cellSize),
			float64(s.Y * cellSize),
			cellSize,
			cellSize,
			color.RGBA{0,255,0,255},
		)
	}

	//food
	ebitenutil.DrawRect(
		screen,
		float64(g.food.X*cellSize),
		float64(g.food.Y*cellSize),
		cellSize,
		cellSize,
		color.RGBA{255, 0, 0, 255},
	)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return sWidth, sHeight
}