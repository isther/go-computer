package gate

import "github.com/isther/go-computer/circuit"

// @Summary AND gate
// @Description Implement the function of an AND gate
// @Tags Gate

type ANDGate struct {
	output circuit.Wire
}

func NewANDGate() *ANDGate {
	return &ANDGate{
		output: *circuit.NewWire("Z", false),
	}
}

func (andGate *ANDGate) Update(valueA, valueB bool) *ANDGate {
	andGate.output.Update(valueA && valueB)
	return andGate
}

func (andGate *ANDGate) Value() bool {
	return andGate.output.Value()
}
