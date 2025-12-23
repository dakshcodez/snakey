package main

import (
	//"fmt"
	"log"
	//"math"

	"github.com/hajimehoshi/ebiten/v2"
	//"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	//"github.com/hajimehoshi/ebiten/v2/inpututil"
)

const (
	sWidth = 1000
	sHeight = 800
)

type Game struct {}

func(g *Game) Update() error {

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {

}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return sWidth, sHeight
}

func main() {
	ebiten.SetWindowSize(sWidth, sHeight)
	ebiten.SetWindowTitle("snakey")
	game := Game{}
	
	if err := ebiten.RunGame(&game); err != nil {
		log.Fatal(err)
	}
}
