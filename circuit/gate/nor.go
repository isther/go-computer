package gate

import "github.com/isther/go-computer/circuit"

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
