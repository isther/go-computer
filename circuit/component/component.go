package component

type Component interface {
	SetInputWire(int, bool)
	GetOutputWire(int) bool
}

func setInputByUint16(c Component, value uint16) {
	var x = 0
	for i := BUS_WIDTH - 1; i >= 0; i-- {
		r := (value & (1 << uint16(x)))
		if r != 0 {
			c.SetInputWire(i, true)
		} else {
			c.SetInputWire(i, false)
		}
		x++
	}
}

func getOutput(c Component) uint16 {
	var (
		result uint16
		x      uint16
	)
	for i := BUS_WIDTH - 1; i >= 0; i-- {
		if c.GetOutputWire(i) {
			result = result | (1 << uint16(x))
		} else {
			result = result & ^(1 << uint16(x))
		}
		x++
	}
	return result
}
