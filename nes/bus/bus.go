package bus

// Bus connects the CPU to difference components that support reading and writing
type Bus struct {
	RAM [0xFFFF]uint8
}

func (b *Bus) Write(address uint16, data uint8) {
	if address >= 0x0000 && address < 0xFFFF {
		b.RAM[address] = data
	}
}

func (b *Bus) Read(address uint16) uint8 {
	if address >= 0x0000 && address < 0xFFFF {
		return b.RAM[address]
	}
	return 0
}
