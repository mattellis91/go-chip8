package main

import (
	"github.com/hajimehoshi/ebiten/v2"
)

type Keyboard struct {
	key [16]uint8
	keyMap map[ebiten.Key]uint8
}


//INPUT KEYS

//Chip8 	//Keyboard
// 1 2 3 C  | 1 2 3 4
// 4 5 6 D  | Q W E R
// 7 8 9 E	| A S D F 
// A 0 B F	| Z X C V

func NewKeyboard() *Keyboard {
	return &Keyboard{
		key: [16]uint8{},
		keyMap: map[ebiten.Key]uint8{
			ebiten.Key1: 0x1,
			ebiten.Key2: 0x2,
			ebiten.Key3: 0x3,
			ebiten.Key4: 0xC,
			ebiten.KeyQ: 0x4,
			ebiten.KeyW: 0x5,
			ebiten.KeyE: 0x6,
			ebiten.KeyR: 0xD,
			ebiten.KeyA: 0x7,
			ebiten.KeyS: 0x8,
			ebiten.KeyD: 0x9,
			ebiten.KeyF: 0xE,
			ebiten.KeyZ: 0xA,
			ebiten.KeyX: 0x0,
			ebiten.KeyC: 0xB,
			ebiten.KeyV: 0xF,
		},
	}
}

func UpdateKeyboard() {
	
}