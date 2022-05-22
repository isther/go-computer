package component

const BUS_WIDTH = 16

type Component interface {
	SetInputWire(int, bool)
	GetOutputWire(int) bool
}
