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
		{"BRK", c.brk, Imp, 7, 0}, {"ORA", c.ora, Izx, 0, 0}, {"FUT", c.fut, Und, 0, 0}, {"FUT", c.fut, Und, 0, 0}, {"FUT", c.fut, Und, 0, 0},
		{"ORA", c.ora, Zp0, 0, 0}, {"ASL", c.asl, Zp0, 0, 0}, {"FUT", c.fut, Und, 0, 0}, {"PHP", c.php, Imp, 0, 0}, {"ORA", c.ora, Imm, 0, 0},
		{"ASL", c.asl, Acc, 0, 0}, {"FUT", c.fut, Und, 0, 0}, {"FUT", c.fut, Und, 0, 0}, {"ORA", c.ora, Abs, 0, 0}, {"ASL", c.asl, Abs, 0, 0},
		{"FUT", c.fut, Und, 0, 0}, {"BPL", c.bpl, Imp, 0, 0}, {"ORA", c.ora, Izy, 0, 0}, {"FUT", c.fut, Und, 0, 0}, {"FUT", c.fut, Und, 0, 0},
		{"FUT", c.fut, Und, 0, 0}, {"ORA", c.ora, Zpx, 0, 0}, {"ASL", c.asl, Zpx, 0, 0}, {"FUT", c.fut, Und, 0, 0}, {"CLC", c.clc, Imp, 0, 0},
		{"ORA", c.ora, Aby, 0, 0}, {"FUT", c.fut, Und, 0, 0}, {"FUT", c.fut, Und, 0, 0}, {"FUT", c.fut, Und, 0, 0}, {"ORA", c.ora, Abx, 0, 0},
		{"ASL", c.asl, Abx, 0, 0}, {"FUT", c.fut, Und, 0, 0}, {"JSR", c.jsr, Imp, 0, 0}, {"AND", c.and, Izx, 0, 0}, {"FUT", c.fut, Und, 0, 0},
		{"FUT", c.fut, Und, 0, 0}, {"BIT", c.bit, Zp0, 0, 0}, {"AND", c.and, Zp0, 0, 0}, {"ROL", c.rol, Zp0, 0, 0}, {"FUT", c.fut, Und, 0, 0},
		{"PLP", c.plp, Imp, 0, 0}, {"AND", c.and, Imm, 0, 0}, {"ROL", c.rol, Acc, 0, 0}, {"FUT", c.fut, Und, 0, 0}, {"BIT", c.bit, Abs, 0, 0},
		{"AND", c.and, Abs, 0, 0}, {"ROL", c.rol, Abs, 0, 0}, {"FUT", c.fut, Und, 0, 0}, {"BMI", c.bmi, Imp, 0, 0}, {"AND", c.and, Izy, 0, 0},
		{"FUT", c.fut, Und, 0, 0}, {"FUT", c.fut, Und, 0, 0}, {"FUT", c.fut, Und, 0, 0}, {"AND", c.and, Zpx, 0, 0}, {"ROL", c.rol, Zpx, 0, 0},
		{"FUT", c.fut, Und, 0, 0}, {"SEC", c.sec, Imp, 0, 0}, {"AND", c.and, Aby, 0, 0}, {"FUT", c.fut, Und, 0, 0}, {"FUT", c.fut, Und, 0, 0},
		{"FUT", c.fut, Und, 0, 0}, {"AND", c.and, Abx, 0, 0}, {"ROL", c.rol, Abx, 0, 0}, {"FUT", c.fut, Und, 0, 0}, {"RTI", c.rti, Imp, 0, 0},
		{"EOR", c.eor, Izx, 0, 0}, {"FUT", c.fut, Und, 0, 0}, {"FUT", c.fut, Und, 0, 0}, {"FUT", c.fut, Und, 0, 0}, {"EOR", c.eor, Zp0, 0, 0},
		{"LSR", c.lsr, Zp0, 0, 0}, {"FUT", c.fut, Und, 0, 0}, {"PHA", c.pha, Imp, 0, 0}, {"EOR", c.eor, Imm, 0, 0}, {"LSR", c.lsr, Acc, 0, 0},
		{"FUT", c.fut, Und, 0, 0}, {"JMP", c.jmp, Abs, 0, 0}, {"EOR", c.eor, Abs, 0, 0}, {"LSR", c.lsr, Abs, 0, 0}, {"FUT", c.fut, Und, 0, 0},
		{"BVC", c.bvc, Imp, 0, 0}, {"EOR", c.eor, Izy, 0, 0}, {"FUT", c.fut, Und, 0, 0}, {"FUT", c.fut, Und, 0, 0}, {"FUT", c.fut, Und, 0, 0},
		{"EOR", c.eor, Zpx, 0, 0}, {"LSR", c.lsr, Zpx, 0, 0}, {"FUT", c.fut, Und, 0, 0}, {"CLI", c.cli, Imp, 0, 0}, {"EOR", c.eor, Aby, 0, 0},
		{"FUT", c.fut, Und, 0, 0}, {"FUT", c.fut, Und, 0, 0}, {"FUT", c.fut, Und, 0, 0}, {"EOR", c.eor, Abx, 0, 0}, {"LSR", c.lsr, Abx, 0, 0},
		{"FUT", c.fut, Und, 0, 0}, {"RTS", c.rts, Imp, 0, 0}, {"ADC", c.adc, Izx, 6, 0}, {"FUT", c.fut, Und, 0, 0}, {"FUT", c.fut, Und, 0, 0},
		{"FUT", c.fut, Und, 0, 0}, {"ADC", c.adc, Zp0, 3, 0}, {"ROR", c.ror, Zp0, 0, 0}, {"FUT", c.fut, Und, 0, 0}, {"PLA", c.pla, Imp, 0, 0},
		{"ADC", c.adc, Imm, 2, 0}, {"ROR", c.ror, Acc, 0, 0}, {"FUT", c.fut, Und, 0, 0}, {"JMP", c.jmp, Ind, 0, 0}, {"ADC", c.adc, Abs, 4, 0},
		{"ROR", c.ror, Abs, 0, 0}, {"FUT", c.fut, Und, 0, 0}, {"BVS", c.bvs, Imp, 0, 0}, {"ADC", c.adc, Izy, 0, 0}, {"FUT", c.fut, Und, 0, 0},
		{"FUT", c.fut, Und, 0, 0}, {"FUT", c.fut, Und, 0, 0}, {"ADC", c.adc, Zpx, 4, 0}, {"ROR", c.ror, Zpx, 0, 0}, {"FUT", c.fut, Und, 0, 0},
		{"SEI", c.sei, Imp, 0, 0}, {"ADC", c.adc, Aby, 0, 0}, {"FUT", c.fut, Und, 0, 0}, {"FUT", c.fut, Und, 0, 0}, {"FUT", c.fut, Und, 0, 0},
		{"ADC", c.adc, Abx, 4, 0}, {"ROR", c.ror, Abx, 0, 0}, {"FUT", c.fut, Und, 0, 0}, {"FUT", c.fut, Und, 0, 0}, {"STA", c.sta, Izx, 0, 0},
		{"FUT", c.fut, Und, 0, 0}, {"FUT", c.fut, Und, 0, 0}, {"STY", c.sty, Zp0, 0, 0}, {"STA", c.sta, Zp0, 0, 0}, {"STX", c.stx, Zp0, 0, 0},
		{"FUT", c.fut, Und, 0, 0}, {"DEY", c.dey, Imp, 0, 0}, {"FUT", c.fut, Und, 0, 0}, {"TXA", c.txa, Imp, 0, 0}, {"FUT", c.fut, Und, 0, 0},
		{"STY", c.sty, Abs, 0, 0}, {"STA", c.sta, Abs, 0, 0}, {"STX", c.stx, Abs, 0, 0}, {"FUT", c.fut, Und, 0, 0}, {"BCC", c.bcc, Imp, 0, 0},
		{"STA", c.sta, Izy, 0, 0}, {"FUT", c.fut, Und, 0, 0}, {"FUT", c.fut, Und, 0, 0}, {"STY", c.sty, Zpx, 0, 0}, {"STA", c.sta, Zpx, 0, 0},
		{"STX", c.stx, Zpy, 0, 0}, {"FUT", c.fut, Und, 0, 0}, {"TYA", c.tya, Imp, 0, 0}, {"STA", c.sta, Aby, 0, 0}, {"TXS", c.txs, Imp, 0, 0},
		{"FUT", c.fut, Und, 0, 0}, {"FUT", c.fut, Und, 0, 0}, {"STA", c.sta, Abx, 0, 0}, {"FUT", c.fut, Und, 0, 0}, {"FUT", c.fut, Und, 0, 0},
		{"LDY", c.ldy, Imm, 0, 0}, {"LDA", c.lda, Izx, 0, 0}, {"LDX", c.ldx, Imm, 0, 0}, {"FUT", c.fut, Und, 0, 0}, {"LDY", c.ldy, Zp0, 0, 0},
		{"LDA", c.lda, Zp0, 0, 0}, {"LDX", c.ldx, Zp0, 0, 0}, {"FUT", c.fut, Und, 0, 0}, {"TAY", c.tay, Imp, 0, 0}, {"LDA", c.lda, Imm, 0, 0},
		{"TAX", c.tax, Imp, 0, 0}, {"FUT", c.fut, Und, 0, 0}, {"LDY", c.ldy, Abs, 0, 0}, {"LDA", c.lda, Abs, 0, 0}, {"LDX", c.ldx, Abs, 0, 0},
		{"FUT", c.fut, Und, 0, 0}, {"BCS", c.bcs, Imp, 0, 0}, {"LDA", c.lda, Izy, 0, 0}, {"FUT", c.fut, Und, 0, 0}, {"FUT", c.fut, Und, 0, 0},
		{"LDY", c.ldy, Zpx, 0, 0}, {"LDA", c.lda, Zpx, 0, 0}, {"LDX", c.ldx, Zpy, 0, 0}, {"FUT", c.fut, Und, 0, 0}, {"CLV", c.clv, Imp, 0, 0},
		{"LDA", c.lda, Aby, 0, 0}, {"TSX", c.tsx, Imp, 0, 0}, {"FUT", c.fut, Und, 0, 0}, {"LDY", c.ldy, Abx, 0, 0}, {"LDA", c.lda, Abx, 0, 0},
		{"LDX", c.ldx, Aby, 0, 0}, {"FUT", c.fut, Und, 0, 0}, {"Cpy", c.cpy, Imm, 0, 0}, {"CMP", c.cmp, Izx, 0, 0}, {"FUT", c.fut, Und, 0, 0},
		{"FUT", c.fut, Und, 0, 0}, {"CPY", c.cpy, Zp0, 0, 0}, {"CMP", c.cmp, Zp0, 0, 0}, {"DEC", c.dec, Zp0, 0, 0}, {"FUT", c.fut, Und, 0, 0},
		{"INY", c.iny, Imp, 0, 0}, {"CMP", c.cmp, Imm, 0, 0}, {"DEX", c.dex, Imp, 0, 0}, {"FUT", c.fut, Und, 0, 0}, {"CPY", c.cpy, Abs, 0, 0},
		{"CMP", c.cmp, Abs, 0, 0}, {"DEC", c.dec, Abs, 0, 0}, {"FUT", c.fut, Und, 0, 0}, {"BNE", c.bne, Imp, 0, 0}, {"CMP", c.cmp, Izy, 0, 0},
		{"FUT", c.fut, Und, 0, 0}, {"FUT", c.fut, Und, 0, 0}, {"FUT", c.fut, Und, 0, 0}, {"CMP", c.cmp, Zpx, 0, 0}, {"DEC", c.dec, Zpx, 0, 0},
		{"FUT", c.fut, Und, 0, 0}, {"CLD", c.cld, Imp, 0, 0}, {"CMP", c.cmp, Aby, 0, 0}, {"FUT", c.fut, Und, 0, 0}, {"FUT", c.fut, Und, 0, 0},
		{"FUT", c.fut, Und, 0, 0}, {"CMP", c.cmp, Abx, 0, 0}, {"DEC", c.dec, Abx, 0, 0}, {"FUT", c.fut, Und, 0, 0}, {"CPX", c.cpx, Imm, 0, 0},
		{"SBC", c.sbc, Izx, 0, 0}, {"FUT", c.fut, Und, 0, 0}, {"FUT", c.fut, Und, 0, 0}, {"CPX", c.cpx, Zp0, 0, 0}, {"SBC", c.sbc, Zp0, 0, 0},
		{"INC", c.inc, Zp0, 0, 0}, {"FUT", c.fut, Und, 0, 0}, {"INX", c.inx, Imp, 0, 0}, {"SBC", c.sbc, Imm, 0, 0}, {"NOP", c.nop, Imp, 0, 0},
		{"FUT", c.fut, Und, 0, 0}, {"CPX", c.cpx, Abs, 0, 0}, {"SBC", c.sbc, Abs, 0, 0}, {"INC", c.inc, Abs, 0, 0}, {"FUT", c.fut, Und, 0, 0},
		{"BEQ", c.beq, Imp, 0, 0}, {"SBC", c.sbc, Izy, 0, 0}, {"FUT", c.fut, Und, 0, 0}, {"FUT", c.fut, Und, 0, 0}, {"FUT", c.fut, Und, 0, 0},
		{"SBC", c.sbc, Zpx, 0, 0}, {"INC", c.inc, Zpx, 0, 0}, {"FUT", c.fut, Und, 0, 0}, {"SED", c.sed, Imp, 0, 0}, {"SBC", c.sbc, Aby, 0, 0},
		{"FUT", c.fut, Und, 0, 0}, {"FUT", c.fut, Und, 0, 0}, {"FUT", c.fut, Und, 0, 0}, {"SBC", c.sbc, Abx, 0, 0}, {"INC", c.inc, Abx, 0, 0},
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
