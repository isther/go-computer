package component

type Component interface {
	SetInputWire(int, bool)
	GetOutputWire(int) bool
}
