package main

import (
	"nesgo/nes/bus"
	"nesgo/nes/clock"
	"nesgo/nes/cpu6502"
	"runtime"
)

const windowWidth = 1024
const windowHeight = 720

func init() {
	runtime.LockOSThread()
}

func main() {
	cpu := cpu6502.NewCPU()
	bus := bus.Bus{}
	cpu.ConnectBus(&bus)
	clock := clock.Clock{}
	clock.RegisterComponent(cpu, 3)

	// cpu.Dump()
	cpu.Reset()
	clock.Start()
}
