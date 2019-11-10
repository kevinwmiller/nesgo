package cpu6502

// Instruction represents a single 6502 instruction.
// Every instruction has a name, an implementation, an addressing mode, and a set number of cycle to run
type Instruction struct {
	Name           string
	Operation      func() uint8
	AddressingMode func() uint8
	Cycles         uint8
}
