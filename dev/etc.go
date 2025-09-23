package devices

import "machine"

type (
	PinPack [2]machine.Pin
)

func NewPinPack(a, b machine.Pin) *PinPack {
	p_ := PinPack{}
	p_[0] = a
	p_[1] = b
	return &p_
}
