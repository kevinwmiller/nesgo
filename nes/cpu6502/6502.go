package cpu6502

import (
	"nesgo/nes/bus"
	"nesgo/nes/cpu6502/flags"
)

// CPU represents an instance of a 2A03 chip which is based on the 6502 processor with the exception of BCD instructions
type CPU struct {
	// Register descriptions from nesdev.com - http://nesdev.com/6502.txt

	// THE ACCUMULATOR
	//   This is THE most important register in the microprocessor. Various ma-
	//   chine language instructions allow you to copy the contents of a memory
	//   location into the accumulator, copy the contents of the accumulator into
	//   a memory location, modify the contents of the accumulator or some other
	//   register directly, without affecting any memory. And the accumulator is
	//   the only register that has instructions for performing math.
	A uint8

	// THE X INDEX REGISTER
	//   This is a very important register. There are instructions for nearly
	//   all of the transformations you can make to the accumulator. But there are
	//   other instructions for things that only the X register can do. Various
	//   machine language instructions allow you to copy the contents of a memory
	//   location into the X register, copy the contents of the X register into a
	//   memory location, and modify the contents of the X, or some other register
	//   directly.
	X uint8

	// THE Y INDEX REGISTER
	//   This is a very important register. There are instructions for nearly
	//   all of the transformations you can make to the accumulator, and the X
	//   register. But there are other instructions for things that only the Y
	//   register can do. Various machine language instructions allow you to copy
	//   the contents of a memory location into the Y register, copy the contents
	//   of the Y register into a memory location, and modify the contents of the
	//   Y, or some other register directly.
	Y uint8

	// THE STACK POINTER
	// 	 This register contains the location of the first empty place on the
	//	 stack. The stack is used for temporary storage by machine language pro-
	// 	 grams, and by the computer.
	SP uint8

	// THE PROGRAM COUNTER
	// This contains the address of the current machine language instruction
	// being executed. Since the operating system is always "RUN"ning in the
	// Commodore VIC-20 (or, for that matter, any computer), the program counter
	// is always changing. It could only be stopped by halting the microprocessor
	// in some way.
	PC uint16

	// 	THE STATUS REGISTER
	//   This register consists of eight "flags" (a flag = something that indi-
	//   cates whether something has, or has not occurred). Bits of this register
	//   are altered depending on the result of arithmetic and logical operations.
	//   These bits are described below:
	Status uint8

	bus *bus.Bus
}

// NewCPU returns a new CPU object with all flags initialized.
func NewCPU() *CPU {
	cpu := CPU{}
	// The unused flag should be set at all times
	cpu.Status = flags.SetFlag(0x00, flags.U)
	return &cpu
}

func (c *CPU) Write(address uint16, data uint8) {
	c.bus.Write(address, data)
}

func (c *CPU) Read(address uint16) uint8 {
	return c.bus.Read(address)
}

func (c *CPU) fetch() {

}

func (c *CPU) ConnectBus(bus *bus.Bus) {
	c.bus = bus
}
