package devices

import "machine"

type (
	tcs34725 struct {
		i2c machine.I2C
		scl machine.Pin
		sda machine.Pin
	}
)
