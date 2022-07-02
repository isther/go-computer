package gate

import (
	"github.com/isther/go-computer/circuit"
)

// @Summary NAND gate
// @Description Implement the function of an NAND gate
// @Tags Gate

type NANDGate struct {
	output circuit.Wire
}

func NewNANDGate() *NANDGate {
	return &NANDGate{
		output: *circuit.NewWire("Z", false),
	}
}

func (nandGate *NANDGate) Update(valueA, valueB bool) *NANDGate {
	nandGate.output.Update(!(valueA && valueB))
	return nandGate
}

func (nandGate *NANDGate) Value() bool {
	return nandGate.output.Value()
}
