package component

import (
	"github.com/isther/go-computer/circuit"
	"github.com/isther/go-computer/circuit/gate"
)

type Decoder8x256 struct {
	decoderSelector Decoder4x16
	decoders4x16    [16]Decoder4x16
	index           int
}

func NewDecoder8x256() *Decoder8x256 {
	d := new(Decoder8x256)

	d.decoderSelector = *NewDecoder4x16()

	for i := range d.decoders4x16 {
		d.decoders4x16[i] = *NewDecoder4x16()
	}

	return d
}

// Returns the index which is enabled
func (d *Decoder8x256) Index() int {
	return d.index
}

func (dc *Decoder8x256) Update(a, b, c, d, e, f, g, h bool) {
	dc.index = 0

	dc.decoderSelector.Update(e, f, g, h)
	for i := 0; i < 16; i++ {
		dc.updateDecoder(a, b, c, d, i, 16*i)
	}
}

func (dc *Decoder8x256) updateDecoder(a, b, c, d bool, decoderIndex int, outputWireStart int) {
	if dc.decoderSelector.GetOutputWire(decoderIndex) {
		dc.decoders4x16[decoderIndex].Update(a, b, c, d)

		for i := 0; i < 16; i++ {
			if dc.decoders4x16[decoderIndex].outputs[i].Value() {
				dc.index = outputWireStart + i
			}
		}
	}
}

type Decoder4x16 struct {
	notGates [4]gate.NOTGate
	andGates [16]ANDGate4
	outputs  [16]circuit.Wire
	index    int
}

func NewDecoder4x16() *Decoder4x16 {
	d := new(Decoder4x16)

	for i, _ := range d.notGates {
		d.notGates[i] = *gate.NewNOTGate()
	}

	for i, _ := range d.andGates {
		d.andGates[i] = *NewANDGate4()
	}

	return d
}

func (d *Decoder4x16) Index() int {
	return d.index
}

func (d *Decoder4x16) GetOutputWire(index int) bool {
	return d.outputs[index].Value()
}

func (a *Decoder4x16) SetInputWire(index int, value bool) {}

func (d *Decoder4x16) Update(inputA, inputB, inputC, inputD bool) {
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
