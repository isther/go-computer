package component

import (
	"github.com/isther/go-computer/circuit"
	"github.com/isther/go-computer/circuit/gate"
)

type FourBitParallelCarry struct {
}

type Adder32Bit struct {
	input   [64]circuit.Wire
	carryIn circuit.Wire

	adders [32]FullAdder

	carryOut circuit.Wire
	output   [32]circuit.Wire
}

func NewAdder32Bit() *Adder32Bit {
	adder32Bit := new(Adder32Bit)

	for i, _ := range adder32Bit.adders {
		adder32Bit.adders[i] = *NewFullAdder()
	}

	return adder32Bit
}

func (adder32Bit *Adder32Bit) SetCarryIn(carryIn bool) *Adder32Bit {
	adder32Bit.carryIn.Update(carryIn)
	return adder32Bit
}

func (adder32Bit *Adder32Bit) SetInputWire(index int, value bool) {
	adder32Bit.input[index].Update(value)
	return
}

func (adder32Bit *Adder32Bit) GetOutputWire(index int) bool {
	return adder32Bit.output[index].Value()
}

func (adder32Bit *Adder32Bit) Update() *Adder32Bit {
	aIndex := 63
	bIndex := 31
	for i := len(adder32Bit.adders) - 1; i >= 0; i-- {
		adder32Bit.adders[i].Update(adder32Bit.input[aIndex].Value(), adder32Bit.input[bIndex].Value(), adder32Bit.carryIn.Value())

		adder32Bit.output[i].Update(adder32Bit.adders[i].Sum())
		adder32Bit.carryOut.Update(adder32Bit.adders[i].Carry())

		adder32Bit.carryIn.Update(adder32Bit.adders[i].Carry())

		aIndex--
		bIndex--
	}
	return adder32Bit
}

func (adder32Bit *Adder32Bit) Carry() bool {
	return adder32Bit.carryOut.Value()
}

type FullAdder struct {
	inputA  circuit.Wire
	inputB  circuit.Wire
	carryIn circuit.Wire

	xor1   gate.XORGate
	xor2   gate.XORGate
	and1   gate.ANDGate
	and2   gate.ANDGate
	orGate gate.ORGate

	carryOut circuit.Wire
	sum      circuit.Wire
}

func NewFullAdder() *FullAdder {
	return &FullAdder{
		inputA:  *circuit.NewWire("A", false),
		inputB:  *circuit.NewWire("B", false),
		carryIn: *circuit.NewWire("C", false),

		xor1:   *gate.NewXORGate(),
		xor2:   *gate.NewXORGate(),
		and1:   *gate.NewANDGate(),
		and2:   *gate.NewANDGate(),
		orGate: *gate.NewORGate(),

		carryOut: *circuit.NewWire("CO", false),
		sum:      *circuit.NewWire("SO", false),
	}
}

func (fullAdder *FullAdder) Update(valueA, valueB, carryIn bool) *FullAdder {
	fullAdder.inputA.Update(valueA)
	fullAdder.inputB.Update(valueB)
	fullAdder.carryIn.Update(carryIn)

	fullAdder.sum.Update(
		fullAdder.xor2.Update(
			fullAdder.xor1.Update(
				fullAdder.inputA.Value(), fullAdder.inputB.Value()).Value(),
			fullAdder.carryIn.Value()).Value())

	fullAdder.carryOut.Update(
		fullAdder.orGate.Update(
			fullAdder.and1.Update(fullAdder.inputA.Value(), fullAdder.inputB.Value()).Value(),
			fullAdder.and2.Update(fullAdder.xor1.Value(), fullAdder.carryIn.Value()).Value()).Value())
	return fullAdder
}

func (fullAdder *FullAdder) Sum() bool {
	return fullAdder.sum.Value()
}
func (fullAdder *FullAdder) Carry() bool {
	return fullAdder.carryOut.Value()
}
