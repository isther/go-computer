package component

import (
	"github.com/isther/go-computer/circuit"
	"github.com/isther/go-computer/circuit/gate"
)

type Enabler struct {
	inputs  [BUS_WIDTH]circuit.Wire
	gates   [BUS_WIDTH]gate.ANDGate
	outputs [BUS_WIDTH]circuit.Wire
	next    Component
}

func NewEnabler() *Enabler {
	e := new(Enabler)

	for i := range e.gates {
		e.gates[i] = *gate.NewANDGate()
	}
	return e
}

func (e *Enabler) ConnectOutput(b Component) {
	e.next = b
}

func (e *Enabler) GetOutputWire(index int) bool {
	return e.outputs[index].Value()
}

func (e *Enabler) SetInputWire(index int, value bool) {
	e.inputs[index].Update(value)
}

func (e *Enabler) Update(enable bool) {
	for i := 0; i < len(e.gates); i++ {
		e.gates[i].Update(e.inputs[i].Value(), enable)
		e.outputs[i].Update(e.gates[i].Value())
	}

	if e.next != nil {
		for i := 0; i < len(e.outputs); i++ {
			e.next.SetInputWire(i, e.outputs[i].Value())
		}
	}
}
