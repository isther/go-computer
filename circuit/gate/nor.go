package gate

import "github.com/isther/go-computer/circuit"

// @Summary NOR gate
// @Description Implement the function of an NOR gate
// @Tags Gate

type NORGate struct {
	output circuit.Wire
}

func NewNORGate() *NORGate {
	return &NORGate{
		output: *circuit.NewWire("Z", false),
	}
}

func (norGate *NORGate) Update(valueA, valueB bool) *NORGate {
	norGate.output.Update(!valueA && !valueB)
	return norGate
}
func (norGate *NORGate) Value() bool {
	return norGate.output.Value()
}
