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
				t.Errorf("FullAdder-%s: value: %v %v %v result: %v %v want: %v %v",
					tt.name,
					tt.a, tt.b, tt.carry,
					fullAdder.Sum(), fullAdder.Carry(),
					tt.si, tt.ci)
			}

		})
	}
}

func Test32BitAdder(t *testing.T) {
	tests := []struct {
		name    string
		a       int
		b       int
		carryIn bool
		want    int
		carry   bool
	}{
		{"1", 0, 0, false, 0, false},
		{"2", 1, 0, false, 1, false},
		{"3", 0, 1, false, 1, false},
		{"4", 1, 1, false, 2, false},
		{"5", 0, 0, true, 1, false},
		{"6", 2147483648, 2147483648, false, 0, true},
		{"7", 2147483649, 2147483648, false, 1, true},
		{"8", 4294967295, 2, false, 1, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			adder32Bit := NewAdder32Bit()
			setWireOn32(adder32Bit, tt.a, tt.b)
			adder32Bit.SetCarryIn(tt.carryIn).Update()

			if !reflect.DeepEqual(adder32Bit.Carry(), tt.carry) || !reflect.DeepEqual(getValueOfOutput(adder32Bit, 32), tt.want) {
				t.Errorf("Adder32Bit-%s result: %v %v want: %v %v",
					tt.name,
					getValueOfOutput(adder32Bit, 32),
					adder32Bit.Carry(),
					tt.want,
					tt.carry,
				)
			}

		})
	}
}

func setWireOn32(c Component, inputA int, inputB int) {
	var x uint32 = 0
	for i := 32 - 1; i >= 0; i-- {
		r := (inputA & (1 << x))
		if r != 0 {
			c.SetInputWire(i, true)
		} else {
			c.SetInputWire(i, false)
		}
		x++
	}

	x = 0
	for i := 64 - 1; i >= 32; i-- {
		r := (inputB & (1 << x))
		if r != 0 {
			c.SetInputWire(i, true)
		} else {
			c.SetInputWire(i, false)
		}
		x++
	}
}

func getValueOfOutput(c Component, outputBits int) int {
	var x int = 0
	var result int
	for i := (outputBits - 1); i >= 0; i-- {
		if c.GetOutputWire(i) {
			result = result | (1 << uint32(x))
		} else {
			result = result & ^(1 << uint32(x))
		}
		x++
	}
	return result
}
