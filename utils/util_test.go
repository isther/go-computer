package utils

import (
	"reflect"
	"testing"
)

func TestUint16ToBinary(t *testing.T) {
	tests := []struct {
		name  string
		value uint16
		want  [16]bool
	}{
		{
			name:  "1",
			value: 65535,
			want: [16]bool{
				true, true, true, true, true, true, true, true,
				true, true, true, true, true, true, true, true,
			},
		}, {
			name:  "2",
			value: 0,
			want: [16]bool{
				false, false, false, false, false, false, false, false,
				false, false, false, false, false, false, false, false,
			},
		}, {
			name:  "3",
			value: 32768,
			want: [16]bool{
				true, false, false, false, false, false, false, false,
				false, false, false, false, false, false, false, false,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if !reflect.DeepEqual(uint16ToBinary(tt.value), tt.want) {
				t.Errorf("BinaryToUint16-%s", tt.name)
				t.Log(uint16ToBinary(tt.value))
			}
		})
	}
}

func TestBinaryToUint16(t *testing.T) {
	tests := []struct {
		name string
		b    [16]bool
		want uint16
	}{
		{
			name: "1",
			b: [16]bool{
				true, true, true, true, true, true, true, true,
				true, true, true, true, true, true, true, true,
			},
			want: 65535,
		}, {
			name: "2",
			b: [16]bool{
				false, false, false, false, false, false, false, false,
				false, false, false, false, false, false, false, false,
			},
			want: 0,
		}, {
			name: "3",
			b: [16]bool{
				true, false, false, false, false, false, false, false,
				false, false, false, false, false, false, false, false,
			},
			want: 32768,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if !reflect.DeepEqual(binaryToUint16(tt.b), tt.want) {
				t.Errorf("BinaryToUint16-%s", tt.name)
			}
		})
	}
}

func TestUint16ToString(t *testing.T) {
	tests := []struct {
		name string
		x    uint16
		want string
	}{
		{"1", 0, "0x0000"},
		{"2", 1, "0x0001"},
		{"3", 0xF, "0x000F"},
		{"4", 0xFF, "0x00FF"},
		{"5", 0xFFF, "0x0FFF"},
		{"6", 0xFFFF, "0xFFFF"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			str := Uint16ToString(tt.x)
			if !reflect.DeepEqual(str, tt.want) {
				t.Errorf("Uint16ToString-%s result: %s want: %s", tt.name, str, tt.want)
			}
		})
	}
}
