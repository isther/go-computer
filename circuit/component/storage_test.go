package component

import (
	"reflect"
	"testing"
)

func TestWord(t *testing.T) {
	tests := []struct {
		name  string
		value uint16
		s     bool
		want  uint16
	}{
		{"1", 0x0000, true, 0x0000},
		{"2", 0x00FF, true, 0x00FF},
		{"3", 0xFFFF, true, 0xFFFF},
		{"4", 0xFFFF, false, 0x0000},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tempComponent := new(TempComponent)
			setInputByUint16(tempComponent, tt.value)

			word := NewWord()
			tempComponent.ConnectOutput(word)
			tempComponent.Update()
			word.Update(tt.s)

			if !reflect.DeepEqual(word.Value(), tt.want) {
				t.Log(word.outputs)
				t.Errorf("Word-%s result: %v want: %v", tt.name, word.Value(), tt.want)
			}
		})
	}
}

func TestBit(t *testing.T) {
	tests := []struct {
		name string
		s, e bool
		want bool
	}{
		{"1", true, true, true},
		{"2", false, true, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			bit := NewBit()
			bit.Update(tt.s, tt.e)
			if !reflect.DeepEqual(bit.Value(), tt.want) {
				t.Errorf("Bit-%s s: %v e: %v result: %v want: %v", tt.name, tt.s, tt.e, bit.Value(), tt.want)
			}
		})
	}
}
