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
			want:  "0000000000000000",
		}, {
			name:  "2",
			value: 1,
			want:  "0000000000000001",
		}, {
			name:  "3",
			value: 65535,
			want:  "1111111111111111",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			bus := NewBus(16)
			bus.SetValue(tt.value)

			if !reflect.DeepEqual(bus.String(), tt.want) {
				t.Errorf("Bus-%s result: %s want: %s", tt.name, bus.String(), tt.want)
			}
		})
	}
}
