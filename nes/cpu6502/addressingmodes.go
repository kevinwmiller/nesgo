package cpu6502

//  uint8_t IMP();	uint8_t IMM();
// 	uint8_t ZP0();	uint8_t ZPX();
// 	uint8_t ZPY();	uint8_t REL();
// 	uint8_t ABS();	uint8_t ABX();
// 	uint8_t ABY();	uint8_t IND();
// 	uint8_t IZX();	uint8_t IZY();

func (c *CPU) IMP() (uint16, bool, uint8) {
	return 0, false, 0
}

// IMM df
func (c *CPU) IMM() (uint16, bool, uint8) {
	return c.PC, false, 0
}

func (c *CPU) ABS() (uint16, bool, uint8) {
	return 0, false, 0
}

func (c *CPU) ZPA() (uint16, bool, uint8) {
	return 0, false, 0
}

func (c *CPU) ZPX() (uint16, bool, uint8) {
	return 0, false, 0
}

func (c *CPU) ZPY() (uint16, bool, uint8) {
	return 0, false, 0
}

func (c *CPU) ACC() (uint16, bool, uint8) {
	return 0, false, 0
}

func (c *CPU) IND() (uint16, bool, uint8) {
	return 0, false, 0
}

func (c *CPU) PRE() (uint16, bool, uint8) {
	return 0, false, 0
}

func (c *CPU) PST() (uint16, bool, uint8) {
	return 0, false, 0
}

func (c *CPU) REL() (uint16, bool, uint8) {
	return 0, false, 0
}
