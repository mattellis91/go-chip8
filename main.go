package main

import (
	"log"
	"os"
	"github.com/hajimehoshi/ebiten/v2"
)

const (
	PcStartAddress = 0x200
	FontStartAddress = 0x50
)

type Game struct {
	display Display
	cpu CPU
	speed int
	paused bool
}

func (g *Game) Update() error {

	for row := 0; row < cellsHigh; row++ {
		for col := 0; col < cellsWide; col++ {

			imagePixelIndex := getDisplayImageIndex(row, col)

			c := uint8(0x0)
			if g.display.displayCells[row][col] == 1{ 
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

	g.display.displayCells[0][10] = 1
	g.display.displayCells[cellsHigh - 1][10] = 1

	g.loadFontSet()
	g.loadRom("test_opcode.ch8")

	
	//test clear display
	g.cpu.ExecuteInstruction(0x00E0, &g.display)

	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}

func getDisplayImageIndex(row, col int) int {
	return row*cellsWide + col
}

func (g *Game) loadFontSet() {
	for i := 0; i < len(fontSet); i++ {
		g.cpu.memory[FontStartAddress + i] = fontSet[i]
	}
}

func (g *Game) loadRom(filePath string) error {
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
	for i := 0; i < len(buffer); i++ {
		g.cpu.memory[PcStartAddress + i] = buffer[i]
	}

	return nil
}

func (g *Game) cycle() {
	for i:=0; i < g.speed; i++ {
		opcode := (uint16(g.cpu.memory[g.cpu.pc])) << 8 | uint16(g.cpu.memory[g.cpu.pc + 1])
		g.cpu.ExecuteInstruction(uint16(opcode), &g.display)
	}
	if !g.paused {
		g.updateTimers()
	}
}

func (g *Game) updateTimers() {
	// if g.cpu.delayTimer > 0 {
	// 	g.cpu.delayTimer--
	// }
	// if g.cpu.soundTimer > 0 {
	// 	g.cpu.soundTimer--
	// }
}
