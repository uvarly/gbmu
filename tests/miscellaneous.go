package tests

import (
	"goboy/cpu"
	"goboy/memory"
	"math/rand"
	"testing"
	"time"
)

func Equals(t *testing.T, expected interface{}, got interface{}, errorf string) {
	if expected != got {
		t.Errorf(errorf, expected, got)
	}
}

func InitCPU() cpu.CPU {
	c := cpu.NewCPU()

	// d := InitDisplay()
	// j := InitJoypad()
	m := InitMemory()
	// s := InitSound()

	// c.BindDisplay(d)
	// c.BindJoypad(j)
	c.SetMemory(m)
	// c.BindSound(s)

	return c
}

// func InitDisplay() *display.Display {
// 	return nil
// }

// func InitJoypad() *joypad.Joypad {
// 	return nil
// }

func InitMemory() memory.Memory {
	m := memory.NewMemory()
	slice := RandSlice(0x10000)
	for i := range slice {
		m[i] = uint8(slice[i])
	}
	return m
}

// func InitSound() *sound.Sound {
// 	return nil
// }

func RandSlice(n int) []int {
	rand.Seed(time.Now().UTC().UnixNano())
	slice := make([]int, n)
	for i := range slice {
		slice[i] = rand.Int()
	}
	return slice
}

func RandRegisters(c *cpu.CPU) {
	rand.Seed(time.Now().UTC().UnixNano())
	c.Registers.SetA(uint8(rand.Int()))
	c.Registers.SetF(uint8(rand.Int()))
	c.Registers.SetB(uint8(rand.Int()))
	c.Registers.SetC(uint8(rand.Int()))
	c.Registers.SetD(uint8(rand.Int()))
	c.Registers.SetE(uint8(rand.Int()))
	c.Registers.SetH(uint8(rand.Int()))
	c.Registers.SetL(uint8(rand.Int()))
	c.Registers.SetSP(uint16(rand.Int()))
	c.Registers.SetPC(uint16(rand.Int()))
	c.Flags.SetFlagsFromValue(c.Registers.GetA())
}
