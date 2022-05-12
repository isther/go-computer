package circuit

type Wire struct {
	value bool
}

func NewWire(value bool) *Wire {
	return &Wire{
		value: value,
	}
}

func (w *Wire) Update(value bool) {
	w.value = value
}

func (w *Wire) Value() bool {
	return w.value
}
