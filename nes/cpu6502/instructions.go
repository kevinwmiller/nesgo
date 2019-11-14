package cpu6502

// Instruction represents a single 6502 instruction.
// Every instruction has a name, an implementation, an addressing mode, and a set number of cycle to run
type Instruction struct {
	Name           string
	Execute        func() uint8
	AddressingMode uint8
	Cycles         uint8
}

func (c *CPU) buildInstructionTable() [256]Instruction {
	return [...]Instruction{
		{"BRK", c.brk, Imp, 7}, {"ORA", c.ora, Izx, 0}, {"UND", c.und, Und, 0}, {"UND", c.und, Und, 0}, {"UND", c.und, Und, 0},
		{"ORA", c.ora, Zp0, 0}, {"ASL", c.asl, Zp0, 0}, {"UND", c.und, Und, 0}, {"PHP", c.php, Imp, 0}, {"ORA", c.ora, Imm, 0},
		{"ASL", c.asl, Acc, 0}, {"UND", c.und, Und, 0}, {"UND", c.und, Und, 0}, {"ORA", c.ora, Abs, 0}, {"ASL", c.asl, Abs, 0},
		{"UND", c.und, Und, 0}, {"BPL", c.bpl, Imp, 0}, {"ORA", c.ora, Izy, 0}, {"UND", c.und, Und, 0}, {"UND", c.und, Und, 0},
		{"UND", c.und, Und, 0}, {"ORA", c.ora, Zpx, 0}, {"ASL", c.asl, Zpx, 0}, {"UND", c.und, Und, 0}, {"CLC", c.clc, Imp, 0},
		{"ORA", c.ora, Aby, 0}, {"UND", c.und, Und, 0}, {"UND", c.und, Und, 0}, {"UND", c.und, Und, 0}, {"ORA", c.ora, Abx, 0},
		{"ASL", c.asl, Abx, 0}, {"UND", c.und, Und, 0}, {"JSR", c.jsr, Imp, 0}, {"AND", c.and, Izx, 0}, {"UND", c.und, Und, 0},
		{"UND", c.und, Und, 0}, {"BIT", c.bit, Zp0, 0}, {"AND", c.and, Zp0, 0}, {"ROL", c.rol, Zp0, 0}, {"UND", c.und, Und, 0},
		{"PLP", c.plp, Imp, 0}, {"AND", c.and, Imm, 0}, {"ROL", c.rol, Acc, 0}, {"UND", c.und, Und, 0}, {"BIT", c.bit, Abs, 0},
		{"AND", c.and, Abs, 0}, {"ROL", c.rol, Abs, 0}, {"UND", c.und, Und, 0}, {"BMI", c.bmi, Imp, 0}, {"AND", c.and, Izy, 0},
		{"UND", c.und, Und, 0}, {"UND", c.und, Und, 0}, {"UND", c.und, Und, 0}, {"AND", c.and, Zpx, 0}, {"ROL", c.rol, Zpx, 0},
		{"UND", c.und, Und, 0}, {"SEC", c.sec, Imp, 0}, {"AND", c.and, Aby, 0}, {"UND", c.und, Und, 0}, {"UND", c.und, Und, 0},
		{"UND", c.und, Und, 0}, {"AND", c.and, Abx, 0}, {"ROL", c.rol, Abx, 0}, {"UND", c.und, Und, 0}, {"RTI", c.rti, Imp, 0},
		{"EOR", c.eor, Izx, 0}, {"UND", c.und, Und, 0}, {"UND", c.und, Und, 0}, {"UND", c.und, Und, 0}, {"EOR", c.eor, Zp0, 0},
		{"LSR", c.lsr, Zp0, 0}, {"UND", c.und, Und, 0}, {"PHA", c.pha, Imp, 0}, {"EOR", c.eor, Imm, 0}, {"LSR", c.lsr, Acc, 0},
		{"UND", c.und, Und, 0}, {"JMP", c.jmp, Abs, 0}, {"EOR", c.eor, Abs, 0}, {"LSR", c.lsr, Abs, 0}, {"UND", c.und, Und, 0},
		{"BVC", c.bvc, Imp, 0}, {"EOR", c.eor, Izy, 0}, {"UND", c.und, Und, 0}, {"UND", c.und, Und, 0}, {"UND", c.und, Und, 0},
		{"EOR", c.eor, Zpx, 0}, {"LSR", c.lsr, Zpx, 0}, {"UND", c.und, Und, 0}, {"CLI", c.cli, Imp, 0}, {"EOR", c.eor, Aby, 0},
		{"UND", c.und, Und, 0}, {"UND", c.und, Und, 0}, {"UND", c.und, Und, 0}, {"EOR", c.eor, Abx, 0}, {"LSR", c.lsr, Abx, 0},
		{"UND", c.und, Und, 0}, {"RTS", c.rts, Imp, 0}, {"ADC", c.adc, Izx, 6}, {"UND", c.und, Und, 0}, {"UND", c.und, Und, 0},
		{"UND", c.und, Und, 0}, {"ADC", c.adc, Zp0, 3}, {"ROR", c.ror, Zp0, 0}, {"UND", c.und, Und, 0}, {"PLA", c.pla, Imp, 0},
		{"ADC", c.adc, Imm, 2}, {"ROR", c.ror, Acc, 0}, {"UND", c.und, Und, 0}, {"JMP", c.jmp, Ind, 0}, {"ADC", c.adc, Abs, 0},
		{"ROR", c.ror, Abs, 0}, {"UND", c.und, Und, 0}, {"BVS", c.bvs, Imp, 0}, {"ADC", c.adc, Izy, 0}, {"UND", c.und, Und, 0},
		{"UND", c.und, Und, 0}, {"UND", c.und, Und, 0}, {"ADC", c.adc, Zpx, 0}, {"ROR", c.ror, Zpx, 0}, {"UND", c.und, Und, 0},
		{"SEI", c.sei, Imp, 0}, {"ADC", c.adc, Aby, 0}, {"UND", c.und, Und, 0}, {"UND", c.und, Und, 0}, {"UND", c.und, Und, 0},
		{"ADC", c.adc, Abx, 0}, {"ROR", c.ror, Abx, 0}, {"UND", c.und, Und, 0}, {"UND", c.und, Und, 0}, {"STA", c.sta, Izx, 0},
		{"UND", c.und, Und, 0}, {"UND", c.und, Und, 0}, {"STY", c.sty, Zp0, 0}, {"STA", c.sta, Zp0, 0}, {"STX", c.stx, Zp0, 0},
		{"UND", c.und, Und, 0}, {"DEY", c.dey, Imp, 0}, {"UND", c.und, Und, 0}, {"TXA", c.txa, Imp, 0}, {"UND", c.und, Und, 0},
		{"STY", c.sty, Abs, 0}, {"STA", c.sta, Abs, 0}, {"STX", c.stx, Abs, 0}, {"UND", c.und, Und, 0}, {"BCC", c.bcc, Imp, 0},
		{"STA", c.sta, Izy, 0}, {"UND", c.und, Und, 0}, {"UND", c.und, Und, 0}, {"STY", c.sty, Zpx, 0}, {"STA", c.sta, Zpx, 0},
		{"STX", c.stx, Zpy, 0}, {"UND", c.und, Und, 0}, {"TYA", c.tya, Imp, 0}, {"STA", c.sta, Aby, 0}, {"TXS", c.txs, Imp, 0},
		{"UND", c.und, Und, 0}, {"UND", c.und, Und, 0}, {"STA", c.sta, Abx, 0}, {"UND", c.und, Und, 0}, {"UND", c.und, Und, 0},
		{"LDY", c.ldy, Imm, 0}, {"LDA", c.lda, Izx, 0}, {"LDX", c.ldx, Imm, 0}, {"UND", c.und, Und, 0}, {"LDY", c.ldy, Zp0, 0},
		{"LDA", c.lda, Zp0, 0}, {"LDX", c.ldx, Zp0, 0}, {"UND", c.und, Und, 0}, {"TAY", c.tay, Imp, 0}, {"LDA", c.lda, Imm, 0},
		{"TAX", c.tax, Imp, 0}, {"UND", c.und, Und, 0}, {"LDY", c.ldy, Abs, 0}, {"LDA", c.lda, Abs, 0}, {"LDX", c.ldx, Abs, 0},
		{"UND", c.und, Und, 0}, {"BCS", c.bcs, Imp, 0}, {"LDA", c.lda, Izy, 0}, {"UND", c.und, Und, 0}, {"UND", c.und, Und, 0},
		{"LDY", c.ldy, Zpx, 0}, {"LDA", c.lda, Zpx, 0}, {"LDX", c.ldx, Zpy, 0}, {"UND", c.und, Und, 0}, {"CLV", c.clv, Imp, 0},
		{"LDA", c.lda, Aby, 0}, {"TSX", c.tsx, Imp, 0}, {"UND", c.und, Und, 0}, {"LDY", c.ldy, Abx, 0}, {"LDA", c.lda, Abx, 0},
		{"LDX", c.ldx, Aby, 0}, {"UND", c.und, Und, 0}, {"Cpy", c.cpy, Imm, 0}, {"CMP", c.cmp, Izx, 0}, {"UND", c.und, Und, 0},
		{"UND", c.und, Und, 0}, {"CPY", c.cpy, Zp0, 0}, {"CMP", c.cmp, Zp0, 0}, {"DEC", c.dec, Zp0, 0}, {"UND", c.und, Und, 0},
		{"INY", c.iny, Imp, 0}, {"CMP", c.cmp, Imm, 0}, {"DEX", c.dex, Imp, 0}, {"UND", c.und, Und, 0}, {"CPY", c.cpy, Abs, 0},
		{"CMP", c.cmp, Abs, 0}, {"DEC", c.dec, Abs, 0}, {"UND", c.und, Und, 0}, {"BNE", c.bne, Imp, 0}, {"CMP", c.cmp, Izy, 0},
		{"UND", c.und, Und, 0}, {"UND", c.und, Und, 0}, {"UND", c.und, Und, 0}, {"CMP", c.cmp, Zpx, 0}, {"DEC", c.dec, Zpx, 0},
		{"UND", c.und, Und, 0}, {"CLD", c.cld, Imp, 0}, {"CMP", c.cmp, Aby, 0}, {"UND", c.und, Und, 0}, {"UND", c.und, Und, 0},
		{"UND", c.und, Und, 0}, {"CMP", c.cmp, Abx, 0}, {"DEC", c.dec, Abx, 0}, {"UND", c.und, Und, 0}, {"CPX", c.cpx, Imm, 0},
		{"SBC", c.sbc, Izx, 0}, {"UND", c.und, Und, 0}, {"UND", c.und, Und, 0}, {"CPX", c.cpx, Zp0, 0}, {"SBC", c.sbc, Zp0, 0},
		{"INC", c.inc, Zp0, 0}, {"UND", c.und, Und, 0}, {"INX", c.inx, Imp, 0}, {"SBC", c.sbc, Imm, 0}, {"NOP", c.nop, Imp, 0},
		{"UND", c.und, Und, 0}, {"CPX", c.cpx, Abs, 0}, {"SBC", c.sbc, Abs, 0}, {"INC", c.inc, Abs, 0}, {"UND", c.und, Und, 0},
		{"BEQ", c.beq, Imp, 0}, {"SBC", c.sbc, Izy, 0}, {"UND", c.und, Und, 0}, {"UND", c.und, Und, 0}, {"UND", c.und, Und, 0},
		{"SBC", c.sbc, Zpx, 0}, {"INC", c.inc, Zpx, 0}, {"UND", c.und, Und, 0}, {"SED", c.sed, Imp, 0}, {"SBC", c.sbc, Aby, 0},
		{"UND", c.und, Und, 0}, {"UND", c.und, Und, 0}, {"UND", c.und, Und, 0}, {"SBC", c.sbc, Abx, 0}, {"INC", c.inc, Abx, 0},
		{"UND", c.und, Und, 0},
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

func (c *CPU) and() uint8 {
	return 0
}

func (c *CPU) bit() uint8 {
	return 0
}

func (c *CPU) rol() uint8 {
	return 0
}

func (c *CPU) plp() uint8 {
	return 0
}

func (c *CPU) bmi() uint8 {
	return 0
}

func (c *CPU) sec() uint8 {
	return 0
}

func (c *CPU) rti() uint8 {
	return 0
}

func (c *CPU) eor() uint8 {
	return 0
}

func (c *CPU) lsr() uint8 {
	return 0
}

func (c *CPU) pha() uint8 {
	return 0
}

func (c *CPU) jmp() uint8 {
	return 0
}

func (c *CPU) bvc() uint8 {
	return 0
}

func (c *CPU) cli() uint8 {
	return 0
}

func (c *CPU) rts() uint8 {
	return 0
}

func (c *CPU) adc() uint8 {
	return 0
}

func (c *CPU) ror() uint8 {
	return 0
}

func (c *CPU) pla() uint8 {
	return 0
}

func (c *CPU) bvs() uint8 {
	return 0
}

func (c *CPU) sei() uint8 {
	return 0
}

func (c *CPU) sta() uint8 {
	return 0
}

func (c *CPU) sty() uint8 {
	return 0
}

func (c *CPU) stx() uint8 {
	return 0
}

func (c *CPU) dey() uint8 {
	return 0
}

func (c *CPU) txa() uint8 {
	return 0
}

func (c *CPU) bcc() uint8 {
	return 0
}

func (c *CPU) tya() uint8 {
	return 0
}

func (c *CPU) txs() uint8 {
	return 0
}

func (c *CPU) ldy() uint8 {
	return 0
}

func (c *CPU) lda() uint8 {
	return 0
}

func (c *CPU) ldx() uint8 {
	return 0
}

func (c *CPU) tay() uint8 {
	return 0
}

func (c *CPU) tax() uint8 {
	return 0
}

func (c *CPU) bcs() uint8 {
	return 0
}

func (c *CPU) clv() uint8 {
	return 0
}

func (c *CPU) tsx() uint8 {
	return 0
}

func (c *CPU) cpy() uint8 {
	return 0
}

func (c *CPU) cmp() uint8 {
	return 0
}

func (c *CPU) dec() uint8 {
	return 0
}

func (c *CPU) iny() uint8 {
	return 0
}

func (c *CPU) dex() uint8 {
	return 0
}

func (c *CPU) bne() uint8 {
	return 0
}

func (c *CPU) cld() uint8 {
	return 0
}

func (c *CPU) cpx() uint8 {
	return 0
}

func (c *CPU) sbc() uint8 {
	return 0
}

func (c *CPU) inc() uint8 {
	return 0
}

func (c *CPU) inx() uint8 {
	return 0
}

func (c *CPU) nop() uint8 {
	return 0
}

func (c *CPU) beq() uint8 {
	return 0
}

func (c *CPU) sed() uint8 {
	return 0
}
