package gate

import "github.com/isther/computer/circuit"

type ORGate struct {
	output circuit.Wire
}

func NewORGate() *ORGate {
	return &ORGate{
		output: *circuit.NewWire(false),
	}
}

func (orGate *ORGate) Update(valueA, valueB bool) *ORGate {
	orGate.output.Update(!(!valueA && !valueB))
	return orGate
}

func (orGate *ORGate) Value() bool {
	return orGate.output.Value()
}
