package integrated

import (
	"github.com/isther/go-computer/circuit"
	"github.com/isther/go-computer/circuit/gate"
)

type HalfAdder struct {
	xorGate gate.XORGate
	andGate gate.ANDGate

	s0 circuit.Wire
	c0 circuit.Wire
}

func NewHalfAdder() *HalfAdder {
	return &HalfAdder{
		xorGate: *gate.NewXORGate(),
		andGate: *gate.NewANDGate(),

		s0: *circuit.NewWire("S0", false),
		c0: *circuit.NewWire("C0", false),
	}
}

func (halfAdder *HalfAdder) Update(valueA, valueB bool) *HalfAdder {
	halfAdder.s0.Update(halfAdder.xorGate.Update(valueA, valueB).Value())
	halfAdder.c0.Update(halfAdder.andGate.Update(valueA, valueB).Value())
	return halfAdder
}

func (halfAdder *HalfAdder) Value() (bool, bool) {
	return halfAdder.s0.Value(), halfAdder.c0.Value()
}
func (halfAdder *HalfAdder) S0() bool {
	return halfAdder.s0.Value()
}
func (halfAdder *HalfAdder) C0() bool {
	return halfAdder.c0.Value()
}
