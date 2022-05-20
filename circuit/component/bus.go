package component

import (
	"github.com/isther/go-computer/circuit"
)

const BUS_WIDTH = 16

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

func (bus *Bus) SetInputWire(index int, value bool) {
	bus.wires[index].Update(value)
}

func (bus *Bus) GetOutputWire(index int) bool {
	return bus.wires[index].Value()
}

func (bus *Bus) SetValue(value uint16) {
	var x = 0
	for i := bus.width - 1; i >= 0; i-- {
		r := (value & (1 << uint16(x)))
		if r != 0 {
			bus.SetInputWire(i, true)
		} else {
			bus.SetInputWire(i, false)
		}
		x++
	}
}

func (bus *Bus) Value() uint16 {
	var (
		result uint16
		x      uint16
	)
	for i := BUS_WIDTH - 1; i >= 0; i-- {
		if bus.GetOutputWire(i) {
			result = result | (1 << uint16(x))
		} else {
			result = result & ^(1 << uint16(x))
		}
		x++
	}
	return result
}

func (bus *Bus) String() string {
	ret := ""
	for i := 0; i < bus.width; i++ {
		if bus.GetOutputWire(i) {
			ret += "1"
		} else {
			ret += "0"
		}
	}
	return ret
}
