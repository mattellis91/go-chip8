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

