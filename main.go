package main

import (
	"image/color"
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
}

func (g *Game) Update() error {
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	print(ebiten.ActualFPS()) 

	for x := 0; x < cellsWide; x++ {
		for y := 0; y < cellsHigh; y++ {
			cellColor := color.White
			randInt := rand.Intn(2)
			
			if randInt == 1 {
				cellColor = color.Black
			}

			img := ebiten.NewImage(cellSize, cellSize)
			op := &ebiten.DrawImageOptions{}
			op.GeoM.Translate(float64(x*cellSize), float64(y*cellSize))
			img.Fill(cellColor)
			screen.DrawImage(img, op)
		}
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return windowWidth, windowHeight
}

func main() {
	ebiten.SetWindowSize(windowWidth, windowHeight)
	ebiten.SetWindowTitle("Hello, World!")
	if err := ebiten.RunGame(&Game{}); err != nil {
		log.Fatal(err)
	}
}