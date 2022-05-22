package memory

import (
	"reflect"
	"testing"

	"github.com/isther/go-computer/circuit/component"
)

func TestMemory(t *testing.T) {
	tests := []struct {
		name   string
		pos    uint16
		input  uint16
		enable bool
		want   uint16
	}{
		{"1", 0x0000, 0x00FF, true, 0x00FF},
		{"2", 0xFFFF, 0x0000, true, 0x0000},
		{"3", 0x0000, 0x00FF, false, 0x0000},
		{"4", 0xFFFF, 0x0000, false, 0xFFFF},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			bus := component.NewBus(component.BUS_WIDTH)
			m := NewMemory64K(bus)

			// stored in memory
			bus.SetValue(tt.pos)
			m.AddressRegister.Set()
			m.Update()

			m.AddressRegister.Unset()
			m.Update()

			bus.SetValue(tt.input)
			m.Set()
			m.Update()

			m.Unset()
			m.Update()

			// fetch from memory
			m.AddressRegister.Set()
			bus.SetValue(tt.pos)
			m.Update()

			m.AddressRegister.Unset()
			m.Update()

			if tt.enable {
				m.Enable()
				m.Update()
			}

			m.Disable()
			m.Update()

			if !reflect.DeepEqual(bus.Value(), tt.want) {
				t.Errorf("Memory-%s", tt.name)
			}
		})
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
			inputBus.SetValue(tt.input)

			cell := NewCell(inputBus, outputBus)
			cell.Update(tt.set, tt.enable)

			if !reflect.DeepEqual(outputBus.Value(), tt.input) {
				t.Errorf("Cell-%s ", tt.name)
			}
		})
	}

}
