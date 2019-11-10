package flags

// Bit 0 - C - Carry flag: this holds the carry out of the most significant
// bit in any arithmetic operation. In subtraction operations however, this
// flag is cleared - set to 0 - if a borrow is required, set to 1 - if no
// borrow is required. The carry flag is also used in shift and rotate
// logical operations.
//
// Bit 1 - Z - Zero flag: this is set to 1 when any arithmetic or logical
// operation produces a zero result, and is set to 0 if the result is
// non-zero.
//
// Bit 2 - I: this is an interrupt enable/disable flag. If it is set,
// interrupts are disabled. If it is cleared, interrupts are enabled.
//
// Bit 3 - D: this is the decimal mode status flag. When set, and an Add with
// Carry or Subtract with Carry instruction is executed, the source values are
// treated as valid BCD (Binary Coded Decimal, eg. 0x00-0x99 = 0-99) numbers.
// The result generated is also a BCD number.
//
// Bit 4 - B - this is set when a software interrupt (BRK instruction) is
// executed.
//
// Bit 5 - U - Unused. Supposed to be logical 1 at all times.
//
// Bit 6 - V - Overflow flag: when an arithmetic operation produces a result
// too large to be represented in a byte, V is set.
//
// Bit 7 - S - Sign flag: this is set if the result of an operation is
// negative, cleared if positive.
const (
	C uint8 = 1 << iota
	Z
	I
	D
	B
	U
	V
	S
)

// SetFlag returns a copy of flags with the given flag bit set
func SetFlag(flags, flag uint8) uint8 {
	return flags | flag
}

// ClearFlag returns a copy of flags with the given flag bit cleared
func ClearFlag(flags, flag uint8) uint8 {
	return flags &^ flag
}

// IsFlagSet returns true if the given flag is set in flags
func IsFlagSet(flags, flag uint8) bool {
	return flags&flag != 0
}
