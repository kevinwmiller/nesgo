package main

import (
	"nesgo/nes/bus"
	"nesgo/nes/clock"
	"nesgo/nes/cpu6502"
)

func main() {
	cpu := cpu6502.NewCPU()
	bus := bus.Bus{}
	cpu.ConnectBus(&bus)
	clock := clock.Clock{}
	clock.RegisterComponent(cpu, 3)
	cpu.Dump()

	// runtime.LockOSThread()
	// if err := glfw.Init(); err != nil {
	// 	panic(fmt.Errorf("could not initialize glfw: %v", err))
	// }

	// glfw.WindowHint(glfw.ContextVersionMajor, 4)
	// glfw.WindowHint(glfw.ContextVersionMinor, 1)
	// glfw.WindowHint(glfw.Resizable, glfw.True)
	// glfw.WindowHint(glfw.OpenGLProfile, glfw.OpenGLCoreProfile)
	// glfw.WindowHint(glfw.OpenGLForwardCompatible, glfw.True)

	// win, err := glfw.CreateWindow(800, 600, "Hello world", nil, nil)if err != nil {
	// 	panic(fmt.Errorf("could not create opengl renderer: %v", err))
	// }

	// win.MakeContextCurrent()

	// for !win.ShouldClose() {
	// 	win.SwapBuffers()
	// 	glfw.PollEvents()
	//  }

	// clock.Start()
}
