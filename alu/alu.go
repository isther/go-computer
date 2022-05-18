package alu

import (
	"github.com/isther/go-computer/circuit"
	"github.com/isther/go-computer/circuit/component"
)

const (
	ADD = iota
	SUB
	MUL
	DIV

	NOT
	OR
	AND
	XOR

	CMP
)

type ALU struct {
	inputBusA *component.Bus
	inputBusB *component.Bus
	outputBus *component.Bus

	CarryIn circuit.Wire

	Op        [4]circuit.Wire
	opDecoder component.Decoder

	xor   component.XORGates
	or    component.ORGates
	and   component.ANDGates
	not   component.NOTGates
	adder component.Adder16Bit

	enablers [16]component.Enabler
}

func NewALU(inputBusA, inputBusB, outputBus *component.Bus) *ALU {
	alu := new(ALU)
	alu.inputBusA = inputBusA
	alu.inputBusB = inputBusB
	alu.outputBus = outputBus

	alu.opDecoder = *component.NewDecoder()

	alu.xor = *component.NewXORGates()
	alu.or = *component.NewORGates()
	alu.and = *component.NewANDGates()
	alu.not = *component.NewNOTGates()
	alu.adder = *component.NewAdder16Bit()

	for i := range alu.enablers {
		alu.enablers[i] = *component.NewEnabler()
	}

	return alu
}

func (alu *ALU) Update() {
	alu.updateOpDecoder()
	enabler := alu.opDecoder.Index()

	switch enabler {
	case ADD:
		alu.updateAdder()
	}

	for i := 0; i < component.BUS_WIDTH; i++ {
		alu.outputBus.SetInputWire(i, alu.enablers[enabler].GetOutputWire(i))
	}
}

func (alu *ALU) updateOpDecoder() {
	alu.opDecoder.Update(alu.Op[2].Value(), alu.Op[2].Value(), alu.Op[1].Value(), alu.Op[0].Value())
}

func (alu *ALU) setWireOnComponent(c component.Component) {
	for i := component.BUS_WIDTH - 1; i >= 0; i-- {
		c.SetInputWire(i, alu.inputBusA.GetOutputWire(i))
	}

	for i := (component.BUS_WIDTH * 2) - 1; i >= component.BUS_WIDTH; i-- {
		c.SetInputWire(i, alu.inputBusB.GetOutputWire(i-16))
	}
}

func (alu *ALU) updateAdder() {
	alu.setWireOnComponent(&alu.adder)
	alu.adder.SetCarryIn(alu.CarryIn.Value())
	alu.adder.Update()
	alu.wireToEnabler(&alu.adder, 0)
}

func (a *ALU) wireToEnabler(b component.Component, enablerIndex int) {
	for i := 0; i < component.BUS_WIDTH; i++ {
		a.enablers[enablerIndex].SetInputWire(i, b.GetOutputWire(i))
	}
	a.enablers[enablerIndex].Update(true)
}
