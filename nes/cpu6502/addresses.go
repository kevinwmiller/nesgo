package cpu6502

func buildAddress(hiByte, lowByte uint8) uint16 {
	return uint16(hiByte)<<8 + uint16(lowByte)
}

// SetProgramCounter sets the CPU program counter given two bytes that make up an address
func (c *CPU) SetProgramCounter(hiByte, lowByte uint8) {
	c.PC = buildAddress(hiByte, lowByte)
}
