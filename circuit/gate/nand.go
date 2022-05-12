package gate

import (
	"github.com/isther/computer/circuit"
)

type NANDGate struct {
	output circuit.Wire
}

func NewNANDGate() *NANDGate {
	return &NANDGate{
		output: *circuit.NewWire(false),
	}
}

func (nandGate *NANDGate) Update(valueA, valueB bool) *NANDGate {
	nandGate.output.Update(!(valueA && valueB))
	return nandGate
}

func (nandGate *NANDGate) Value() bool {
	return nandGate.output.Value()
}
