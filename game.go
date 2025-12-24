package main

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

const (
	sWidth = 1000
	sHeight = 800
	cellSize = 20
)

type Point struct {
	X int
	Y int
}

type Game struct {
	snake []Point
	food Point
}

func NewGame() *Game {
	g := &Game{
		snake : []Point{{5,5}, {4,5}, {3,5}},
	}
	g.food = SpawnFood(g.snake)
	return g
}

func(g *Game) Update() error {
	if inpututil.IsKeyJustPressed(ebiten.KeyQ) || inpututil.IsKeyJustPressed(ebiten.KeyEscape) {
		return ebiten.Termination
	}
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