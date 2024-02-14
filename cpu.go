package main

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

func (c *CPU) ExecuteInstruction(instruction uint16) {
	c.pc += 2

	x := (instruction & 0x0F00) >> 8
	y := (instruction & 0x00F0) >> 4
	n := instruction & 0x000F
	nn := instruction & 0x00FF
	nnn := instruction & 0x0FFF

	switch instruction & 0xF000 {
		case 0x0000:
			switch instruction {
				case 0x00E0:
					//CLS - Clear the display
				case 0x00EE:
					//RET - Return from a subroutine
			}
		case 0x1000:
			//JP addr - Jump to location nnn
		case 0x2000:
			//CALL addr - Call subroutine at nnn
		case 0x3000:
			//SE Vx, byte - Skip next instruction if Vx = nn
		case 0x4000:
			//SNE Vx, byte - Skip next instruction if Vx != nn
		case 0x5000:
			//SE Vx, Vy - Skip next instruction if Vx = Vy
		case 0x6000:
			//LD Vx, byte - Set Vx = nn
		case 0x7000:
			//ADD Vx, byte - Set Vx = Vx + nn
		case 0x8000:
			//TODO: Check the rest of the 8xxx instructions
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
		case 0xF000:
			//TODO: Check the rest of the Fxxx instructions
		default:
			panic("INVALID OPCODE")
	}

}

