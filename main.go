package main

import (
	"image"
	"log"
	"github.com/hajimehoshi/ebiten/v2"
)

const (
	cellScale = 16
	cellsWide = 64
	cellsHigh = 32
	windowWidth  = cellsWide * cellScale
	windowHeight = cellsHigh * cellScale
)

type Game struct {
	display [cellsHigh][cellsWide]bool
	displayImage *image.RGBA
}

func (g *Game) Update() error {
	
	for row := 0; row < cellsHigh; row++ {
		for col := 0; col < cellsWide; col++ {

			imagePixelIndex := getDisplayImageIndex(row, col)

			c := uint8(0x0)
			if g.display[row][col] { 
				c = 0xff
			}

			//set rgba value of corresponding pixel in displayImage
			g.displayImage.Pix[imagePixelIndex*4] = c //R
			g.displayImage.Pix[imagePixelIndex*4+1] = c //G
			g.displayImage.Pix[imagePixelIndex*4+2] = c //B
			g.displayImage.Pix[imagePixelIndex*4+3] = 0xff //A
		}
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
		display: [cellsHigh][cellsWide]bool{},
		displayImage: image.NewRGBA(image.Rect(0, 0, cellsWide, cellsHigh)),
	}

	g.display[0][10] = true
	g.display[cellsHigh - 1][10] = true

	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}

func getDisplayImageIndex(row, col int) int {
	return row*cellsWide + col
}