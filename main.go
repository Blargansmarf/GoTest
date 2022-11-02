package main

import (
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type Game struct{}

var dir = "no dir"

func (g *Game) Update() error {
	if ebiten.IsKeyPressed(ebiten.KeyArrowRight) {
		dir = "right"
	} else if ebiten.IsKeyPressed(ebiten.KeyArrowLeft) {
		dir = "left"
	} else if ebiten.IsKeyPressed(ebiten.KeyArrowDown) {
		dir = "down"
	} else if ebiten.IsKeyPressed(ebiten.KeyArrowUp) {
		dir = "up"
	} else {
		dir = "no dir"
	}
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(color.RGBA{0x50, 0x50, 0x50, 0xFF})
	ebitenutil.DebugPrint(screen, dir)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 320, 240
}

func main() {
	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowTitle("Hello, World!")
	if err := ebiten.RunGame(&Game{}); err != nil {
		log.Fatal(err)
	}
}
