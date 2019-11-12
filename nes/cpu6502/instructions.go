package cpu6502

// Instruction represents a single 6502 instruction.
// Every instruction has a name, an implementation, an addressing mode, and a set number of cycle to run
type Instruction struct {
	Name           string
	Execute        func() uint8
	AddressingMode func() uint8
	Cycles         uint8
}

func (c *CPU) buildInstructionTable() [256]Instruction {
	return [256]Instruction{
		{"BRK", c.brk, c.imp, 7}, {"ORA", c.ora, c.izx, 6}, {"UND", c.und, c.imp, 0}, {"UND", c.und, c.imp, 0}, {"ORA", c.ora, c.zp0, 3},
		{"ASL", c.asl, c.zp0, 5}, {"UND", c.und, c.imp, 0}, {"PHP", c.php, c.imp, 3}, {"ORA", c.ora, c.imm, 2}, {"ASL", c.asl, c.acc, 2},
		{"UND", c.und, c.imp, 0}, {"UND", c.und, c.imp, 0}, {"ORA", c.ora, c.abs, 3}, {"ORA", c.ora, c.abs, 4}, {"ASL", c.asl, c.abs, 6},
		{"BPL", c.bpl, c.rel, 2}, {"ORA", c.ora, c.izy, 5}, {"UND", c.und, c.imp, 0}, {"UND", c.und, c.imp, 0}, {"UND", c.und, c.imp, 0},
		{"ORA", c.ora, c.zpx, 4}, {"ASL", c.asl, c.zpx, 6}, {"UND", c.und, c.imp, 0}, {"CLC", c.clc, c.imp, 2}, {"ORA", c.ora, c.aby, 4},
		{"JSR", c.jsr, c.abs, 6}, //21
	}
}

func (c *CPU) und() uint8 {
	return 0
}

func (c *CPU) brk() uint8 {
	return 0
}

func (c *CPU) ora() uint8 {
	return 0
}

func (c *CPU) asl() uint8 {
	return 0
}

func (c *CPU) php() uint8 {
	return 0
}

func (c *CPU) bpl() uint8 {
	return 0
}

func (c *CPU) clc() uint8 {
	return 0
}

func (c *CPU) jsr() uint8 {
	return 0
}
