package component

import (
	"github.com/isther/go-computer/circuit"
	"github.com/isther/go-computer/circuit/gate"
)

type Comparator struct {
	inputs       [BUS_WIDTH * 2]circuit.Wire
	equalIn      circuit.Wire
	aIsLargerIn  circuit.Wire
	compares     [BUS_WIDTH]Compare
	outputs      [BUS_WIDTH]circuit.Wire
	equalOut     circuit.Wire
	aIsLargerOut circuit.Wire
	next         Component
}

func NewComparator() *Comparator {
	c := new(Comparator)

	for i, _ := range c.compares {
		c.compares[i] = *NewCompare()
	}

	return c
}

func (c *Comparator) ConnectOutput(b Component) {
	c.next = b
}

func (c *Comparator) GetOutputWire(index int) bool {
	return c.outputs[index].Value()
}

func (c *Comparator) SetInputWire(index int, value bool) {
	c.inputs[index].Update(value)
}

func (g *Comparator) Equal() bool {
	return g.equalOut.Value()
}

func (g *Comparator) Larger() bool {
	return g.aIsLargerOut.Value()
}

func (c *Comparator) Update() {
	// these start out as 1 and 0 respectively
	c.equalIn.Update(true)
	c.aIsLargerIn.Update(false)

	// top 16 bits are <b>, bottom 16 bits are <a>
	awire := 0
	bwire := BUS_WIDTH

	for i := range c.compares {
		c.compares[i].Update(c.inputs[awire].Value(), c.inputs[bwire].Value(), c.equalIn.Value(), c.aIsLargerIn.Value())
		c.outputs[i].Update(c.compares[i].Value())
		c.equalOut.Update(c.compares[i].Equal())
		c.aIsLargerOut.Update(c.compares[i].Larger())

		c.equalIn.Update(c.compares[i].Equal())
		c.aIsLargerIn.Update(c.compares[i].Larger())
		awire++
		bwire++
	}
}

type Compare struct {
	inputA circuit.Wire
	inputB circuit.Wire

	xor  gate.XORGate
	not  gate.NOTGate
	or   gate.ORGate
	and  gate.ANDGate
	and3 ANDGate3

	equalIn   circuit.Wire
	largerIn  circuit.Wire
	equalOut  circuit.Wire
	largerOut circuit.Wire
	out       circuit.Wire
}

func NewCompare() *Compare {
	c := new(Compare)

	c.inputA = *circuit.NewWire("a", false)
	c.inputB = *circuit.NewWire("b", false)
	c.equalIn = *circuit.NewWire("equal_in", false)
	c.equalOut = *circuit.NewWire("equal_out", false)
	c.largerIn = *circuit.NewWire("larger_in", false)
	c.largerOut = *circuit.NewWire("larger_out", false)
	c.out = *circuit.NewWire("out", false)

	c.xor = *gate.NewXORGate()
	c.not = *gate.NewNOTGate()
	c.or = *gate.NewORGate()
	c.and = *gate.NewANDGate()
	c.and3 = *NewANDGate3()
	return c
}

func (g *Compare) Equal() bool {
	return g.equalOut.Value()
}

func (g *Compare) Larger() bool {
	return g.largerOut.Value()
}

func (g *Compare) Value() bool {
	return g.out.Value()
}

func (g *Compare) Update(inputA, inputB, equalIn, isLargerIn bool) {
	g.inputA.Update(inputA)
	g.inputB.Update(inputB)
	g.equalIn.Update(equalIn)
	g.largerIn.Update(isLargerIn)

	g.xor.Update(g.inputA.Value(), g.inputB.Value())
	g.not.Update(g.xor.Value())
	g.and.Update(g.not.Value(), g.equalIn.Value())
	g.equalOut.Update(g.and.Value())

	g.and3.Update(g.equalIn.Value(), g.inputA.Value(), g.xor.Value())
	g.or.Update(g.and3.Value(), g.largerIn.Value())
	g.largerOut.Update(g.or.Value())

	g.out.Update(g.xor.Value())
}
