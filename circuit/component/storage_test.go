package component

import (
	"reflect"
	"testing"
)

func TestWord(t *testing.T) {
	tests := []struct {
		name   string
		input  uint16
		set    bool
		expect uint16
	}{
		{"1", 0x0000, true, 0x0000},
		{"2", 0x00FF, true, 0x00FF},
		{"3", 0xFFFF, true, 0xFFFF},
		{"4", 0x0000, false, 0x0000},
		{"5", 0x00FF, false, 0x0000},
		{"6", 0xFFFF, false, 0x0000},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tempComponent := new(TempComponent)
			setWireOn16(tempComponent, tt.input)

			word := NewWord()
			tempComponent.ConnectOutput(word)
			tempComponent.Update()
			word.Update(tt.set)

			t.Log(word.inputs[0], tt.set, word.bits[0].Value(), word.outputs[0])

			if !reflect.DeepEqual(word.Value(), tt.expect) {
				t.Errorf("Word-%s result: %v expect: %v", tt.name, word.Value(), tt.expect)
			}
		})
	}
}

func TestBit(t *testing.T) {
	tests := []struct {
		name       string
		init       bool
		input, set bool
		expect     bool
	}{
		{"1", false, true, true, true},
		{"2", false, false, true, false},
		{"3", false, true, false, false},
		{"4", false, false, false, false},
		{"5", true, true, true, true},
		{"6", true, false, true, false},
		{"7", true, true, false, true},
		{"8", true, false, false, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			bit := NewBit()
			bit.Update(tt.init, true)
			bit.Update(tt.input, tt.set)
			if !reflect.DeepEqual(bit.Value(), tt.expect) {
				t.Errorf("Bit-%s s: %v e: %v result: %v expect: %v", tt.name, tt.input, tt.set, bit.Value(), tt.expect)
			}
		})
	}
}
