package main

import (
	"nesgo/nes/bus"
	"nesgo/nes/cpu6502"
)

func main() {
	cpu := cpu6502.NewCPU()
	bus := bus.Bus{}
	cpu.ConnectBus(&bus)
}
