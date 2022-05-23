package component

import (
	"github.com/isther/go-computer/circuit"
	"github.com/isther/go-computer/circuit/gate"
)

type XORGates struct {
	inputs  [BUS_WIDTH * 2]circuit.Wire
	gates   [BUS_WIDTH]gate.XORGate
	outputs [BUS_WIDTH]circuit.Wire
	next    Component
}

func NewXORGates() *XORGates {
	a := new(XORGates)

	for i := range a.gates {
		a.gates[i] = *gate.NewXORGate()
	}

	return a
}

func (a *XORGates) ConnectOutput(c Component) {
	a.next = c
}

func (a *XORGates) GetOutputWire(index int) bool {
	return a.outputs[index].Value()
}

func (a *XORGates) SetInputWire(index int, value bool) {
	a.inputs[index].Update(value)
}

func (a *XORGates) Update() {
	awire := BUS_WIDTH
	bwire := 0
	for i := range a.gates {
		a.gates[i].Update(a.inputs[awire].Value(), a.inputs[bwire].Value())
		a.outputs[i].Update(a.gates[i].Value())
		awire++
		bwire++
	}
}

type NOTGates struct {
	inputs  [BUS_WIDTH]circuit.Wire
	gates   [BUS_WIDTH]gate.NOTGate
	outputs [BUS_WIDTH]circuit.Wire
	next    Component
}

func NewNOTGates() *NOTGates {
	a := new(NOTGates)

	for i := range a.gates {
		a.gates[i] = *gate.NewNOTGate()
	}

	return a
}

func (a *NOTGates) ConnectOutput(c Component) {
	a.next = c
}

func (a *NOTGates) GetOutputWire(index int) bool {
	return a.outputs[index].Value()
}

func (a *NOTGates) SetInputWire(index int, value bool) {
	a.inputs[index].Update(value)
}

func (a *NOTGates) Update() {
	wire := 0
	for i := range a.gates {
		a.gates[i].Update(a.inputs[wire].Value())
		a.outputs[i].Update(a.gates[i].Value())
		wire++
	}
}

type ORGates struct {
	Inputs  [BUS_WIDTH * 2]circuit.Wire
	gates   [BUS_WIDTH]gate.ORGate
	outputs [BUS_WIDTH]circuit.Wire

	next Component
}

func NewORGates() *ORGates {
	a := new(ORGates)

	for i := range a.gates {
		a.gates[i] = *gate.NewORGate()
	}

	return a
}

func (a *ORGates) ConnectOutput(c Component) {
	a.next = c
}

func (a *ORGates) GetOutputWire(index int) bool {
	return a.outputs[index].Value()
}

func (a *ORGates) SetInputWire(index int, value bool) {
	a.Inputs[index].Update(value)
}

func (a *ORGates) Update() {
	awire := BUS_WIDTH
	bwire := 0
	for i := range a.gates {
		a.gates[i].Update(a.Inputs[awire].Value(), a.Inputs[bwire].Value())
		a.outputs[i].Update(a.gates[i].Value())
		awire++
		bwire++
	}
}

type ANDGates struct {
	inputs  [BUS_WIDTH * 2]circuit.Wire
	gates   [BUS_WIDTH]gate.ANDGate
	outputs [BUS_WIDTH]circuit.Wire

	next Component
}

func NewANDGates() *ANDGates {
	a := new(ANDGates)

	for i := range a.gates {
		a.gates[i] = *gate.NewANDGate()
	}

	return a
}

func (a *ANDGates) ConnectOutput(c Component) {
	a.next = c
}

func (a *ANDGates) GetOutputWire(index int) bool {
	return a.outputs[index].Value()
}

func (a *ANDGates) SetInputWire(index int, value bool) {
	a.inputs[index].Update(value)
}

func (a *ANDGates) Update() {
	awire := BUS_WIDTH
	bwire := 0
	for i := range a.gates {
		a.gates[i].Update(a.inputs[awire].Value(), a.inputs[bwire].Value())
		a.outputs[i].Update(a.gates[i].Value())
		awire++
		bwire++
	}
}

type ANDGate3 struct {
	inputA circuit.Wire
	inputB circuit.Wire
	inputC circuit.Wire
	andA   gate.ANDGate
	andB   gate.ANDGate
	output circuit.Wire
}

func NewANDGate3() *ANDGate3 {
	a := new(ANDGate3)

	a.inputA = *circuit.NewWire("a", false)
	a.inputB = *circuit.NewWire("b", false)
	a.inputC = *circuit.NewWire("c", false)
	a.output = *circuit.NewWire("d", false)

	a.andA = *gate.NewANDGate()
	a.andB = *gate.NewANDGate()

	return a
}

func (g *ANDGate3) Value() bool {
	return g.output.Value()
}

func (g *ANDGate3) Update(inputA bool, inputB bool, inputC bool) {
	g.andA.Update(inputA, inputB)
	g.andB.Update(g.andA.Value(), inputC)

	g.output.Update(g.andB.Value())
}

type ANDGate4 struct {
	inputA circuit.Wire
	inputB circuit.Wire
	inputC circuit.Wire
	inputD circuit.Wire
	andA   gate.ANDGate
	andB   gate.ANDGate
	andC   gate.ANDGate
	output circuit.Wire
}

func NewANDGate4() *ANDGate4 {
	a := new(ANDGate4)

	a.inputA = *circuit.NewWire("a", false)
	a.inputB = *circuit.NewWire("b", false)
	a.inputC = *circuit.NewWire("c", false)
	a.inputD = *circuit.NewWire("d", false)
	a.output = *circuit.NewWire("o", false)

	a.andA = *gate.NewANDGate()
	a.andB = *gate.NewANDGate()
	a.andC = *gate.NewANDGate()

	return a
}

func (g *ANDGate4) Value() bool {
	return g.output.Value()
}

func (g *ANDGate4) Update(inputA, inputB, inputC, inputD bool) {
	g.andA.Update(inputA, inputB)
	g.andB.Update(g.andA.Value(), inputC)
	g.andC.Update(g.andB.Value(), inputD)
	g.output.Update(g.andC.Value())
}
