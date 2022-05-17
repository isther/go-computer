package component

import (
	"reflect"
	"testing"
)

func TestBus(t *testing.T) {
	tests := []struct {
		name  string
		value uint16
		want  string
	}{
		{
			name:  "1",
			value: 0,
			want:  "00000000",
		}, {
			name:  "2",
			value: 1,
			want:  "00000001",
		}, {
			name:  "3",
			value: 128,
			want:  "10000000",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			bus := NewBus(8)
			bus.SetValue(tt.value)

			if !reflect.DeepEqual(bus.String(), tt.want) {
				t.Errorf("Bus-%s result: %s want: %s", tt.name, bus.String(), tt.want)
			}
		})
	}
}
