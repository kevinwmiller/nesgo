package bus

// Bus connects the CPU to difference components that support reading and writing
type Bus struct {
	maxAddressableSpace uint16
	addresses           []uint16
	AddressRanges       []AddressRange
	components          []Component
}

type AddressRange struct {
	min           uint16
	max           uint16
	mirrored      bool
	mirroredRange *AddressRange
}

// A ReadWriter can be read from and/or written to
type ReadWriter interface {
	Write(address uint16, data uint8)
	Read(address uint16) uint8
}

type Component struct {
	addresses AddressRange
	rw        ReadWriter
}

// NewBus returns a new Bus object with the highest address being maxAddressableSpace
func NewBus(maxAddressableSpace uint16) *Bus {
	return &Bus{
		maxAddressableSpace: maxAddressableSpace,
		addresses:           make([]uint16, maxAddressableSpace),
	}
}

func (b *Bus) Write(address uint16, data uint8) {
	for _, component := range b.components {
		component.rw.Write(address, data)
	}
}

func (b *Bus) Read(address uint16) uint8 {
	for _, component := range b.components {
		component.Read(address)
	}
	return 0
}
