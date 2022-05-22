package component

import (
	"reflect"
	"testing"
)

func TestDecoder(t *testing.T) {
	tests := []struct {
		name       string
		a, b, c, d bool
		expect     int
	}{
		{"1", false, false, false, false, 0},
		{"2", false, false, false, true, 1},
		{"3", false, false, true, false, 2},
		{"4", false, false, true, true, 3},
		{"5", false, true, false, false, 4},
		{"6", false, true, false, true, 5},
		{"7", false, true, true, false, 6},
		{"8", false, true, true, true, 7},
		{"9", true, false, false, false, 8},
		{"10", true, false, false, true, 9},
		{"11", true, false, true, false, 10},
		{"12", true, false, true, true, 11},
		{"13", true, true, false, false, 12},
		{"14", true, true, false, true, 13},
		{"15", true, true, true, false, 14},
		{"16", true, true, true, true, 15},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			decoder := NewDecoder4x16()
			decoder.Update(tt.a, tt.b, tt.c, tt.d)
			if !reflect.DeepEqual(decoder.Index(), tt.expect) {
				t.Errorf("Decoder-%s", tt.name)
			}

		})
	}
}
