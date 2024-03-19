package main

import "fmt"

type CPU struct {
	memory [4096]uint8 // memory size 4k
	vx     [16]uint8   // cpu registers V0-VF
	key    [16]uint8   // input key
	stack  [16]uint16  // program counter stack
	oc     uint16 // current opcode
	pc     uint16 // program counter
	sp     uint16 // stack pointer
	iv     uint16 // index register
}

func NewCPU() *CPU {
	return &CPU{
		memory: [4096]uint8{}, 
		vx:     [16]uint8{},
		key:    [16]uint8{},
		stack:  [16]uint16{},
		oc:     0,
		pc:     0x200,
		sp:     0,
		iv:     0,
	}
}

func (c *CPU) ExecuteInstruction(instruction uint16, display *Display) {
	c.nextInstruction()

	x := (instruction & 0x0F00) >> 8
	//y := (instruction & 0x00F0) >> 4
	//n := instruction & 0x000F
	nn := instruction & 0x00FF
	//nnn := instruction & 0x0FFF

	fmt.Println("Executing instruction: ", instruction)

	switch instruction & 0xF000 {
		case 0x0000:
			switch instruction {
				case 0x00E0:
					//CLS - Clear the display
					display.Clear()
				case 0x00EE:
					//RET - Return from a subroutine
					c.pc = c.stack[c.sp]
					c.sp--
				default:
					panic("INVALID OPCODE")
			}
		case 0x1000:
			//JP addr - Jump to location nnn
			c.pc = instruction & 0x0FFF
		case 0x2000:
			//CALL addr - Call subroutine at nnn
			c.sp++

			if c.sp == 16 {
				panic("STACK OVERFLOW")
			}

			c.stack[c.sp] = c.pc
			c.pc = instruction & 0x0FFF
		case 0x3000:
			//SE Vx, byte - Skip next instruction if Vx = nn
			if c.vx[x] == uint8(nn) {
				c.pc += 2
			}
		case 0x4000:
			//SNE Vx, byte - Skip next instruction if Vx != nn
			if c.vx[x] != uint8(nn) {
				c.pc += 2
			}
		case 0x5000:
			//SE Vx, Vy - Skip next instruction if Vx = Vy
		case 0x6000:
			//LD Vx, byte - Set Vx = nn
		case 0x7000:
			//ADD Vx, byte - Set Vx = Vx + nn
		case 0x8000:
			switch instruction & 0xF {
				case 0x0:
					//LD, Vx, VY
				case 0x1:
					//OR Vx, Vy
				case 0x2:
					//AND Vx, Vy
				case 0x3:
					//XOR Vx, Vy
				case 0x4:
					//ADD Vx, Vy
				case 0x5:
					//SUB Vx, Vy
				case 0x6:
					//SHR Vx {, Vy}
				case 0x7:
					//SUBN Vx, Vy
				case 0xE:
					//SHL Vx {, Vy}
			}
		case 0x9000:
			//SNE Vx, Vy - Skip next instruction if Vx != Vy
		case 0xA000:
			//LD I, addr - Set I = nnn
		case 0xB000:
			//JP V0, addr - Jump to location nnn + V0
		case 0xC000:
			//RND Vx, byte - Set Vx = random byte AND nn
		case 0xD000:
			//DRW Vx, Vy, nibble - Display n-byte sprite starting at memory location I at (Vx, Vy), set VF = collision
		case 0xE000:
			//TODO: Check the rest of the Exxx instructions
			switch instruction & 0xFF {
				case 0x9E:
					//SKP Vx - Skip next instruction if key with the value of Vx is pressed
				case 0xA1:
					//SKNP Vx - Skip next instruction if key with the value of Vx is not pressed
			}
		case 0xF000:
			//TODO: Check the rest of the Fxxx instructions
			switch instruction & 0xFF {
				case 0x07:
					//LD Vx, DT - Set Vx = delay timer value
				case 0x0A:
					//LD Vx, K - Wait for a key press, store the value of the key in Vx
				case 0x15:
					//LD DT, Vx - Set delay timer = Vx
				case 0x18:
					//LD ST, Vx - Set sound timer = Vx
				case 0x1E:
					//ADD I, Vx - Set I = I + Vx
				case 0x29:
					//LD F, Vx - Set I = location of sprite for digit Vx
				case 0x33:
					//LD B, Vx - Store BCD representation of Vx in memory locations I, I+1, and I+2
				case 0x55:
					//LD [I], Vx - Store registers V0 through Vx in memory starting at location I
				case 0x65:
					//LD Vx, [I] - Read registers V0 through Vx from memory starting at location I
			}
		default:
			panic("INVALID OPCODE")
	}

}

func (c *CPU) nextInstruction() {
	c.pc += 2
}

