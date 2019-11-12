package clock

// Clock is a counter that will call registeredComponents on their specified clock cycles
type Clock struct {
	clockCount           uint64
	registeredComponents []Component
}

// Component represents an object that should be used on a regular interval
type Component struct {
	tickOn uint64
	ticker Ticker
}

// Ticker contains a function Tick that executes a single fetch/decode/execute cycle
type Ticker interface {
	Tick()
}

// RegisterComponent registers a ticker that will tick every $tickOn clock cycles
func (c *Clock) RegisterComponent(ticker Ticker, tickOn uint64) {
	c.registeredComponents = append(c.registeredComponents, Component{
		tickOn: tickOn,
		ticker: ticker,
	})
}

// Start starts the clock
func (c *Clock) Start() {
	for {
		c.Tick()
	}
}

// Tick goes to the next clock cycle
func (c *Clock) Tick() {
	c.clockCount++
	for _, component := range c.registeredComponents {
		if c.clockCount%component.tickOn == 0 {
			component.ticker.Tick()
		}
	}
}
