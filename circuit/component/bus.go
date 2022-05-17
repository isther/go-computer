package component

import (
	"github.com/isther/go-computer/circuit"
)

const BUS_WIDTH = 32

type Bus struct {
	wires []circuit.Wire
	width int
}

func NewBus(width int) *Bus {
	bus := new(Bus)
	bus.width = width
	bus.wires = make([]circuit.Wire, width)
	for i, _ := range bus.wires {
		bus.wires[i] = *circuit.NewWire("", false)
	}

	return bus
}

func (bus *Bus) Input(index int, value bool) *Bus {
	bus.wires[index].Update(value)
	return bus
}

func (bus *Bus) Output(index int) bool {
	return bus.wires[index].Value()
}

func (bus *Bus) SetValue(value uint16) {
	var x = 0
	for i := bus.width - 1; i >= 0; i-- {
		r := (value & (1 << uint16(x)))
		if r != 0 {
			bus.Input(i, true)
		} else {
			bus.Input(i, false)
		}
		x++
	}
}

func (bus *Bus) String() string {
	ret := ""
	for i := 0; i < bus.width; i++ {
		if bus.Output(i) {
			ret += "1"
		} else {
			ret += "0"
		}
	}
	return ret
}
