package integrated

import (
	"github.com/isther/go-computer/circuit"
	"github.com/isther/go-computer/circuit/gate"
)

type FullAdder struct {
	halfAdder1 HalfAdder
	halfAdder2 HalfAdder
	orGate     gate.ORGate

	si circuit.Wire
	ci circuit.Wire
}

func NewFullAdder() *FullAdder {
	return &FullAdder{
		halfAdder1: *NewHalfAdder(),
		halfAdder2: *NewHalfAdder(),
		orGate:     *gate.NewORGate(),
		si:         *circuit.NewWire("S_i", false),
		ci:         *circuit.NewWire("C_i", false),
	}
}

func (fullAdder *FullAdder) Update(valueA, valueB, valueC bool) *FullAdder {
	fullAdder.halfAdder1.Update(valueA, valueB)
	fullAdder.halfAdder2.Update(fullAdder.halfAdder1.S0(), valueC)
	fullAdder.ci.Update(fullAdder.orGate.Update(fullAdder.halfAdder1.C0(), fullAdder.halfAdder2.C0()).Value())
	fullAdder.si.Update(fullAdder.halfAdder2.S0())
	return fullAdder
}

func (fullAdder *FullAdder) Value() (bool, bool) {
	return fullAdder.si.Value(), fullAdder.ci.Value()
}
func (fullAdder *FullAdder) Si() bool {
	return fullAdder.si.Value()
}
func (fullAdder *FullAdder) Ci() bool {
	return fullAdder.ci.Value()
}
