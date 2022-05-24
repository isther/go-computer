package alu

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/isther/go-computer/circuit/component"
)

var inputBusA *component.Bus = component.NewBus(component.BUS_WIDTH)
var inputBusB *component.Bus = component.NewBus(component.BUS_WIDTH)
var outputBus *component.Bus = component.NewBus(component.BUS_WIDTH)

func TestAluNOT(t *testing.T) {
	tests := []struct {
		name      string
		inputA    uint16
		wantValue uint16
	}{
		{"1", 0x0000, 0xFFFF},
		{"2", 0xFFFF, 0x0000},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			inputBusA.SetValue(tt.inputA)
			alu := NewALU(inputBusA, inputBusB, outputBus)
			setOp(alu, NOT)

			alu.Update()

			if !reflect.DeepEqual(outputBus.Value(), tt.wantValue) {
				t.Errorf("AluNOT-%s result: %v expect: %v", tt.name, outputBus.Value(), tt.wantValue)
			}
		})
	}
}

func TestAluAND(t *testing.T) {
	tests := []struct {
		name      string
		inputA    uint16
		inputB    uint16
		wantValue uint16
	}{
		{"1", 0x0000, 0x0000, 0x0000},
		{"2", 0x0000, 0xFF00, 0x0000},
		{"3", 0x00FF, 0x0000, 0x0000},
		{"4", 0xFFFF, 0xFFFF, 0xFFFF},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			inputBusA.SetValue(tt.inputA)
			inputBusB.SetValue(tt.inputB)
			alu := NewALU(inputBusA, inputBusB, outputBus)
			setOp(alu, AND)

			alu.Update()

			if !reflect.DeepEqual(outputBus.Value(), tt.wantValue) {
				t.Errorf("AluAND-%s result: %v expect: %v", tt.name, outputBus.Value(), tt.wantValue)
			}
		})
	}
}

func TestAluOR(t *testing.T) {
	tests := []struct {
		name      string
		inputA    uint16
		inputB    uint16
		wantValue uint16
	}{
		{"1", 0x0000, 0x0000, 0x0000},
		{"2", 0x0000, 0xFF00, 0xFF00},
		{"3", 0x00FF, 0x0000, 0x00FF},
		{"4", 0xFFFF, 0xFFFF, 0xFFFF},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			inputBusA.SetValue(tt.inputA)
			inputBusB.SetValue(tt.inputB)
			alu := NewALU(inputBusA, inputBusB, outputBus)
			setOp(alu, OR)

			alu.Update()

			if !reflect.DeepEqual(outputBus.Value(), tt.wantValue) {
				t.Errorf("AluOR-%s result: %v expect: %v", tt.name, outputBus.Value(), tt.wantValue)
			}
		})
	}
}

func TestAluXOR(t *testing.T) {
	tests := []struct {
		name      string
		inputA    uint16
		inputB    uint16
		wantValue uint16
	}{
		{"1", 0x0000, 0x0000, 0x0000},
		{"2", 0x0000, 0xFF00, 0xFF00},
		{"3", 0x00FF, 0x0000, 0x00FF},
		{"4", 0xFFFF, 0xFFFF, 0x0000},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			inputBusA.SetValue(tt.inputA)
			inputBusB.SetValue(tt.inputB)
			alu := NewALU(inputBusA, inputBusB, outputBus)
			setOp(alu, XOR)

			alu.Update()

			if !reflect.DeepEqual(outputBus.Value(), tt.wantValue) {
				t.Errorf("AluXOR-%s result: %v expect: %v", tt.name, outputBus.Value(), tt.wantValue)
			}
		})
	}
}

func TestSHL(t *testing.T) {
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
			t.Run(tt.name, func(t *testing.T) {
				inputBusA.SetValue(tt.input)
				alu := NewALU(inputBusA, inputBusB, outputBus)
				setOp(alu, SHL)
				alu.carryIn.Update(tt.shiftIn)
				alu.Update()

				if !reflect.DeepEqual(outputBus.Value(), tt.expectOut) || !reflect.DeepEqual(alu.carryOut.Value(), tt.expectShiftOut) {
					t.Errorf("AluSHL-%s", tt.name)
				}
			})
		})
	}
}

func TestSHR(t *testing.T) {
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
			inputBusA.SetValue(tt.input)
			alu := NewALU(inputBusA, inputBusB, outputBus)
			setOp(alu, SHR)
			alu.carryIn.Update(tt.shiftIn)
			alu.Update()

			if !reflect.DeepEqual(outputBus.Value(), tt.expectOut) || !reflect.DeepEqual(alu.carryOut.Value(), tt.expectShiftOut) {
				t.Errorf("AluSHR-%s", tt.name)
			}
		})
	}
}
func TestCMP(t *testing.T) {
	tests := []struct {
		name   string
		inputA uint16
		inputB uint16

		expectEqualOut  bool
		expectLargerOut bool
	}{
		{"1", 0x0000, 0x0000, true, false},
		{"2", 0x00FF, 0x0000, false, true},
		{"3", 0x00FF, 0x00FF, true, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			inputBusA.SetValue(tt.inputA)
			inputBusB.SetValue(tt.inputB)
			alu := NewALU(inputBusA, inputBusB, outputBus)
			setOp(alu, CMP)

			alu.Update()

			if !reflect.DeepEqual(alu.isEqual.Value(), tt.expectEqualOut) || !reflect.DeepEqual(alu.isLarger.Value(), tt.expectLargerOut) {
				t.Errorf("AluCMP-%s", tt.name)
			}
		})
	}
}

func TestAluAdd(t *testing.T) {
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

		{"11", 0xFFFF, 0xFFFF, true, 0xFFFF, true},
		{"12", 0x7777, 0x8888, false, 0xFFFF, false},
		{"13", 0x7777, 0x8888, true, 0x0000, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			inputBusA.SetValue(tt.inputA)
			inputBusB.SetValue(tt.inputB)
			alu := NewALU(inputBusA, inputBusB, outputBus)
			setOp(alu, ADD)

			alu.carryIn.Update(tt.carryIn)
			alu.Update()

			if !reflect.DeepEqual(outputBus.Value(), tt.wantSum) {
				t.Errorf("AluADD-%s result: %v expect: %v", tt.name, outputBus.Value(), tt.wantSum)
			}
		})
	}
}

func TestOP(t *testing.T) {
	alu := NewALU(inputBusA, inputBusA, outputBus)
	setOp(alu, ADD)
	fmt.Println(alu.Op)
}

func setOp(a *ALU, value int) {
	value = value & 0xF
	for i := 3; i >= 0; i-- {
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
