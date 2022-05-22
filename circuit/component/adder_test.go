package component

import (
	"reflect"
	"testing"
)

func TestFullAdder(t *testing.T) {
	tests := []struct {
		name        string
		a, b, carry bool
		si          bool
		ci          bool
	}{
		{"1", false, false, false, false, false},
		{"2", true, false, false, true, false},
		{"3", false, true, false, true, false},
		{"4", true, true, false, false, true},
		{"5", false, false, true, true, false},
		{"6", true, false, true, false, true},
		{"7", false, true, true, false, true},
		{"8", true, true, true, true, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fullAdder := NewFullAdder()
			fullAdder.Update(tt.a, tt.b, tt.carry)

			if !reflect.DeepEqual(fullAdder.Sum(), tt.si) || !reflect.DeepEqual(fullAdder.Carry(), tt.ci) {
				t.Errorf("FullAdder-%s: value: %v %v %v result: %v %v expect: %v %v",
					tt.name,
					tt.a, tt.b, tt.carry,
					fullAdder.Sum(), fullAdder.Carry(),
					tt.si, tt.ci)
			}

		})
	}
}

func Test16BitAdder(t *testing.T) {
	tests := []struct {
		name    string
		inputA  uint16
		inputB  uint16
		carryIn bool
		expect  uint16
		carry   bool
	}{
		{"1", 0, 0, false, 0, false},
		{"2", 1, 0, false, 1, false},
		{"3", 0, 1, false, 1, false},
		{"4", 1, 1, false, 2, false},
		{"5", 0, 0, true, 1, false},
		{"6", 32768, 32768, false, 0, true},
		{"7", 32769, 32768, false, 1, true},
		{"8", 65535, 2, false, 1, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			adder16Bit := NewAdder16Bit()
			setWireOn16x2(adder16Bit, tt.inputA, tt.inputB)
			adder16Bit.SetCarryIn(tt.carryIn).Update()

			if !reflect.DeepEqual(adder16Bit.Carry(), tt.carry) || !reflect.DeepEqual(getComponentOutput(adder16Bit), tt.expect) {
				t.Errorf("Adder16Bit-%s result: %v %v expect: %v %v",
					tt.name,
					getComponentOutput(adder16Bit),
					adder16Bit.Carry(),
					tt.expect,
					tt.carry,
				)
			}

		})
	}
}
