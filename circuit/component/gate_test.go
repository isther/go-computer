package component

import (
	"reflect"
	"testing"
)

func TestXORGates(t *testing.T) {
	tests := []struct {
		name   string
		inputA uint16
		inputB uint16
		expect uint16
	}{
		{"1", 0x0000, 0x0000, 0x0000},
		{"2", 0x0000, 0x00FF, 0x00FF},
		{"3", 0xFF00, 0x0000, 0xFF00},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			xorGates := NewXORGates()
			setWireOn16x2(xorGates, tt.inputA, tt.inputB)
			xorGates.Update()

			if !reflect.DeepEqual(getComponentOutput(xorGates), tt.expect) {
				t.Errorf("XORGates-%s", tt.name)
			}
		})
	}
}

func TestNOTGates(t *testing.T) {
	tests := []struct {
		name   string
		input  uint16
		expect uint16
	}{
		{"1", 0x00FF, 0xFF00},
		{"2", 0xFF00, 0x00FF},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			notGates := NewNOTGates()
			setWireOn16(notGates, tt.input)
			notGates.Update()

			if !reflect.DeepEqual(getComponentOutput(notGates), tt.expect) {
				t.Errorf("NOTGates-%s", tt.name)
			}
		})
	}
}

func TestORGates(t *testing.T) {
	tests := []struct {
		name   string
		inputA uint16
		inputB uint16
		expect uint16
	}{
		{"1", 0x0000, 0x0000, 0x0000},
		{"2", 0x0000, 0x00FF, 0x00FF},
		{"3", 0xFF00, 0x0000, 0xFF00},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			orGates := NewORGates()
			setWireOn16x2(orGates, tt.inputA, tt.inputB)
			orGates.Update()

			if !reflect.DeepEqual(getComponentOutput(orGates), tt.expect) {
				t.Errorf("ORGates-%s: result: %v expect: %v", tt.name, getComponentOutput(orGates), tt.expect)
			}
		})
	}
}

func TestANDGates(t *testing.T) {

	tests := []struct {
		name   string
		inputA uint16
		inputB uint16
		expect uint16
	}{
		{"1", 0x0000, 0x0000, 0x0000},
		{"2", 0x0000, 0x00FF, 0x0000},
		{"3", 0xFF00, 0x0000, 0x0000},
		{"4", 0xFFFF, 0xFFFF, 0xFFFF},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			andGates := NewANDGates()
			setWireOn16x2(andGates, tt.inputA, tt.inputB)
			andGates.Update()

			if !reflect.DeepEqual(getComponentOutput(andGates), tt.expect) {
				t.Errorf("ANDGates-%s", tt.name)
			}
		})
	}
}
