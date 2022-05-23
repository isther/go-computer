package component

import "github.com/isther/go-computer/circuit"

type LeftShifter struct {
	inputs   [BUS_WIDTH]circuit.Wire
	outputs  [BUS_WIDTH]circuit.Wire
	shiftIn  circuit.Wire
	shiftOut circuit.Wire
	next     Component
}

func NewLeftShifter() *LeftShifter {
	return new(LeftShifter)
}

func (l *LeftShifter) ConnectOutput(b Component) {
	l.next = b
}

func (l *LeftShifter) GetOutputWire(index int) bool {
	return l.outputs[index].Value()
}

func (l *LeftShifter) SetInputWire(index int, value bool) {
	l.inputs[index].Update(value)
}

func (l *LeftShifter) ShiftOut() bool {
	return l.shiftOut.Value()
}

func (l *LeftShifter) Update(shiftIn bool) {
	l.shiftIn.Update(shiftIn)
	l.shiftOut.Update(l.inputs[0].Value())
	l.outputs[0].Update(l.inputs[1].Value())
	l.outputs[1].Update(l.inputs[2].Value())
	l.outputs[2].Update(l.inputs[3].Value())
	l.outputs[3].Update(l.inputs[4].Value())
	l.outputs[4].Update(l.inputs[5].Value())
	l.outputs[5].Update(l.inputs[6].Value())
	l.outputs[6].Update(l.inputs[7].Value())
	l.outputs[7].Update(l.inputs[8].Value())
	l.outputs[8].Update(l.inputs[9].Value())
	l.outputs[9].Update(l.inputs[10].Value())
	l.outputs[10].Update(l.inputs[11].Value())
	l.outputs[11].Update(l.inputs[12].Value())
	l.outputs[12].Update(l.inputs[13].Value())
	l.outputs[13].Update(l.inputs[14].Value())
	l.outputs[14].Update(l.inputs[15].Value())
	l.outputs[15].Update(l.shiftIn.Value())
}

type RightShifter struct {
	inputs   [BUS_WIDTH]circuit.Wire
	outputs  [BUS_WIDTH]circuit.Wire
	shiftIn  circuit.Wire
	shiftOut circuit.Wire
	next     Component
}

func NewRightShifter() *RightShifter {
	return new(RightShifter)
}

func (r *RightShifter) ConnectOutput(b Component) {
	r.next = b
}

func (r *RightShifter) GetOutputWire(index int) bool {
	return r.outputs[index].Value()
}

func (r *RightShifter) SetInputWire(index int, value bool) {
	r.inputs[index].Update(value)
}

func (r *RightShifter) ShiftOut() bool {
	return r.shiftOut.Value()
}

func (r *RightShifter) Update(shiftIn bool) {
	r.shiftIn.Update(shiftIn)
	r.outputs[0].Update(r.shiftIn.Value())
	r.outputs[1].Update(r.inputs[0].Value())
	r.outputs[2].Update(r.inputs[1].Value())
	r.outputs[3].Update(r.inputs[2].Value())
	r.outputs[4].Update(r.inputs[3].Value())
	r.outputs[5].Update(r.inputs[4].Value())
	r.outputs[6].Update(r.inputs[5].Value())
	r.outputs[7].Update(r.inputs[6].Value())
	r.outputs[8].Update(r.inputs[7].Value())
	r.outputs[9].Update(r.inputs[8].Value())
	r.outputs[10].Update(r.inputs[9].Value())
	r.outputs[11].Update(r.inputs[10].Value())
	r.outputs[12].Update(r.inputs[11].Value())
	r.outputs[13].Update(r.inputs[12].Value())
	r.outputs[14].Update(r.inputs[13].Value())
	r.outputs[15].Update(r.inputs[14].Value())
	r.shiftOut.Update(r.inputs[15].Value())
}
