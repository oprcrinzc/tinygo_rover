package devices

import (
	"machine"
)

type (
	Tcs34725 Device
)

func NewTcs34725(i2cConfig machine.I2CConfig, channel uint8) *Tcs34725 {
	var tcs *Tcs34725 = new(Tcs34725)

	err := setI2cIfhave((*Device)(tcs), i2cConfig, channel)
	if err != 0 {
		println("(tcs34725) NewTcs34725() : set i2c config error ", err)
		return nil
	}
	print("(tcs34725) Loaded")
	return tcs
}
