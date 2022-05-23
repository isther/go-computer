package alu

import (
	"fmt"

	"github.com/isther/go-computer/circuit"
	"github.com/isther/go-computer/circuit/component"
	"github.com/isther/go-computer/circuit/gate"
)

const (
	NOT = iota
	AND
	OR
	XOR
	SHL
	SHR
	CMP

	ADD
	SUB
	MUL
	DIV
)

type ALU struct {
	inputBusA *component.Bus
	inputBusB *component.Bus
	outputBus *component.Bus

	carryIn  circuit.Wire
	carryOut circuit.Wire
	isLarger circuit.Wire
	isEqual  circuit.Wire

	Op        [4]circuit.Wire
	opDecoder component.Decoder4x16

	xor         component.XORGates
	or          component.ORGates
	and         component.ANDGates
	not         component.NOTGates
	leftShifer  component.LeftShifter
	rightShifer component.RightShifter
	comparator  component.Comparator
	adder       component.Adder16Bit

	enablers [16]component.Enabler
	andGates [3]gate.ANDGate
}

func NewALU(inputBusA, inputBusB, outputBus *component.Bus) *ALU {
	alu := new(ALU)
	alu.inputBusA = inputBusA
	alu.inputBusB = inputBusB
	alu.outputBus = outputBus

	alu.opDecoder = *component.NewDecoder4x16()

	alu.xor = *component.NewXORGates()
	alu.or = *component.NewORGates()
	alu.and = *component.NewANDGates()
	alu.not = *component.NewNOTGates()
	alu.leftShifer = *component.NewLeftShifter()
	alu.rightShifer = *component.NewRightShifter()
	alu.comparator = *component.NewComparator()
	alu.adder = *component.NewAdder16Bit()

	for i := range alu.enablers {
		alu.enablers[i] = *component.NewEnabler()
	}

	alu.andGates[0] = *gate.NewANDGate()
	alu.andGates[1] = *gate.NewANDGate()
	alu.andGates[2] = *gate.NewANDGate()

	return alu
}

func (alu *ALU) Update() {
	alu.updateOpDecoder()
	enabler := alu.opDecoder.Index()
	switch enabler {
	case NOT:
		alu.updateNotter()
	case AND:
		alu.updateAnder()
	case OR:
		alu.updateOrer()
	case XOR:
		alu.updateXorer()
	case SHL:
		alu.updateLeftShifter()
	case SHR:
		alu.updateRightShifter()
	case CMP:
		alu.updateComparator()
	case ADD:
		alu.updateAdder()
	// case SUB:
	// case MUL:
	// case DIV:
	default:
		fmt.Println("ERROR")
	}

	if enabler != CMP {
		switch enabler {
		case ADD:
			alu.andGates[0].Update(alu.adder.Carry(), alu.opDecoder.GetOutputWire(ADD))
			alu.carryOut.Update(alu.andGates[0].Value())
		case SHR:
			alu.andGates[1].Update(alu.rightShifer.ShiftOut(), alu.opDecoder.GetOutputWire(SHR))
			alu.carryOut.Update(alu.andGates[1].Value())
		case SHL:
			alu.andGates[2].Update(alu.leftShifer.ShiftOut(), alu.opDecoder.GetOutputWire(SHL))
			alu.carryOut.Update(alu.andGates[2].Value())
		}

		for i := 0; i < component.BUS_WIDTH; i++ {
			alu.outputBus.SetInputWire(i, alu.enablers[enabler].GetOutputWire(i))
		}
	}
}

func (alu *ALU) setWireOnComponent(c component.Component) {
	for i := component.BUS_WIDTH - 1; i >= 0; i-- {
		c.SetInputWire(i, alu.inputBusA.GetOutputWire(i))
	}

	for i := (component.BUS_WIDTH * 2) - 1; i >= component.BUS_WIDTH; i-- {
		c.SetInputWire(i, alu.inputBusB.GetOutputWire(i-component.BUS_WIDTH))
	}
}

func (a *ALU) wireToEnabler(b component.Component, enablerIndex int) {
	for i := 0; i < component.BUS_WIDTH; i++ {
		a.enablers[enablerIndex].SetInputWire(i, b.GetOutputWire(i))
	}
	a.enablers[enablerIndex].Update(true)
}
