package component

import (
	"reflect"
	"testing"
)

func TestXORGates(t *testing.T) {
	tests := []struct {
		name  string
		input []bool
		want  []bool
	}{
		{"1", []bool{
			false, false, false, false, false, false, false, true,
			false, false, false, false, false, false, false, false,

			false, false, false, false, false, false, false, true,
			true, true, true, true, true, true, true, true,
		}, []bool{
			false, false, false, false, false, false, false, false,
			true, true, true, true, true, true, true, true,
		}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			xorGates := NewXORGates()
			initInputs(xorGates, tt.input)
			xorGates.Update()

			if !reflect.DeepEqual(getOutputs(xorGates), tt.want) {
				t.Errorf("XORGates-%s", tt.name)
			}
		})
	}
}

func TestNOTGates(t *testing.T) {
	tests := []struct {
		name  string
		input []bool
		want  []bool
	}{
		{"1", []bool{
			false, false, false, false, false, false, false, true,
			true, true, true, true, true, true, true, false,
		}, []bool{
			true, true, true, true, true, true, true, false,
			false, false, false, false, false, false, false, true,
		}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			notGates := NewNOTGates()
			initInputs(notGates, tt.input)
			notGates.Update()

			if !reflect.DeepEqual(getOutputs(notGates), tt.want) {
				t.Errorf("NOTGates-%s", tt.name)
			}
		})
	}
}

func TestORGates(t *testing.T) {
	tests := []struct {
		name  string
		input []bool
		want  []bool
	}{
		{"1", []bool{
			false, false, false, false, false, false, false, true,
			false, false, false, false, false, false, false, false,

			false, false, false, false, false, false, false, true,
			true, true, true, true, true, true, true, true,
		}, []bool{
			false, false, false, false, false, false, false, true,
			true, true, true, true, true, true, true, true,
		}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			orGates := NewORGates()
			initInputs(orGates, tt.input)
			orGates.Update()

			if !reflect.DeepEqual(getOutputs(orGates), tt.want) {
				t.Errorf("ORGates-%s", tt.name)
			}
		})
	}
}

func TestANDGates(t *testing.T) {

	tests := []struct {
		name  string
		input []bool
		want  []bool
	}{
		{"1", []bool{
			false, false, false, false, false, false, false, true,
			false, false, false, false, false, false, false, false,

			false, false, false, false, false, false, false, true,
			true, true, true, true, true, true, true, true,
		}, []bool{
			false, false, false, false, false, false, false, true,
			false, false, false, false, false, false, false, false,
		}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			andGates := NewANDGates()
			initInputs(andGates, tt.input)
			andGates.Update()

			if !reflect.DeepEqual(getOutputs(andGates), tt.want) {
				t.Errorf("ANDGates-%s", tt.name)
			}
		})
	}
}

func initInputs(c Component, value []bool) {
	for i := 0; i < len(value); i++ {
		c.SetInputWire(i, value[i])
	}
}

func getOutputs(c Component) []bool {
	res := []bool{}
	for i := 0; i < BUS_WIDTH; i++ {
		res = append(res, c.GetOutputWire(i))
	}
	return res
}
