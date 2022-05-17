package component

import (
	"github.com/isther/go-computer/circuit"
	"github.com/isther/go-computer/circuit/gate"
)

type Decoder struct {
	notGates [4]gate.NOTGate
	andGates [16]ANDGate4
	outputs  [16]circuit.Wire
	index    int
}

func NewDecoder() *Decoder {
	d := new(Decoder)

	for i, _ := range d.notGates {
		d.notGates[i] = *gate.NewNOTGate()
	}

	for i, _ := range d.andGates {
		d.andGates[i] = *NewANDGate4()
	}

	return d
}

func (d *Decoder) Index() int {
	return d.index
}

func (d *Decoder) GetOutputWire(index int) bool {
	return d.outputs[index].Value()
}

func (a *Decoder) SetInputWire(index int, value bool) {}

func (d *Decoder) Update(inputA, inputB, inputC, inputD bool) {
	// https://www.elprocus.com/designing-4-to-16-decoder-using-3-to-8-decoder/
	d.notGates[0].Update(inputA)
	d.notGates[1].Update(inputB)
	d.notGates[2].Update(inputC)
	d.notGates[3].Update(inputD)

	d.andGates[0].Update(d.notGates[0].Value(), d.notGates[1].Value(), d.notGates[2].Value(), d.notGates[3].Value())
	d.andGates[1].Update(d.notGates[0].Value(), d.notGates[1].Value(), d.notGates[2].Value(), inputD)
	d.andGates[2].Update(d.notGates[0].Value(), d.notGates[1].Value(), inputC, d.notGates[3].Value())
	d.andGates[3].Update(d.notGates[0].Value(), d.notGates[1].Value(), inputC, inputD)

	d.andGates[4].Update(d.notGates[0].Value(), inputB, d.notGates[2].Value(), d.notGates[3].Value())
	d.andGates[5].Update(d.notGates[0].Value(), inputB, d.notGates[2].Value(), inputD)
	d.andGates[6].Update(d.notGates[0].Value(), inputB, inputC, d.notGates[3].Value())
	d.andGates[7].Update(d.notGates[0].Value(), inputB, inputC, inputD)

	d.andGates[8].Update(inputA, d.notGates[1].Value(), d.notGates[2].Value(), d.notGates[3].Value())
	d.andGates[9].Update(inputA, d.notGates[1].Value(), d.notGates[2].Value(), inputD)
	d.andGates[10].Update(inputA, d.notGates[1].Value(), inputC, d.notGates[3].Value())
	d.andGates[11].Update(inputA, d.notGates[1].Value(), inputC, inputD)

	d.andGates[12].Update(inputA, inputB, d.notGates[2].Value(), d.notGates[3].Value())
	d.andGates[13].Update(inputA, inputB, d.notGates[2].Value(), inputD)
	d.andGates[14].Update(inputA, inputB, inputC, d.notGates[3].Value())
	d.andGates[15].Update(inputA, inputB, inputC, inputD)

	d.index = 0
	for i := 0; i < len(d.outputs); i++ {
		d.outputs[i].Update(d.andGates[i].output.Value())
		if d.outputs[i].Value() {
			d.index += i
		}
	}
}
