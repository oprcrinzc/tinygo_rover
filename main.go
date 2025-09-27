package main

import (
	"machine"
	"time"

	//"tinygo.org/x/machine"

	dev "rover/dev"
	//
	// "machine"
)

func main() {
	var ikb1z *dev.Ikb1z = dev.NewIkb1z(machine.I2CConfig{
		SCL:       machine.GPIO3,
		SDA:       machine.GPIO2,
		Frequency: 100e3,
	})

	var pcf8574 *dev.Pcf8574 = dev.NewPcf8574(0x20, machine.I2CConfig{
		SCL:       machine.GPIO3,
		SDA:       machine.GPIO2,
		Frequency: 100e3,
	})
	// for {
	/*	ikb1z.Servo(10, 0)
		time.Sleep(time.Second * 3)
		ikb1z.Servo(10, 90)
		time.Sleep(time.Second * 3)
		ikb1z.Servo(10, 180)
	*/

	pcf8574.SetFlag(0x01)
	println("run . . . ")
	for {

		ikb1z.Servo(10, 0)
		ikb1z.Servo(15, 190)
		time.Sleep(time.Second)
		ikb1z.Servo(10, 190)
		ikb1z.Servo(15, 0)
		time.Sleep(time.Second)
		// pcf8574.Write(1, 1).Write(1, 0).Write(0, 1)
	}

	//}
}
