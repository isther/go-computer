package integrated

import "github.com/isther/go-computer/circuit"

type FourBitAdder struct {
	fullAdder FullAdder

	si [4]circuit.Wire
	ci circuit.Wire
}

func NewFourBitAdder() *FourBitAdder {
	return &FourBitAdder{
		fullAdder: *NewFullAdder(),
		si: [4]circuit.Wire{
			*circuit.NewWire("S0", false),
			*circuit.NewWire("S1", false),
			*circuit.NewWire("S2", false),
			*circuit.NewWire("S3", false),
		},
		ci: *circuit.NewWire("ci", false),
	}
}

func (fourBitAdder *FourBitAdder) Update(valueA [4]bool, valueB [4]bool) *FourBitAdder {
	for i := 0; i < 4; i++ {
		fourBitAdder.fullAdder.Update(valueA[i], valueB[i], fourBitAdder.fullAdder.Ci())
		fourBitAdder.si[i].Update(fourBitAdder.fullAdder.Si())
	}
	fourBitAdder.ci.Update(fourBitAdder.fullAdder.Ci())
	return fourBitAdder
}

func (fourBitAdder *FourBitAdder) Value() (si [4]bool, ci bool) {
	return fourBitAdder.Si(), fourBitAdder.Ci()
}

func (fourBitAdder *FourBitAdder) Si() (si [4]bool) {
	for i := 0; i < 4; i++ {
		si[i] = fourBitAdder.si[i].Value()
	}
	return si
}

func (fourBitAdder *FourBitAdder) Ci() bool {
	return fourBitAdder.ci.Value()
}
