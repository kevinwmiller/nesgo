package cpu6502

// Instruction represents a single 6502 instruction.
// Every instruction has a name, an implementation, an addressing mode, and a set number of cycle to run
type Instruction struct {
	Name           string
	Execute        func(uint8, uint16) uint8
	AddressingMode uint8
	Cycles         uint8
	// The number of cycles to add if a page boundary is crossed
	AdditionalCycles uint8
}

func (c *CPU) buildInstructionTable() [256]Instruction {
	return [...]Instruction{
		{"BRK", c.brk, Rel, 7, 0}, {"ORA", c.ora, Izx, 6, 0}, {"FUT", c.fut, Und, 0, 0}, {"FUT", c.fut, Und, 0, 0}, {"FUT", c.fut, Und, 0, 0},
		{"ORA", c.ora, Zp0, 3, 0}, {"ASL", c.asl, Zp0, 5, 0}, {"FUT", c.fut, Und, 0, 0}, {"PHP", c.php, Imp, 3, 0}, {"ORA", c.ora, Imm, 2, 0},
		{"ASL", c.asl, Acc, 2, 0}, {"FUT", c.fut, Und, 0, 0}, {"FUT", c.fut, Und, 0, 0}, {"ORA", c.ora, Abs, 4, 0}, {"ASL", c.asl, Abs, 6, 0},
		{"FUT", c.fut, Und, 0, 0}, {"BPL", c.bpl, Rel, 2, 1}, {"ORA", c.ora, Izy, 5, 1}, {"FUT", c.fut, Und, 0, 0}, {"FUT", c.fut, Und, 0, 0},
		{"FUT", c.fut, Und, 0, 0}, {"ORA", c.ora, Zpx, 4, 0}, {"ASL", c.asl, Zpx, 6, 0}, {"FUT", c.fut, Und, 0, 0}, {"CLC", c.clc, Imp, 2, 0},
		{"ORA", c.ora, Aby, 4, 1}, {"FUT", c.fut, Und, 0, 0}, {"FUT", c.fut, Und, 0, 0}, {"FUT", c.fut, Und, 0, 0}, {"ORA", c.ora, Abx, 4, 1},
		{"ASL", c.asl, Abx, 7, 0}, {"FUT", c.fut, Und, 0, 0}, {"JSR", c.jsr, Abs, 6, 0}, {"AND", c.and, Izx, 6, 0}, {"FUT", c.fut, Und, 0, 0},
		{"FUT", c.fut, Und, 0, 0}, {"BIT", c.bit, Zp0, 3, 0}, {"AND", c.and, Zp0, 3, 0}, {"ROL", c.rol, Zp0, 5, 0}, {"FUT", c.fut, Und, 0, 0},
		{"PLP", c.plp, Imp, 4, 0}, {"AND", c.and, Imm, 2, 0}, {"ROL", c.rol, Acc, 2, 0}, {"FUT", c.fut, Und, 0, 0}, {"BIT", c.bit, Abs, 4, 0},
		{"AND", c.and, Abs, 4, 0}, {"ROL", c.rol, Abs, 6, 0}, {"FUT", c.fut, Und, 0, 0}, {"BMI", c.bmi, Rel, 2, 1}, {"AND", c.and, Izy, 5, 1},
		{"FUT", c.fut, Und, 0, 0}, {"FUT", c.fut, Und, 0, 0}, {"FUT", c.fut, Und, 0, 0}, {"AND", c.and, Zpx, 4, 0}, {"ROL", c.rol, Zpx, 6, 0},
		{"FUT", c.fut, Und, 0, 0}, {"SEC", c.sec, Imp, 2, 0}, {"AND", c.and, Aby, 4, 1}, {"FUT", c.fut, Und, 0, 0}, {"FUT", c.fut, Und, 0, 0},
		{"FUT", c.fut, Und, 0, 0}, {"AND", c.and, Abx, 4, 1}, {"ROL", c.rol, Abx, 7, 0}, {"FUT", c.fut, Und, 0, 0}, {"RTI", c.rti, Imp, 6, 0},
		{"EOR", c.eor, Izx, 6, 0}, {"FUT", c.fut, Und, 0, 0}, {"FUT", c.fut, Und, 0, 0}, {"FUT", c.fut, Und, 0, 0}, {"EOR", c.eor, Zp0, 3, 0},
		{"LSR", c.lsr, Zp0, 5, 0}, {"FUT", c.fut, Und, 0, 0}, {"PHA", c.pha, Imp, 3, 0}, {"EOR", c.eor, Imm, 2, 0}, {"LSR", c.lsr, Acc, 2, 0},
		{"FUT", c.fut, Und, 0, 0}, {"JMP", c.jmp, Abs, 3, 0}, {"EOR", c.eor, Abs, 4, 0}, {"LSR", c.lsr, Abs, 6, 0}, {"FUT", c.fut, Und, 0, 0},
		{"BVC", c.bvc, Rel, 2, 1}, {"EOR", c.eor, Izy, 5, 1}, {"FUT", c.fut, Und, 0, 0}, {"FUT", c.fut, Und, 0, 0}, {"FUT", c.fut, Und, 0, 0},
		{"EOR", c.eor, Zpx, 4, 0}, {"LSR", c.lsr, Zpx, 6, 0}, {"FUT", c.fut, Und, 0, 0}, {"CLI", c.cli, Imp, 2, 0}, {"EOR", c.eor, Aby, 4, 1},
		{"FUT", c.fut, Und, 0, 0}, {"FUT", c.fut, Und, 0, 0}, {"FUT", c.fut, Und, 0, 0}, {"EOR", c.eor, Abx, 4, 1}, {"LSR", c.lsr, Abx, 7, 0},
		{"FUT", c.fut, Und, 0, 0}, {"RTS", c.rts, Imp, 6, 0}, {"ADC", c.adc, Izx, 6, 0}, {"FUT", c.fut, Und, 0, 0}, {"FUT", c.fut, Und, 0, 0},
		{"FUT", c.fut, Und, 0, 0}, {"ADC", c.adc, Zp0, 3, 0}, {"ROR", c.ror, Zp0, 5, 0}, {"FUT", c.fut, Und, 0, 0}, {"PLA", c.pla, Imp, 4, 0},
		{"ADC", c.adc, Imm, 2, 0}, {"ROR", c.ror, Acc, 2, 0}, {"FUT", c.fut, Und, 0, 0}, {"JMP", c.jmp, Ind, 5, 0}, {"ADC", c.adc, Abs, 4, 0},
		{"ROR", c.ror, Abs, 6, 0}, {"FUT", c.fut, Und, 0, 0}, {"BVS", c.bvs, Rel, 2, 1}, {"ADC", c.adc, Izy, 5, 1}, {"FUT", c.fut, Und, 0, 0},
		{"FUT", c.fut, Und, 0, 0}, {"FUT", c.fut, Und, 0, 0}, {"ADC", c.adc, Zpx, 4, 0}, {"ROR", c.ror, Zpx, 6, 0}, {"FUT", c.fut, Und, 0, 0},
		{"SEI", c.sei, Imp, 2, 0}, {"ADC", c.adc, Aby, 4, 1}, {"FUT", c.fut, Und, 0, 0}, {"FUT", c.fut, Und, 0, 0}, {"FUT", c.fut, Und, 0, 0},
		{"ADC", c.adc, Abx, 4, 1}, {"ROR", c.ror, Abx, 7, 0}, {"FUT", c.fut, Und, 0, 0}, {"FUT", c.fut, Und, 0, 0}, {"STA", c.sta, Izx, 6, 0},
		{"FUT", c.fut, Und, 0, 0}, {"FUT", c.fut, Und, 0, 0}, {"STY", c.sty, Zp0, 3, 0}, {"STA", c.sta, Zp0, 3, 0}, {"STX", c.stx, Zp0, 3, 0},
		{"FUT", c.fut, Und, 0, 0}, {"DEY", c.dey, Imp, 2, 0}, {"FUT", c.fut, Und, 0, 0}, {"TXA", c.txa, Imp, 2, 0}, {"FUT", c.fut, Und, 0, 0},
		{"STY", c.sty, Abs, 4, 0}, {"STA", c.sta, Abs, 4, 0}, {"STX", c.stx, Abs, 4, 0}, {"FUT", c.fut, Und, 0, 0}, {"BCC", c.bcc, Rel, 2, 1},
		{"STA", c.sta, Izy, 6, 0}, {"FUT", c.fut, Und, 0, 0}, {"FUT", c.fut, Und, 0, 0}, {"STY", c.sty, Zpx, 4, 0}, {"STA", c.sta, Zpx, 4, 0},
		{"STX", c.stx, Zpy, 4, 0}, {"FUT", c.fut, Und, 0, 0}, {"TYA", c.tya, Imp, 2, 0}, {"STA", c.sta, Aby, 5, 0}, {"TXS", c.txs, Imp, 2, 0},
		{"FUT", c.fut, Und, 0, 0}, {"FUT", c.fut, Und, 0, 0}, {"STA", c.sta, Abx, 5, 0}, {"FUT", c.fut, Und, 0, 0}, {"FUT", c.fut, Und, 0, 0},
		{"LDY", c.ldy, Imm, 2, 0}, {"LDA", c.lda, Izx, 6, 0}, {"LDX", c.ldx, Imm, 2, 0}, {"FUT", c.fut, Und, 0, 0}, {"LDY", c.ldy, Zp0, 3, 0},
		{"LDA", c.lda, Zp0, 3, 0}, {"LDX", c.ldx, Zp0, 3, 0}, {"FUT", c.fut, Und, 0, 0}, {"TAY", c.tay, Imp, 2, 0}, {"LDA", c.lda, Imm, 2, 0},
		{"TAX", c.tax, Imp, 2, 0}, {"FUT", c.fut, Und, 0, 0}, {"LDY", c.ldy, Abs, 4, 0}, {"LDA", c.lda, Abs, 4, 0}, {"LDX", c.ldx, Abs, 4, 0},
		{"FUT", c.fut, Und, 0, 0}, {"BCS", c.bcs, Rel, 2, 1}, {"LDA", c.lda, Izy, 5, 1}, {"FUT", c.fut, Und, 0, 0}, {"FUT", c.fut, Und, 0, 0},
		{"LDY", c.ldy, Zpx, 4, 0}, {"LDA", c.lda, Zpx, 4, 0}, {"LDX", c.ldx, Zpy, 4, 0}, {"FUT", c.fut, Und, 0, 0}, {"CLV", c.clv, Imp, 2, 0},
		{"LDA", c.lda, Aby, 4, 1}, {"TSX", c.tsx, Imp, 2, 0}, {"FUT", c.fut, Und, 0, 0}, {"LDY", c.ldy, Abx, 4, 1}, {"LDA", c.lda, Abx, 4, 1},
		{"LDX", c.ldx, Aby, 4, 1}, {"FUT", c.fut, Und, 0, 0}, {"Cpy", c.cpy, Imm, 2, 0}, {"CMP", c.cmp, Izx, 6, 0}, {"FUT", c.fut, Und, 0, 0},
		{"FUT", c.fut, Und, 0, 0}, {"CPY", c.cpy, Zp0, 3, 0}, {"CMP", c.cmp, Zp0, 3, 0}, {"DEC", c.dec, Zp0, 5, 0}, {"FUT", c.fut, Und, 0, 0},
		{"INY", c.iny, Imp, 2, 0}, {"CMP", c.cmp, Imm, 2, 0}, {"DEX", c.dex, Imp, 2, 0}, {"FUT", c.fut, Und, 0, 0}, {"CPY", c.cpy, Abs, 4, 0},
		{"CMP", c.cmp, Abs, 4, 0}, {"DEC", c.dec, Abs, 6, 0}, {"FUT", c.fut, Und, 0, 0}, {"BNE", c.bne, Rel, 2, 1}, {"CMP", c.cmp, Izy, 5, 1},
		{"FUT", c.fut, Und, 0, 0}, {"FUT", c.fut, Und, 0, 0}, {"FUT", c.fut, Und, 0, 0}, {"CMP", c.cmp, Zpx, 4, 0}, {"DEC", c.dec, Zpx, 6, 0},
		{"FUT", c.fut, Und, 0, 0}, {"CLD", c.cld, Imp, 2, 0}, {"CMP", c.cmp, Aby, 4, 1}, {"FUT", c.fut, Und, 0, 0}, {"FUT", c.fut, Und, 0, 0},
		{"FUT", c.fut, Und, 0, 0}, {"CMP", c.cmp, Abx, 4, 1}, {"DEC", c.dec, Abx, 7, 0}, {"FUT", c.fut, Und, 0, 0}, {"CPX", c.cpx, Imm, 2, 0},
		{"SBC", c.sbc, Izx, 6, 0}, {"FUT", c.fut, Und, 0, 0}, {"FUT", c.fut, Und, 0, 0}, {"CPX", c.cpx, Zp0, 3, 0}, {"SBC", c.sbc, Zp0, 3, 0},
		{"INC", c.inc, Zp0, 5, 0}, {"FUT", c.fut, Und, 0, 0}, {"INX", c.inx, Imp, 2, 0}, {"SBC", c.sbc, Imm, 2, 0}, {"NOP", c.nop, Imp, 2, 0},
		{"FUT", c.fut, Und, 0, 0}, {"CPX", c.cpx, Abs, 4, 0}, {"SBC", c.sbc, Abs, 4, 0}, {"INC", c.inc, Abs, 6, 0}, {"FUT", c.fut, Und, 0, 0},
		{"BEQ", c.beq, Rel, 2, 1}, {"SBC", c.sbc, Izy, 5, 1}, {"FUT", c.fut, Und, 0, 0}, {"FUT", c.fut, Und, 0, 0}, {"FUT", c.fut, Und, 0, 0},
		{"SBC", c.sbc, Zpx, 4, 0}, {"INC", c.inc, Zpx, 6, 0}, {"FUT", c.fut, Und, 0, 0}, {"SED", c.sed, Imp, 2, 0}, {"SBC", c.sbc, Aby, 4, 1},
		{"FUT", c.fut, Und, 0, 0}, {"FUT", c.fut, Und, 0, 0}, {"FUT", c.fut, Und, 0, 0}, {"SBC", c.sbc, Abx, 4, 1}, {"INC", c.inc, Abx, 7, 0},
		{"FUT", c.fut, Und, 0, 0},
	}
}

// Future Expansion
func (c *CPU) fut(data uint8, address uint16) uint8 {
	return 0
}

func (c *CPU) brk(data uint8, address uint16) uint8 {
	return 0
}

func (c *CPU) ora(data uint8, address uint16) uint8 {
	return 0
}

func (c *CPU) asl(data uint8, address uint16) uint8 {
	return 0
}

func (c *CPU) php(data uint8, address uint16) uint8 {
	return 0
}

func (c *CPU) bpl(data uint8, address uint16) uint8 {
	return 0
}

func (c *CPU) clc(data uint8, address uint16) uint8 {
	return 0
}

func (c *CPU) jsr(data uint8, address uint16) uint8 {
	return 0
}

func (c *CPU) and(data uint8, address uint16) uint8 {
	return 0
}

func (c *CPU) bit(data uint8, address uint16) uint8 {
	return 0
}

func (c *CPU) rol(data uint8, address uint16) uint8 {
	return 0
}

func (c *CPU) plp(data uint8, address uint16) uint8 {
	return 0
}

func (c *CPU) bmi(data uint8, address uint16) uint8 {
	return 0
}

func (c *CPU) sec(data uint8, address uint16) uint8 {
	return 0
}

func (c *CPU) rti(data uint8, address uint16) uint8 {
	return 0
}

func (c *CPU) eor(data uint8, address uint16) uint8 {
	return 0
}

func (c *CPU) lsr(data uint8, address uint16) uint8 {
	return 0
}

func (c *CPU) pha(data uint8, address uint16) uint8 {
	return 0
}

func (c *CPU) jmp(data uint8, address uint16) uint8 {
	return 0
}

func (c *CPU) bvc(data uint8, address uint16) uint8 {
	return 0
}

func (c *CPU) cli(data uint8, address uint16) uint8 {
	return 0
}

func (c *CPU) rts(data uint8, address uint16) uint8 {
	return 0
}

func (c *CPU) adc(data uint8, address uint16) uint8 {
	return 0
}

func (c *CPU) ror(data uint8, address uint16) uint8 {
	return 0
}

func (c *CPU) pla(data uint8, address uint16) uint8 {
	return 0
}

func (c *CPU) bvs(data uint8, address uint16) uint8 {
	return 0
}

func (c *CPU) sei(data uint8, address uint16) uint8 {
	return 0
}

func (c *CPU) sta(data uint8, address uint16) uint8 {
	return 0
}

func (c *CPU) sty(data uint8, address uint16) uint8 {
	return 0
}

func (c *CPU) stx(data uint8, address uint16) uint8 {
	return 0
}

func (c *CPU) dey(data uint8, address uint16) uint8 {
	return 0
}

func (c *CPU) txa(data uint8, address uint16) uint8 {
	return 0
}

func (c *CPU) bcc(data uint8, address uint16) uint8 {
	return 0
}

func (c *CPU) tya(data uint8, address uint16) uint8 {
	return 0
}

func (c *CPU) txs(data uint8, address uint16) uint8 {
	return 0
}

func (c *CPU) ldy(data uint8, address uint16) uint8 {
	return 0
}

func (c *CPU) lda(data uint8, address uint16) uint8 {
	return 0
}

func (c *CPU) ldx(data uint8, address uint16) uint8 {
	return 0
}

func (c *CPU) tay(data uint8, address uint16) uint8 {
	return 0
}

func (c *CPU) tax(data uint8, address uint16) uint8 {
	return 0
}

func (c *CPU) bcs(data uint8, address uint16) uint8 {
	return 0
}

func (c *CPU) clv(data uint8, address uint16) uint8 {
	return 0
}

func (c *CPU) tsx(data uint8, address uint16) uint8 {
	return 0
}

func (c *CPU) cpy(data uint8, address uint16) uint8 {
	return 0
}

func (c *CPU) cmp(data uint8, address uint16) uint8 {
	return 0
}

func (c *CPU) dec(data uint8, address uint16) uint8 {
	return 0
}

func (c *CPU) iny(data uint8, address uint16) uint8 {
	return 0
}

func (c *CPU) dex(data uint8, address uint16) uint8 {
	return 0
}

func (c *CPU) bne(data uint8, address uint16) uint8 {
	return 0
}

func (c *CPU) cld(data uint8, address uint16) uint8 {
	return 0
}

func (c *CPU) cpx(data uint8, address uint16) uint8 {
	return 0
}

func (c *CPU) sbc(data uint8, address uint16) uint8 {
	return 0
}

func (c *CPU) inc(data uint8, address uint16) uint8 {
	return 0
}

func (c *CPU) inx(data uint8, address uint16) uint8 {
	return 0
}

func (c *CPU) nop(data uint8, address uint16) uint8 {
	return 0
}

func (c *CPU) beq(data uint8, address uint16) uint8 {
	return 0
}

func (c *CPU) sed(data uint8, address uint16) uint8 {
	return 0
}
