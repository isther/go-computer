package gate

import "github.com/isther/go-computer/circuit"

type XORGate struct {
	output circuit.Wire
}

func NewXORGate() *XORGate {
	return &XORGate{
		output: *circuit.NewWire("Z", false),
	}
}

func (xorGate *XORGate) Update(valueA, valueB bool) *XORGate {
	xorGate.output.Update(!(!valueA && !valueB) && !(valueA && valueB))
	return xorGate
}

func (xorGate *XORGate) Value() bool {
	return xorGate.output.Value()
}
