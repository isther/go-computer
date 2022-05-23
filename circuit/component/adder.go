package component

import (
	"github.com/isther/go-computer/circuit"
	"github.com/isther/go-computer/circuit/gate"
)

type FourBitParallelCarry struct {
}

type Adder16Bit struct {
	Input   [32]circuit.Wire
	carryIn circuit.Wire

	adders [16]FullAdder

	output   [16]circuit.Wire
	carryOut circuit.Wire

	next Component
}

func NewAdder16Bit() *Adder16Bit {
	adder16Bit := new(Adder16Bit)

	for i := range adder16Bit.adders {
		adder16Bit.adders[i] = *NewFullAdder()
	}

	return adder16Bit
}

func (adder16Bit *Adder16Bit) SetCarryIn(carryIn bool) *Adder16Bit {
	adder16Bit.carryIn.Update(carryIn)
	return adder16Bit
}

func (adder16Bit *Adder16Bit) ConnectOutput(c Component) {
	adder16Bit.next = c
}

func (adder16Bit *Adder16Bit) SetInputWire(index int, value bool) {
	adder16Bit.Input[index].Update(value)
	return
}

func (adder16Bit *Adder16Bit) GetOutputWire(index int) bool {
	return adder16Bit.output[index].Value()
}

func (adder16Bit *Adder16Bit) Update() *Adder16Bit {
	aIndex := 32 - 1
	bIndex := 16 - 1
	for i := len(adder16Bit.adders) - 1; i >= 0; i-- {
		adder16Bit.adders[i].Update(adder16Bit.Input[aIndex].Value(), adder16Bit.Input[bIndex].Value(), adder16Bit.carryIn.Value())

		adder16Bit.output[i].Update(adder16Bit.adders[i].Sum())
		adder16Bit.carryOut.Update(adder16Bit.adders[i].Carry())

		adder16Bit.carryIn.Update(adder16Bit.adders[i].Carry())

		aIndex--
		bIndex--
	}
	return adder16Bit
}

func (adder16Bit *Adder16Bit) Carry() bool {
	return adder16Bit.carryOut.Value()
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
