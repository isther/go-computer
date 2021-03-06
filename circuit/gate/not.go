package gate

import "github.com/isther/go-computer/circuit"

// @Summary NOT gate
// @Description Implement the function of an NOT gate
// @Tags Gate

type NOTGate struct {
	output circuit.Wire
}

func NewNOTGate() *NOTGate {
	return &NOTGate{
		output: *circuit.NewWire("Z", false),
	}
}

func (notGate *NOTGate) Update(valueA bool) *NOTGate {
	notGate.output.Update(!valueA)
	return notGate
}

func (notGate *NOTGate) Value() bool {
	return notGate.output.Value()
}
