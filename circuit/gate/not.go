package gate

import "github.com/isther/computer/circuit"

type NOTGate struct {
	output circuit.Wire
}

func NewNOTGate() *NOTGate {
	return &NOTGate{
		output: *circuit.NewWire(false),
	}
}

func (notGate *NOTGate) Update(valueA bool) *NOTGate {
	notGate.output.Update(!valueA)
	return notGate
}

func (notGate *NOTGate) Value() bool {
	return notGate.output.Value()
}
