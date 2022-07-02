package gate

import "github.com/isther/go-computer/circuit"

// @Summary OR gate
// @Description Implement the function of an OR gate
// @Tags Gate

type ORGate struct {
	output circuit.Wire
}

func NewORGate() *ORGate {
	return &ORGate{
		output: *circuit.NewWire("Z", false),
	}
}

func (orGate *ORGate) Update(valueA, valueB bool) *ORGate {
	orGate.output.Update(!(!valueA && !valueB))
	return orGate
}

func (orGate *ORGate) Value() bool {
	return orGate.output.Value()
}
