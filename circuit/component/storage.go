package component

import (
	"github.com/isther/go-computer/circuit"
	"github.com/isther/go-computer/circuit/gate"
)

type Word struct {
	inputs  [16]circuit.Wire
	bits    [16]Bit
	outputs [16]circuit.Wire
	next    Component
}

func NewWord() *Word {
	word := new(Word)
	for i, _ := range word.bits {
		word.bits[i].Update(false, true)
	}
	return word
}

func (w *Word) ConnectOutput(b Component) {
	w.next = b
}

func (w *Word) GetOutputWire(index int) bool {
	return w.outputs[index].Value()
}

func (w *Word) SetInputWire(index int, value bool) {
	w.inputs[index].Update(value)
}

func (w *Word) Value() uint16 {
	var (
		value uint16
		x     uint16
	)
	for i := BUS_WIDTH - 1; i >= 0; i-- {
		if w.GetOutputWire(i) {
			value = value | (1 << x)
		} else {
			value = value &^ (1 << x)
		}
		x++
	}

	return value
}

func (w *Word) Update(set bool) {
	for i := 0; i < len(w.inputs); i++ {
		w.bits[i].Update(w.inputs[i].Value(), set)
		w.outputs[i].Update(w.bits[i].Value())
	}

	if w.next != nil {
		for i := 0; i < len(w.outputs); i++ {
			w.next.SetInputWire(i, w.outputs[i].Value())
		}
	}
}

type Bit struct {
	nandGates [4]gate.NANDGate
	q         circuit.Wire
}

func NewBit() *Bit {
	gates := [4]gate.NANDGate{
		*gate.NewNANDGate(),
		*gate.NewNANDGate(),
		*gate.NewNANDGate(),
		*gate.NewNANDGate(),
	}

	return &Bit{
		nandGates: gates,
		q:         *circuit.NewWire("Q", false),
	}
}

func (m *Bit) Value() bool {
	return m.q.Value()
}

func (m *Bit) Update(wireI bool, wireS bool) {
	for i := 0; i < 2; i++ {
		m.nandGates[0].Update(wireI, wireS)
		m.nandGates[1].Update(m.nandGates[0].Value(), wireS)
		m.nandGates[2].Update(m.nandGates[0].Value(), m.nandGates[3].Value())
		m.nandGates[3].Update(m.nandGates[2].Value(), m.nandGates[1].Value())
		m.q.Update(m.nandGates[2].Value())
	}
}
