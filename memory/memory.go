package memory

import (
	"fmt"
	"strings"

	"github.com/isther/go-computer/circuit"
	"github.com/isther/go-computer/circuit/component"
	"github.com/isther/go-computer/circuit/gate"
)

type Memory64K struct {
	AddressRegister component.Register
	rowDecoder      component.Decoder8x256
	colDecoder      component.Decoder8x256
	data            [256][256]Cell
	set             circuit.Wire
	enable          circuit.Wire
	bus             *component.Bus
}

func NewMemory64K(bus *component.Bus) *Memory64K {
	m := new(Memory64K)
	m.AddressRegister = *component.NewRegister("MAR", bus, bus)
	m.rowDecoder = *component.NewDecoder8x256()
	m.colDecoder = *component.NewDecoder8x256()
	m.bus = bus

	for i := 0; i < 256; i++ {
		for j := 0; j < 256; j++ {
			m.data[i][j] = *NewCell(bus, bus)
		}
	}

	return m
}

func (m *Memory64K) Set() {
	m.set.Update(true)
}

func (m *Memory64K) Unset() {
	m.set.Update(false)
}

func (m *Memory64K) Enable() {
	m.enable.Update(true)
}

func (m *Memory64K) Disable() {
	m.enable.Update(false)
}

func (m *Memory64K) Update() {
	m.AddressRegister.Update()
	m.rowDecoder.Update(
		m.AddressRegister.Bit(0),
		m.AddressRegister.Bit(1),
		m.AddressRegister.Bit(2),
		m.AddressRegister.Bit(3),
		m.AddressRegister.Bit(4),
		m.AddressRegister.Bit(5),
		m.AddressRegister.Bit(6),
		m.AddressRegister.Bit(7),
	)
	m.colDecoder.Update(
		m.AddressRegister.Bit(8),
		m.AddressRegister.Bit(9),
		m.AddressRegister.Bit(10),
		m.AddressRegister.Bit(11),
		m.AddressRegister.Bit(12),
		m.AddressRegister.Bit(13),
		m.AddressRegister.Bit(14),
		m.AddressRegister.Bit(15),
	)

	var row int = m.rowDecoder.Index()
	var col int = m.colDecoder.Index()

	m.data[row][col].Update(m.set.Value(), m.enable.Value())
}

func (m *Memory64K) String() string {
	var row int = m.rowDecoder.Index()
	var col int = m.colDecoder.Index()

	var builder strings.Builder
	builder.WriteString(fmt.Sprint("Memory\n--------------------------------------\n"))
	builder.WriteString(fmt.Sprintf("RD: %d\tCD: %d\tS: %v\tE: %v\t%s\n", row, col, m.set.Value(), m.enable.Value(), m.AddressRegister.String()))

	for i := 0; i < 256; i++ {
		for j := 0; j < 256; j++ {
			val := m.data[i][j].value.Value()
			if val <= 0x000F {
				builder.WriteString(fmt.Sprintf("0x000%X\t", val))
			} else if val <= 0x00FF {
				builder.WriteString(fmt.Sprintf("0x00%X\t", val))
			} else if val <= 0x0FFF {
				builder.WriteString(fmt.Sprintf("0x0%X\t", val))
			} else {
				builder.WriteString(fmt.Sprintf("0x%X\t", val))
			}
		}
		builder.WriteString(fmt.Sprint("\n"))

	}
	return builder.String()

}

type Cell struct {
	value    component.Register
	andGates [3]gate.ANDGate
}

func NewCell(inputBus, outputBus *component.Bus) *Cell {
	c := new(Cell)
	c.value = *component.NewRegister("", inputBus, outputBus)
	c.andGates[0] = *gate.NewANDGate()
	c.andGates[1] = *gate.NewANDGate()
	c.andGates[2] = *gate.NewANDGate()
	return c
}

func (c *Cell) Update(set bool, enable bool) {
	c.andGates[0].Update(true, true)
	c.andGates[1].Update(c.andGates[0].Value(), set)
	c.andGates[2].Update(c.andGates[0].Value(), enable)

	if c.andGates[1].Value() {
		c.value.Set()
	} else {
		c.value.Unset()
	}

	if c.andGates[2].Value() {
		c.value.Enable()
	} else {
		c.value.Disable()
	}
	c.value.Update()
}
