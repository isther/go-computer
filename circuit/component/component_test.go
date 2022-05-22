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

func setWireOn16x2(c Component, inputA uint16, inputB uint16) {
	var x uint16 = 0
	for i := 16 - 1; i >= 0; i-- {
		r := (inputA & (1 << x))
		if r != 0 {
			c.SetInputWire(i, true)
		} else {
			c.SetInputWire(i, false)
		}
		x++
	}

	x = 0
	for i := 32 - 1; i >= 16; i-- {
		r := (inputB & (1 << x))
		if r != 0 {
			c.SetInputWire(i, true)
		} else {
			c.SetInputWire(i, false)
		}
		x++
	}
}

func setWireOn16(c Component, value uint16) {
	var x = 0
	for i := BUS_WIDTH - 1; i >= 0; i-- {
		r := (value & (1 << uint16(x)))
		if r != 0 {
			c.SetInputWire(i, true)
		} else {
			c.SetInputWire(i, false)
		}
		x++
	}
}

func getComponentOutput(c Component) uint16 {
	var x int = 0
	var result uint16
	for i := (16 - 1); i >= 0; i-- {
		if c.GetOutputWire(i) {
			result = result | (1 << uint16(x))
		} else {
			result = result & ^(1 << uint16(x))
		}
		x++
	}
	return result
}
