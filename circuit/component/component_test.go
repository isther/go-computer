package component

import "github.com/isther/go-computer/circuit"

type TempComponent struct {
	wires [BUS_WIDTH]circuit.Wire
	next  Component
}

func (d *TempComponent) ConnectOutput(b Component) {
	d.next = b
}

func (d *TempComponent) SetInputWire(index int, value bool) {
	d.wires[index].Update(value)
}

func (d *TempComponent) GetOutputWire(index int) bool {
	return d.wires[index].Value()
}

func (d *TempComponent) Update() {
	for i, w := range d.wires {
		d.next.SetInputWire(i, w.Value())
	}
}
