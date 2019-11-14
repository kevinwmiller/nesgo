package cpu6502

// An AddressingMode handles how the data for the current operation is fetched
type AddressingMode struct {
	Name   string
	Lookup func() uint8
}

// Addressing Modes:
// Imp - Implied
// Acc - Accumulator
// Imm - Immediate
// Zp0 - Zero Page
// Zpx - Zero Page X
// Zpy - Zero Page Y
// Abs - Absolute
// Abx - Absolute X
// Aby - Absolute Y
// Rel - Relative
// Izx - Indirect X
// Izy - Indirect Y
// Ind - Indirect
const (
	Und uint8 = iota
	Imp
	Acc
	Imm
	Zp0
	Zpx
	Zpy
	Abs
	Abx
	Aby
	Rel
	Izx
	Izy
	Ind
)

func (c *CPU) buildAddressingModeTable() [14]AddressingMode {
	return [...]AddressingMode{
		{"Undefined", c.und},
		{"Implied", c.imp},
		{"Accumulator", c.acc},
		{"Immediate", c.imm},
		{"Zero Page", c.zp0},
		{"Zero Page X", c.zpx},
		{"Zero Page Y", c.zpy},
		{"Absolute", c.abs},
		{"Absolute X", c.abx},
		{"Absolute Y", c.aby},
		{"Relative", c.rel},
		{"Indirect X", c.izx},
		{"Indirect Y", c.izy},
		{"Indirect", c.ind},
	}
}

// Accumulator
func (c *CPU) acc() uint8 {
	return 0
}

// Absolute
func (c *CPU) imm() uint8 {
	return 0
}

// Zero-page
func (c *CPU) zp0() uint8 {
	return 0
}

// Zero-page X
func (c *CPU) zpx() uint8 {
	return 0
}

// Zero-page Y
func (c *CPU) zpy() uint8 {
	return 0
}

// Absolute
func (c *CPU) abs() uint8 {
	return 0
}

// Absolute X
func (c *CPU) abx() uint8 {
	return 0
}

// Absolute Y
func (c *CPU) aby() uint8 {
	return 0
}

// Implied
func (c *CPU) imp() uint8 {
	return 0
}

// Relative
func (c *CPU) rel() uint8 {
	return 0
}

// Indirect X
func (c *CPU) izx() uint8 {
	return 0
}

// Indirect Y
func (c *CPU) izy() uint8 {
	return 0
}

// Indirect
func (c *CPU) ind() uint8 {
	return 0
}
