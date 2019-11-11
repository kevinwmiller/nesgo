package cpu6502

const ()

func (c *CPU) undefined() (uint16, bool, uint8) {
	return c.PC, false, 0
}

func (c *CPU) immediate() (uint16, bool, uint8) {
	return c.PC, false, 0
}

func (c *CPU) absolute() (uint16, bool, uint8) {
	return 0, false, 0
}

func (c *CPU) absoluteX() (uint16, bool, uint8) {
	return 0, false, 0
}

func (c *CPU) absoluteY() (uint16, bool, uint8) {
	return 0, false, 0
}

func (c *CPU) zeroPage() (uint16, bool, uint8) {
	return 0, false, 0
}

func (c *CPU) zeroPageX() (uint16, bool, uint8) {
	return 0, false, 0
}

func (c *CPU) zeroPageY() (uint16, bool, uint8) {
	return 0, false, 0
}

func (c *CPU) implied() (uint16, bool, uint8) {
	return 0, false, 0
}

func (c *CPU) accumulator() (uint16, bool, uint8) {
	return 0, false, 0
}

func (c *CPU) indexed() (uint16, bool, uint8) {
	return 0, false, 0
}

func (c *CPU) indirect() (uint16, bool, uint8) {
	return 0, false, 0
}

func (c *CPU) preIndexIndirect() (uint16, bool, uint8) {
	return 0, false, 0
}

func (c *CPU) postIndexIndirect() (uint16, bool, uint8) {
	return 0, false, 0
}

func (c *CPU) relative() (uint16, bool, uint8) {
	return 0, false, 0
}
