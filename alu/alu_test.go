package alu

import (
	"reflect"
	"testing"

	"github.com/isther/go-computer/circuit/component"
)

func TestAluAdd(t *testing.T) {
	var inputBusA *component.Bus = component.NewBus(component.BUS_WIDTH)
	var inputBusB *component.Bus = component.NewBus(component.BUS_WIDTH)
	var outputBus *component.Bus = component.NewBus(component.BUS_WIDTH)
	tests := []struct {
		name      string
		inputA    uint16
		inputB    uint16
		carryIn   bool
		wantSum   uint16
		wantCarry bool
	}{
		{"1", 0x0000, 0x0000, false, 0x0000, false},
		{"2", 0x0000, 0x0001, false, 0x0001, false},
		{"3", 0x0001, 0x0000, false, 0x0001, false},
		{"4", 0x0001, 0x0001, false, 0x0002, false},
		{"5", 0x0001, 0x0001, true, 0x0003, false},
		{"6", 0x000F, 0x0001, false, 0x0010, false},
		{"7", 0xFFFE, 0x0001, false, 0xFFFF, false},
		{"8", 0xFFFE, 0x0001, true, 0x0000, true},
		{"9", 0xFFFF, 0x0001, false, 0x0000, true},
		{"10", 0xFFFF, 0x0000, true, 0x0000, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			inputBusA.SetValue(tt.inputA)
			inputBusB.SetValue(tt.inputB)
			alu := NewALU(inputBusA, inputBusB, outputBus)
			setOp(alu, 0)

			alu.CarryIn.Update(tt.carryIn)
			alu.Update()

			output := getValueOfBus(alu.outputBus)

			if !reflect.DeepEqual(output, tt.wantSum) {
				t.Errorf("AluADD-%s result: %v want: %v", tt.name, output, tt.wantSum)
			}
		})
	}
}
func setOp(a *ALU, value uint16) {
	value = value & 0x10
	for i := 2; i >= 0; i-- {
		r := (value & (1 << byte(i)))
		if r != 0 {
			a.Op[i].Update(true)
		} else {
			a.Op[i].Update(false)
		}
	}
}

func getValueOfComponent(c component.Component, outputBits int) int {
	var x int = 0
	var result int
	for i := (outputBits - 1); i >= 0; i-- {
		if c.GetOutputWire(i) {
			result = result | (1 << uint16(x))
		} else {
			result = result & ^(1 << uint16(x))
		}
		x++
	}
	return result
}

func getValueOfBus(bus *component.Bus) uint16 {
	var x uint16 = 0
	var result uint16
	for i := component.BUS_WIDTH - 1; i >= 0; i-- {
		if bus.GetOutputWire(i) {
			result = result | (1 << x)
		} else {
			result = result & ^(1 << x)
		}
		x++
	}
	return result
}
