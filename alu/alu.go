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

	Op        [4]circuit.Wire
	opDecoder component.Decoder

	xor component.XORGates
	or  component.ORGates
	and component.ANDGates
	not component.NOTGates

	adder component.Adder32Bit
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
	alu.adder = *component.NewAdder32Bit()

	return alu
}

func (alu *ALU) updateOpDecoder() {
	alu.opDecoder.Update(alu.Op[2].Value(), alu.Op[2].Value(), alu.Op[1].Value(), alu.Op[0].Value())
}
