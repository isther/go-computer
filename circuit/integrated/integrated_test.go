package integrated

import (
	"reflect"
	"testing"

	"github.com/isther/go-computer/circuit"
)

func TestHalfAdder(t *testing.T) {
	tests := []struct {
		name string
		a    circuit.Wire
		b    circuit.Wire
		s    bool
		c    bool
	}{
		{
			name: "1",
			a:    *circuit.NewWire("A", false),
			b:    *circuit.NewWire("B", false),
			s:    false,
			c:    false,
		}, {
			name: "2",
			a:    *circuit.NewWire("A", true),
			b:    *circuit.NewWire("B", false),
			s:    true,
			c:    false,
		}, {
			name: "3",
			a:    *circuit.NewWire("A", false),
			b:    *circuit.NewWire("B", true),
			s:    true,
			c:    false,
		}, {
			name: "4",
			a:    *circuit.NewWire("A", true),
			b:    *circuit.NewWire("B", true),
			s:    false,
			c:    true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			halfAdder := NewHalfAdder()
			halfAdder.Update(tt.a.Value(), tt.b.Value())

			if !reflect.DeepEqual(halfAdder.S0(), tt.s) || !reflect.DeepEqual(halfAdder.C0(), tt.c) {
				t.Errorf("HalfAdder-%s: value: %v %v result: %v %v want: %v %v", tt.name, tt.a.Value(), tt.b.Value(), halfAdder.S0(), halfAdder.C0(), tt.s, tt.c)
			}

		})
	}
}

func TestFullAdder(t *testing.T) {
	tests := []struct {
		name string
		c    circuit.Wire
		a    circuit.Wire
		b    circuit.Wire
		si   bool
		ci   bool
	}{
		{
			name: "1",
			c:    *circuit.NewWire("C", false),
			a:    *circuit.NewWire("A", false),
			b:    *circuit.NewWire("B", false),
			si:   false,
			ci:   false,
		}, {
			name: "2",
			c:    *circuit.NewWire("C", true),
			a:    *circuit.NewWire("A", false),
			b:    *circuit.NewWire("B", false),
			si:   true,
			ci:   false,
		}, {
			name: "3",
			c:    *circuit.NewWire("C", false),
			a:    *circuit.NewWire("A", true),
			b:    *circuit.NewWire("B", false),
			si:   true,
			ci:   false,
		}, {
			name: "4",
			c:    *circuit.NewWire("C", true),
			a:    *circuit.NewWire("A", true),
			b:    *circuit.NewWire("B", false),
			si:   false,
			ci:   true,
		}, {
			name: "5",
			c:    *circuit.NewWire("C", false),
			a:    *circuit.NewWire("A", false),
			b:    *circuit.NewWire("B", true),
			si:   true,
			ci:   false,
		}, {
			name: "6",
			c:    *circuit.NewWire("C", true),
			a:    *circuit.NewWire("A", false),
			b:    *circuit.NewWire("B", true),
			si:   false,
			ci:   true,
		}, {
			name: "7",
			c:    *circuit.NewWire("C", false),
			a:    *circuit.NewWire("A", true),
			b:    *circuit.NewWire("B", true),
			si:   false,
			ci:   true,
		}, {
			name: "8",
			c:    *circuit.NewWire("C", true),
			a:    *circuit.NewWire("A", true),
			b:    *circuit.NewWire("B", true),
			si:   true,
			ci:   true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fullAdder := NewFullAdder()
			fullAdder.Update(tt.a.Value(), tt.b.Value(), tt.c.Value())

			if !reflect.DeepEqual(fullAdder.Si(), tt.si) || !reflect.DeepEqual(fullAdder.Ci(), tt.ci) {
				t.Errorf("FullAdder-%s: value: %v %v %v result: %v %v want: %v %v", tt.name, tt.c.Value(), tt.a.Value(), tt.b.Value(), fullAdder.Si(), fullAdder.Ci(), tt.si, tt.ci)
			}

		})
	}
}
