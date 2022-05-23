package component

import (
	"reflect"
	"testing"
)

func TestLeftShifter(t *testing.T) {
	tests := []struct {
		name    string
		input   uint16
		shiftIn bool

		expectOut      uint16
		expectShiftOut bool
	}{
		{"1", 0x0000, false, 0x0000, false},
		{"2", 0x8000, false, 0x0000, true},
		{"3", 0xFFFF, false, 0xFFFE, true},
		{"4", 0x0000, true, 0x0001, false},
		{"5", 0x8000, true, 0x0001, true},
		{"6", 0xFFFF, true, 0xFFFF, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			leftShifter := NewLeftShifter()
			setWireOn16(leftShifter, tt.input)
			leftShifter.Update(tt.shiftIn)

			if !reflect.DeepEqual(getComponentOutput(leftShifter), tt.expectOut) || !reflect.DeepEqual(leftShifter.ShiftOut(), tt.expectShiftOut) {
				t.Errorf("LeftShifter-%s", tt.name)
			}
		})
	}
}

func TestRightShifter(t *testing.T) {
	tests := []struct {
		name    string
		input   uint16
		shiftIn bool

		expectOut      uint16
		expectShiftOut bool
	}{
		{"1", 0x0000, false, 0x0000, false},
		{"2", 0x0001, false, 0x0000, true},
		{"3", 0x8000, false, 0x4000, false},
		{"4", 0xFFFF, false, 0x7FFF, true},
		{"5", 0x0000, true, 0x8000, false},
		{"6", 0x8000, true, 0xC000, false},
		{"7", 0x4AAA, true, 0xA555, false},
		{"8", 0xFFFF, true, 0xFFFF, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rightShifter := NewRightShifter()
			setWireOn16(rightShifter, tt.input)
			rightShifter.Update(tt.shiftIn)

			if !reflect.DeepEqual(getComponentOutput(rightShifter), tt.expectOut) || !reflect.DeepEqual(rightShifter.ShiftOut(), tt.expectShiftOut) {
				t.Errorf("LeftShifter-%s", tt.name)
			}
		})
	}
}
