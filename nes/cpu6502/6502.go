package cpu6502

import (
	"fmt"
	"github.com/kevinwmiller/nesgo/nes/bus"
	"github.com/kevinwmiller/nesgo/nes/cpu6502/flags"
)

// CPU represents an instance of a 2A03 chip which is based on the 6502 processor with the exception of BCD instructions
type CPU struct {
	// Register descriptions from nesdev.com - http://nesdev.com/6502.txt and http://nesdev.com/6502_cpu.txt

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

	cycles          uint8
	instructions    [256]Instruction
	addressingModes [14]AddressingMode

	bus *bus.Bus
}

// NewCPU returns a new CPU object with all flags initialized.
func NewCPU() *CPU {
	cpu := CPU{}
	// The unused flag should be set at all times
	cpu.instructions = cpu.buildInstructionTable()
	cpu.addressingModes = cpu.buildAddressingModeTable()
	return &cpu
}

func (c *CPU) Write(address uint16, data uint8) {
	if c.bus == nil {
		return
	}
	c.bus.Write(address, data)
}

func (c *CPU) Read(address uint16) uint8 {
	if c.bus == nil {
		return 0
	}
	return c.bus.Read(address)
}

// Reset resets the CPU to a known state.
// The program counter starting address is loaded from 0xFFFC and 0xFFFD
// The registers and flags (with the exception of the unused flag) are initialized to 0
func (c *CPU) Reset() {
	// The initial state of the program counter can be found at 0xFFFC and 0xFFFD
	addr := uint16(0xFFFC)
	lowByte := c.Read(addr)
	hiByte := c.Read(addr + 1)
	c.PC = buildAddress(hiByte, lowByte)
	c.Status = flags.SetFlag(0x00, flags.U)
	c.A = 0
	c.X = 0
	c.Y = 0

	// http://wiki.nesdev.com/w/index.php/CPU_power_up_state
	c.SP = 0xFD

	c.cycles = 6
}

// Tick executes a single fetch/decode/execute cycle
func (c *CPU) Tick() {
	fmt.Println("Ticking")
	if c.cycles == 0 {
		c.Dump()
		opcode := c.fetch()
		instruction := &c.instructions[opcode]
		addressingMode := &c.addressingModes[instruction.AddressingMode]
		c.cycles = instruction.Cycles

		data, address, pageBoundaryCrossed := addressingMode.Lookup()
		if pageBoundaryCrossed {
			c.cycles += instruction.AdditionalCycles
		}
		c.cycles += instruction.Execute(data, address)
	}
	c.cycles--
}

// fetch fetches the next instruction
func (c *CPU) fetch() uint8 {
	opcode := c.Read(c.PC)
	c.PC++
	return opcode
}

// func (c *CPU) decode() *Instruction {
// 	return &instructions[]
// }

// ConnectBus connects a bus to the CPU
func (c *CPU) ConnectBus(bus *bus.Bus) {
	c.bus = bus
}

// Dump prints the CPU state to the console
func (c *CPU) Dump() {
	fmt.Printf("PC: %X\n", c.PC)
	fmt.Printf("A : %X\n", c.A)
	fmt.Printf("X : %X\n", c.X)
	fmt.Printf("Y : %X\n", c.Y)
	fmt.Printf("SP: %X\n", c.SP)
	fmt.Printf("Status: %08b\n", c.Status)
	fmt.Printf("Cycles: %d\n", c.cycles)

	// fmt.Println("Instructions: ")
	// for opcode, instruction := range c.instructions {
	// 	fmt.Printf("    %02X: %s - %-12s - %d cycles +%d*\n",
	// 		opcode,
	// 		instruction.Name,
	// 		c.addressingModes[instruction.AddressingMode].Name,
	// 		instruction.Cycles,
	// 		instruction.AdditionalCycles,
	// 	)
	// }
}
