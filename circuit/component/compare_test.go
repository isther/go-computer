package component

import (
	"reflect"
	"testing"
)

func TestComparator(t *testing.T) {
	tests := []struct {
		name   string
		inputA uint16
		inputB uint16

		expectEqualOut  bool
		expectLargerOut bool
		expectOut       uint16
	}{
		{"1", 0x0000, 0x0000, true, false, 0x0000},
		{"2", 0x00FF, 0x0000, false, true, 0x00FF},
		{"3", 0x00FF, 0x00FF, true, false, 0x0000},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			comparator := NewComparator()
			setWireOn16x2(comparator, tt.inputA, tt.inputB)
			comparator.Update()

			if !reflect.DeepEqual(comparator.Equal(), tt.expectEqualOut) ||
				!reflect.DeepEqual(comparator.Larger(), tt.expectLargerOut) ||
				!reflect.DeepEqual(getComponentOutput(comparator), tt.expectOut) {
				t.Errorf("Compare-%s", tt.name)
			}
		})
	}
}

func TestCompare(t *testing.T) {
	tests := []struct {
		name            string
		inputA          bool
		inputB          bool
		equalIn         bool
		largerIn        bool
		expectEqualOut  bool
		expectLargerOut bool
	}{
		{"1", false, false, false, false, false, false},
		{"2", false, false, true, false, true, false},
		{"3", false, true, false, false, false, false},
		{"4", false, true, true, false, false, false},
		{"5", true, false, false, false, false, false},
		{"6", true, false, true, false, false, true},
		{"7", true, true, false, false, false, false},
		{"8", true, true, true, false, true, false},
		{"9", false, false, false, true, false, true},
		{"10", false, false, true, true, true, true},
		{"11", false, true, false, true, false, true},
		{"12", false, true, true, true, false, true},
		{"13", true, false, false, true, false, true},
		{"14", true, false, true, true, false, true},
		{"15", true, true, false, true, false, true},
		{"16", true, true, true, true, true, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			compare := NewCompare()
			compare.Update(tt.inputA, tt.inputB, tt.equalIn, tt.largerIn)

			if !reflect.DeepEqual(compare.Equal(), tt.expectEqualOut) || !reflect.DeepEqual(compare.Larger(), tt.expectLargerOut) {
				t.Errorf("Compare-%s", tt.name)
			}
		})
	}
}
