package cpu6502

// An AddressingMode handles how the data for the current operation is fetched
type AddressingMode struct {
	Name   string
	Lookup func() (uint8, uint16, bool)
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

// Undefined
func (c *CPU) und() (uint8, uint16, bool) {
	return 0, 0, false
}

// Accumulator
func (c *CPU) acc() (uint8, uint16, bool) {
	return c.A, 0, false
}

// Absolute
func (c *CPU) imm() (uint8, uint16, bool) {
	addr := c.PC
	data := c.Read(addr)
	c.PC++
	return data, addr, false
}

// Zero-page
func (c *CPU) zp0() (uint8, uint16, bool) {
	zpAddr := c.Read(c.PC)
	c.PC++
	data := c.Read(uint16(zpAddr))
	return data, uint16(zpAddr), false
}

// Zero-page X
func (c *CPU) zpx() (uint8, uint16, bool) {
	zpAddr := c.Read(c.PC) + c.X
	c.PC++
	data := c.Read(uint16(zpAddr))
	return data, uint16(zpAddr), false
}

// Zero-page Y
func (c *CPU) zpy() (uint8, uint16, bool) {
	zpAddr := c.Read(c.PC) + c.Y
	c.PC++
	data := c.Read(uint16(zpAddr))
	return data, uint16(zpAddr), false
}

// Absolute
func (c *CPU) abs() (uint8, uint16, bool) {
	hiByte := c.Read(c.PC)
	c.PC++
	lowByte := c.Read(c.PC)
	c.PC++
	addr := buildAddress(hiByte, lowByte)
	data := c.Read(addr)
	return data, addr, false
}

// Absolute X
func (c *CPU) abx() (uint8, uint16, bool) {
	return 0, 0, false
}

// Absolute Y
func (c *CPU) aby() (uint8, uint16, bool) {
	return 0, 0, false
}

// Implied
func (c *CPU) imp() (uint8, uint16, bool) {
	return 0, 0, false
}

// Relative
func (c *CPU) rel() (uint8, uint16, bool) {
	return 0, 0, false
}

// Indirect X
func (c *CPU) izx() (uint8, uint16, bool) {
	return 0, 0, false
}

// Indirect Y
func (c *CPU) izy() (uint8, uint16, bool) {
	return 0, 0, false
}

// Indirect
func (c *CPU) ind() (uint8, uint16, bool) {
	return 0, 0, false
}
