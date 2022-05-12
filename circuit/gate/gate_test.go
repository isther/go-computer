package gate

import (
	"reflect"
	"testing"
)

func TestNOTGate(t *testing.T) {
	tests := []struct {
		name   string
		valueA bool
		want   bool
	}{
		{
			name:   "1",
			valueA: false,
			want:   true,
		}, {
			name:   "2",
			valueA: true,
			want:   false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			notGate := NewNOTGate()

			if value := notGate.Update(tt.valueA).Value(); !reflect.DeepEqual(value, tt.want) {
				t.Errorf("NOTGate-%s: value: %v result: %v want: %v", tt.name, tt.valueA, value, tt.want)
			}
		})
	}
}

func TestORGate(t *testing.T) {
	tests := []struct {
		name   string
		valueA bool
		valueB bool
		want   bool
	}{
		{
			name:   "1",
			valueA: false,
			valueB: false,
			want:   false,
		}, {
			name:   "2",
			valueA: true,
			valueB: false,
			want:   true,
		}, {
			name:   "3",
			valueA: false,
			valueB: true,
			want:   true,
		}, {
			name:   "4",
			valueA: true,
			valueB: true,
			want:   true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			orGate := NewORGate()

			if value := orGate.Update(tt.valueA, tt.valueB).Value(); !reflect.DeepEqual(value, tt.want) {
				t.Errorf("ORGate-%s: value: %v %v result: %v want: %v", tt.name, tt.valueA, tt.valueB, value, tt.want)
			}
		})
	}
}

func TestNORGate(t *testing.T) {
	tests := []struct {
		name   string
		valueA bool
		valueB bool
		want   bool
	}{
		{
			name:   "1",
			valueA: false,
			valueB: false,
			want:   true,
		}, {
			name:   "2",
			valueA: true,
			valueB: false,
			want:   false,
		}, {
			name:   "3",
			valueA: false,
			valueB: true,
			want:   false,
		}, {
			name:   "4",
			valueA: true,
			valueB: true,
			want:   false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			norGate := NewNORGate()

			if value := norGate.Update(tt.valueA, tt.valueB).Value(); !reflect.DeepEqual(value, tt.want) {
				t.Errorf("NORGate-%s: value: %v %v result: %v want: %v", tt.name, tt.valueA, tt.valueB, value, tt.want)
			}
		})
	}
}

func TestANDGate(t *testing.T) {
	tests := []struct {
		name   string
		valueA bool
		valueB bool
		want   bool
	}{
		{
			name:   "1",
			valueA: false,
			valueB: false,
			want:   false,
		}, {
			name:   "2",
			valueA: true,
			valueB: false,
			want:   false,
		}, {
			name:   "3",
			valueA: false,
			valueB: true,
			want:   false,
		}, {
			name:   "4",
			valueA: true,
			valueB: true,
			want:   true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			andGate := NewANDGate()

			if value := andGate.Update(tt.valueA, tt.valueB).Value(); !reflect.DeepEqual(value, tt.want) {
				t.Errorf("ANDGate-%s: value: %v %v result: %v want: %v", tt.name, tt.valueA, tt.valueB, value, tt.want)
			}
		})
	}
}

func TestNANDGate(t *testing.T) {
	tests := []struct {
		name   string
		valueA bool
		valueB bool
		want   bool
	}{
		{
			name:   "1",
			valueA: false,
			valueB: false,
			want:   true,
		}, {
			name:   "2",
			valueA: true,
			valueB: false,
			want:   true,
		}, {
			name:   "3",
			valueA: false,
			valueB: true,
			want:   true,
		}, {
			name:   "4",
			valueA: true,
			valueB: true,
			want:   false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			nandGate := NewNANDGate()

			if value := nandGate.Update(tt.valueA, tt.valueB).Value(); !reflect.DeepEqual(value, tt.want) {
				t.Errorf("NANDGate-%s: value: %v %v result: %v want: %v", tt.name, tt.valueA, tt.valueB, value, tt.want)
			}
		})
	}
}

func TestXORGate(t *testing.T) {
	tests := []struct {
		name   string
		valueA bool
		valueB bool
		want   bool
	}{
		{
			name:   "1",
			valueA: false,
			valueB: false,
			want:   false,
		}, {
			name:   "2",
			valueA: true,
			valueB: false,
			want:   true,
		}, {
			name:   "3",
			valueA: false,
			valueB: true,
			want:   true,
		}, {
			name:   "4",
			valueA: true,
			valueB: true,
			want:   false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			xorGate := NewXORGate()

			if value := xorGate.Update(tt.valueA, tt.valueB).Value(); !reflect.DeepEqual(value, tt.want) {
				t.Errorf("XORGate-%s: value: %v %v result: %v want: %v", tt.name, tt.valueA, tt.valueB, value, tt.want)
			}
		})
	}
}
