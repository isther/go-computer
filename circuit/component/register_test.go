package component

import (
	"reflect"
	"testing"
)

func TestRegister(t *testing.T) {
	tests := []struct {
		name        string
		input       uint16
		set         bool
		enable      bool
		wantStorage uint16
		wantOutput  uint16
	}{
		{"set-and-enable-1", 0x0000, true, true, 0x0000, 0x0000},
		{"set-and-enable-2", 0x00FF, true, true, 0x00FF, 0x00FF},
		{"only-set-1", 0x0000, true, false, 0x0000, 0x0000},
		{"only-set-2", 0x00FF, true, false, 0x00FF, 0x0000},
		{"only-enable-1", 0x0000, false, true, 0x0000, 0x0000},
		{"only-enable-2", 0x00FF, false, true, 0x0000, 0x0000},
		{"no-set-and-enable-1", 0x0000, false, false, 0x0000, 0x0000},
		{"no-set-and-enable-2", 0x00FF, false, false, 0x0000, 0x0000},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := NewBus(BUS_WIDTH)
			o := NewBus(BUS_WIDTH)
			// setBus(i, tt.input)
			i.SetValue(tt.input)
			r := NewRegister("r", i, o)
			if tt.set {
				r.Set()
			}
			r.Disable()
			r.Update()

			if !reflect.DeepEqual(r.Value(), tt.wantStorage) {
				t.Errorf("Register-%s: result: %v expect: %v", tt.name, r.Value(), tt.wantStorage)
			}

			if tt.enable {
				r.Enable()
			}
			r.Update()
			if !reflect.DeepEqual(r.outputBus.Value(), tt.wantOutput) {
				t.Errorf("Register-%s: result: %v expect: %v", tt.name, r.outputBus.Value(), tt.wantOutput)
				t.Log(r.enabler.outputs)
			}
		})
	}
}
