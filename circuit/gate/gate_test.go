package gate

import (
	"reflect"
	"testing"

	"github.com/isther/computer/circuit"
)

func TestNOTGate(t *testing.T) {
	type fields struct {
		output circuit.Wire
	}
	type args struct {
		valueA bool
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "1",
			args: args{
				valueA: false,
			},
			want: true,
		}, {
			name: "2",
			args: args{
				valueA: true,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			notGate := NewNOTGate()

			if value := notGate.Update(tt.args.valueA).Value(); !reflect.DeepEqual(value, tt.want) {
				t.Errorf("ANDGate-%s: value: %v  want: %v", tt.name, tt.args.valueA, tt.want)
			}
		})
	}
}

func TestORGate(t *testing.T) {
	type fields struct {
		output circuit.Wire
	}
	type args struct {
		valueA bool
		valueB bool
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "1",
			args: args{
				valueA: false,
				valueB: false,
			},
			want: false,
		}, {
			name: "2",
			args: args{
				valueA: true,
				valueB: false,
			},
			want: true,
		}, {
			name: "3",
			args: args{
				valueA: false,
				valueB: true,
			},
			want: true,
		}, {
			name: "4",
			args: args{
				valueA: true,
				valueB: true,
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			orGate := NewORGate()

			if value := orGate.Update(tt.args.valueA, tt.args.valueB).Value(); !reflect.DeepEqual(value, tt.want) {
				t.Errorf("ORGate-%s: value: %v %v want: %v", tt.name, tt.args.valueA, tt.args.valueB, tt.want)
			}
		})
	}
}

func TestANDGate(t *testing.T) {
	type fields struct {
		output circuit.Wire
	}
	type args struct {
		valueA bool
		valueB bool
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "1",
			args: args{
				valueA: false,
				valueB: false,
			},
			want: false,
		}, {
			name: "2",
			args: args{
				valueA: true,
				valueB: false,
			},
			want: false,
		}, {
			name: "3",
			args: args{
				valueA: false,
				valueB: true,
			},
			want: false,
		}, {
			name: "4",
			args: args{
				valueA: true,
				valueB: true,
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			andGate := NewANDGate()

			if value := andGate.Update(tt.args.valueA, tt.args.valueB).Value(); !reflect.DeepEqual(value, tt.want) {
				t.Errorf("ANDGate-%s: value: %v %v want: %v", tt.name, tt.args.valueA, tt.args.valueB, tt.want)
			}
		})
	}
}

func TestNANDGate(t *testing.T) {
	type fields struct {
		output circuit.Wire
	}
	type args struct {
		valueA bool
		valueB bool
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "1",
			args: args{
				valueA: false,
				valueB: false,
			},
			want: true,
		}, {
			name: "2",
			args: args{
				valueA: true,
				valueB: false,
			},
			want: true,
		}, {
			name: "3",
			args: args{
				valueA: false,
				valueB: true,
			},
			want: true,
		}, {
			name: "4",
			args: args{
				valueA: true,
				valueB: true,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			nandGate := NewNANDGate()

			if value := nandGate.Update(tt.args.valueA, tt.args.valueB).Value(); !reflect.DeepEqual(value, tt.want) {
				t.Errorf("ANDGate-%s: value: %v %v want: %v", tt.name, tt.args.valueA, tt.args.valueB, tt.want)
			}
		})
	}
}

func TestXORGate(t *testing.T) {
	type fields struct {
		output circuit.Wire
	}
	type args struct {
		valueA bool
		valueB bool
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "1",
			args: args{
				valueA: false,
				valueB: false,
			},
			want: false,
		}, {
			name: "2",
			args: args{
				valueA: true,
				valueB: false,
			},
			want: true,
		}, {
			name: "3",
			args: args{
				valueA: false,
				valueB: true,
			},
			want: true,
		}, {
			name: "4",
			args: args{
				valueA: true,
				valueB: true,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			xorGate := NewXORGate()

			if value := xorGate.Update(tt.args.valueA, tt.args.valueB).Value(); !reflect.DeepEqual(value, tt.want) {
				t.Errorf("ANDGate-%s: value: %v %v want: %v", tt.name, tt.args.valueA, tt.args.valueB, tt.want)
			}
		})
	}
}
