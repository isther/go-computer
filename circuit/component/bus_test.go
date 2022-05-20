package component

import (
	"reflect"
	"testing"
)

func TestBus(t *testing.T) {
	tests := []struct {
		name      string
		value     uint16
		wantValue uint16
		wantStr   string
	}{
		{
			name:      "1",
			value:     0x0000,
			wantValue: 0x0000,
			wantStr:   "0000000000000000",
		}, {
			name:      "2",
			value:     0x0001,
			wantValue: 0x0001,
			wantStr:   "0000000000000001",
		}, {
			name:      "3",
			value:     0x00FF,
			wantValue: 0x00FF,
			wantStr:   "0000000011111111",
		}, {
			name:      "4",
			value:     0xFFFF,
			wantValue: 0xFFFF,
			wantStr:   "1111111111111111",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			bus := NewBus(16)
			bus.SetValue(tt.value)

			if !reflect.DeepEqual(bus.Value(), tt.wantValue) || !reflect.DeepEqual(bus.String(), tt.wantStr) {
				t.Errorf("Bus-%s result: %v %s want: %v %s", tt.name, bus.Value(), bus.String(), tt.wantValue, tt.wantStr)
			}
		})
	}
}
