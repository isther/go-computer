package memory

import (
	"testing"

	"github.com/isther/go-computer/circuit/component"
)

func TestMemory64KWrite(t *testing.T) {
	bus := component.NewBus(component.BUS_WIDTH)
	m := NewMemory64K(bus)

	var i uint16
	var q uint16 = 0xFFFF
	for i = 0x0000; i < 0xFFFF; i++ {
		m.AddressRegister.Set()
		bus.SetValue(i)
		m.Update()

		m.AddressRegister.Unset()
		m.Update()

		bus.SetValue(q)
		m.Set()
		m.Update()

		m.Unset()
		m.Update()

		q--
	}

	var expected uint16 = 0xFFFF
	for i = 0x0000; i < 0xFFFF; i++ {
		m.AddressRegister.Set()
		bus.SetValue(i)
		m.Update()

		m.AddressRegister.Unset()
		m.Update()

		m.Enable()
		m.Update()

		m.Disable()
		m.Update()

		checkBus(bus, expected)
		expected--
	}
}

func TestMemory64KDoesNotUpdateWhenSetFlagIsOff(t *testing.T) {
	bus := component.NewBus(component.BUS_WIDTH)
	m := NewMemory64K(bus)

	var i uint16
	var q uint16 = 0xFFFF
	for i = 0x0000; i < 0xFFFF; i++ {
		m.AddressRegister.Set()
		bus.SetValue(i)
		m.Update()

		m.AddressRegister.Unset()
		m.Update()

		bus.SetValue(q)

		m.Unset()
		m.Update()

		q--
	}

	var expected uint16 = 0xFFFF
	for i = 0x0000; i < 0xFFFF; i++ {
		m.AddressRegister.Set()
		bus.SetValue(i)
		m.Update()

		m.AddressRegister.Unset()
		m.Update()

		m.Enable()
		m.Update()

		m.Disable()
		m.Update()

		checkBus(bus, expected)
	}
}

func TestCell(t *testing.T) {
	tests := []struct {
		name        string
		set, enable bool
		input       uint16
	}{
		{"1", true, true, 0x0000},
		{"2", true, true, 0xFFFF},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			inputBus := component.NewBus(component.BUS_WIDTH)
			outputBus := component.NewBus(component.BUS_WIDTH)
			setBus(inputBus, tt.input)
			cell := NewCell(inputBus, outputBus)
			cell.Update(tt.set, tt.enable)
			if !checkBus(outputBus, tt.input) {
				t.Errorf("Cell-%s ", tt.name)
			}
		})
	}

}

func setBus(b *component.Bus, value uint16) {
	for i := component.BUS_WIDTH - 1; i >= 0; i-- {
		r := (value & (1 << uint16(i)))
		if r != 0 {
			b.SetInputWire(i, true)
		} else {
			b.SetInputWire(i, false)
		}
	}
}

func checkBus(b *component.Bus, expected uint16) bool {
	var result uint16
	for i := component.BUS_WIDTH - 1; i >= 0; i-- {
		if b.GetOutputWire(i) {
			result = result | (1 << uint16(i))
		} else {
			result = result & ^(1 << uint16(i))
		}
	}
	return result == expected
}
