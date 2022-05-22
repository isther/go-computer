package gate

import (
	"reflect"
	"testing"
)

func TestNOTGate(t *testing.T) {
	tests := []struct {
		name   string
		valueA bool
		expect bool
	}{
		{"1", false, true},
		{"2", true, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			notGate := NewNOTGate()

			if value := notGate.Update(tt.valueA).Value(); !reflect.DeepEqual(value, tt.expect) {
				t.Errorf("NOTGate-%s: value: %v result: %v   %v", tt.name, tt.valueA, value, tt.expect)
			}
		})
	}
}

func TestORGate(t *testing.T) {
	tests := []struct {
		name   string
		valueA bool
		valueB bool
		expect bool
	}{
		{"1", false, false, false},
		{"2", true, false, true},
		{"3", false, true, true},
		{"4", true, true, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			orGate := NewORGate()

			if value := orGate.Update(tt.valueA, tt.valueB).Value(); !reflect.DeepEqual(value, tt.expect) {
				t.Errorf("ORGate-%s: value: %v %v result: %v   %v", tt.name, tt.valueA, tt.valueB, value, tt.expect)
			}
		})
	}
}

func TestNORGate(t *testing.T) {
	tests := []struct {
		name   string
		valueA bool
		valueB bool
		expect bool
	}{
		{"1", false, false, true},
		{"2", true, false, false},
		{"3", false, true, false},
		{"4", true, true, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			norGate := NewNORGate()

			if value := norGate.Update(tt.valueA, tt.valueB).Value(); !reflect.DeepEqual(value, tt.expect) {
				t.Errorf("NORGate-%s: value: %v %v result: %v   %v", tt.name, tt.valueA, tt.valueB, value, tt.expect)
			}
		})
	}
}

func TestANDGate(t *testing.T) {
	tests := []struct {
		name   string
		valueA bool
		valueB bool
		expect bool
	}{
		{"1", false, false, false},
		{"2", true, false, false},
		{"3", false, true, false},
		{"4", true, true, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			andGate := NewANDGate()

			if value := andGate.Update(tt.valueA, tt.valueB).Value(); !reflect.DeepEqual(value, tt.expect) {
				t.Errorf("ANDGate-%s: value: %v %v result: %v   %v", tt.name, tt.valueA, tt.valueB, value, tt.expect)
			}
		})
	}
}

func TestNANDGate(t *testing.T) {
	tests := []struct {
		name   string
		valueA bool
		valueB bool
		expect bool
	}{
		{"1", false, false, true},
		{"2", true, false, true},
		{"3", false, true, true},
		{"4", true, true, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			nandGate := NewNANDGate()

			if value := nandGate.Update(tt.valueA, tt.valueB).Value(); !reflect.DeepEqual(value, tt.expect) {
				t.Errorf("NANDGate-%s: value: %v %v result: %v   %v", tt.name, tt.valueA, tt.valueB, value, tt.expect)
			}
		})
	}
}

func TestXORGate(t *testing.T) {
	tests := []struct {
		name   string
		valueA bool
		valueB bool
		expect bool
	}{
		{"1", false, false, false},
		{"2", true, false, true},
		{"3", false, true, true},
		{"4", true, true, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			xorGate := NewXORGate()

			if value := xorGate.Update(tt.valueA, tt.valueB).Value(); !reflect.DeepEqual(value, tt.expect) {
				t.Errorf("XORGate-%s: value: %v %v result: %v   %v", tt.name, tt.valueA, tt.valueB, value, tt.expect)
			}
		})
	}
}
