package alu

import (
	"github.com/isther/go-computer/circuit/component"
)

func (alu *ALU) updateOpDecoder() {
	alu.opDecoder.Update(alu.Op[3].Value(), alu.Op[2].Value(), alu.Op[1].Value(), alu.Op[0].Value())
}

func (a *ALU) updateNotter() {
	for i := (component.BUS_WIDTH - 1); i >= 0; i-- {
		a.not.SetInputWire(i, a.inputBusA.GetOutputWire(i))
	}
	a.not.Update()
	a.wireToEnabler(&a.not, NOT)
}

func (a *ALU) updateOrer() {
	a.setWireOnComponent(&a.or)
	a.or.Update()
	a.wireToEnabler(&a.or, OR)
}

func (a *ALU) updateAnder() {
	a.setWireOnComponent(&a.and)
	a.and.Update()
	a.wireToEnabler(&a.and, AND)
}

func (a *ALU) updateXorer() {
	a.setWireOnComponent(&a.xor)
	a.xor.Update()
	a.wireToEnabler(&a.xor, XOR)
}

func (a *ALU) updateLeftShifter() {
	for i := (component.BUS_WIDTH - 1); i >= 0; i-- {
		a.leftShifer.SetInputWire(i, a.inputBusA.GetOutputWire(i))
	}
	a.leftShifer.Update(a.carryIn.Value())
	a.wireToEnabler(&a.leftShifer, SHL)
}

func (a *ALU) updateRightShifter() {
	for i := (component.BUS_WIDTH - 1); i >= 0; i-- {
		a.rightShifer.SetInputWire(i, a.inputBusA.GetOutputWire(i))
	}
	a.rightShifer.Update(a.carryIn.Value())
	a.wireToEnabler(&a.rightShifer, SHR)
}

func (a *ALU) updateComparator() {
	// comparator is not wired to an enabler and runs all the time
	a.setWireOnComponent(&a.comparator)
	a.comparator.Update()
	a.isLarger.Update(a.comparator.Larger())
	a.isEqual.Update(a.comparator.Equal())
}

func (alu *ALU) updateAdder() {
	alu.setWireOnComponent(&alu.adder)
	alu.adder.SetCarryIn(alu.carryIn.Value())
	alu.adder.Update()
	alu.wireToEnabler(&alu.adder, ADD)
}
