package circuit

type Wire struct {
	name  string
	value bool
}

func NewWire(name string, value bool) *Wire {
	return &Wire{
		name:  name,
		value: value,
	}
}

func (w *Wire) Update(value bool) {
	w.value = value
}

func (w *Wire) Value() bool {
	return w.value
}
