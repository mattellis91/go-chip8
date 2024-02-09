package main

import (
	"image"
	"log"
	"math/rand"
	"github.com/hajimehoshi/ebiten/v2"
)

const (
	cellSize = 24
	cellsWide = 64
	cellsHigh = 32
	windowWidth  = cellsWide * cellSize
	windowHeight = cellsHigh * cellSize
)

type Game struct {
	display [64][32]bool
	displayImage *image.RGBA
}

func (g *Game) Update() error {
	l := cellsWide * cellsHigh

	for i := 0; i < l; i++ { 

		r := rand.Intn(2)
		c := uint8(0xff)
	
		if r == 1 {
			c = 0x0
		}

		g.displayImage.Pix[i*4] = c //R
		g.displayImage.Pix[i*4+1] = c //G
		g.displayImage.Pix[i*4+2] = c //B
		g.displayImage.Pix[i*4+3] = 0xff //A
	
	}
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.WritePixels(g.displayImage.Pix)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return cellsWide, cellsHigh
}

func main() {
	ebiten.SetWindowSize(windowWidth, windowHeight)
	ebiten.SetWindowTitle("Hello, World!")

	g := &Game{
		display: [64][32]bool{},
		displayImage: image.NewRGBA(image.Rect(0, 0, cellsWide, cellsHigh)),
	}

	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}