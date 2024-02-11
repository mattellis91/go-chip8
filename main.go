package main

import (
	"log"
	"os"

	"github.com/hajimehoshi/ebiten/v2"
)

type Game struct {
	display Display
	cpu CPU
}

func (g *Game) Update() error {

	if ebiten.IsKeyPressed(ebiten.KeyW) {
		println("W key pressed");
	} 

	if ebiten.IsKeyPressed(ebiten.KeyA){
		println("A key pressed");	
	}

	if ebiten.IsKeyPressed(ebiten.KeyS){
		println("S key pressed");	
	}

	if ebiten.IsKeyPressed(ebiten.KeyD){
		println("D key pressed");	
	}

	for row := 0; row < cellsHigh; row++ {
		for col := 0; col < cellsWide; col++ {

			imagePixelIndex := getDisplayImageIndex(row, col)

			c := uint8(0x0)
			if g.display.displayCells[row][col] { 
				c = 0xff
			}

			//set rgba value of corresponding pixel in displayImage
			g.display.displayImage.Pix[imagePixelIndex*4] = c //R
			g.display.displayImage.Pix[imagePixelIndex*4+1] = c //G
			g.display.displayImage.Pix[imagePixelIndex*4+2] = c //B
			g.display.displayImage.Pix[imagePixelIndex*4+3] = 0xff //A
		}
	}

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.WritePixels(g.display.displayImage.Pix)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return cellsWide, cellsHigh
}

func main() {
	ebiten.SetWindowSize(windowWidth, windowHeight)
	ebiten.SetWindowTitle("Hello, World!")

	g := &Game{
		display: *NewDisplay(),
		cpu: *NewCPU(),
	}

	g.display.displayCells[0][10] = true
	g.display.displayCells[cellsHigh - 1][10] = true

	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}

func getDisplayImageIndex(row, col int) int {
	return row*cellsWide + col
}

func loadRom(filePath string) error {
	file, err := os.OpenFile(filePath, os.O_RDONLY, 0)
	if err != nil {
		return err
	}
	defer file.Close()

	fStat, err := file.Stat()
	if err != nil {
		return err
	}

	buffer := make([]byte, fStat.Size())
	_, err = file.Read(buffer)
	if err != nil {
		return err
	}

	//TODO: load buffer into cpu memory

	return nil
}