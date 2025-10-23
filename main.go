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
		SCL:       machine.GPIO19,
		SDA:       machine.GPIO18,
		Frequency: 400e3,
	}, 1)

	var pcf8574 *dev.Pcf8574 = dev.NewPcf8574(0x20, machine.I2CConfig{
		SCL:       machine.GPIO17,
		SDA:       machine.GPIO16,
		Frequency: 100e3,
	}, 0)
	// for {
	/*	ikb1z.Servo(10, 0)
		time.Sleep(time.Second * 3)
		ikb1z.Servo(10, 90)
		time.Sleep(time.Second * 3)
		ikb1z.Servo(10, 180)
	*/
	machine.LED.Configure(machine.PinConfig{Mode: machine.PinOutput})

	// machine.GPIO13.Configure(machine.PinConfig{Mode: machine.PinOutput})
	// machine.GPIO13.High()
	ikb1z.Servo(10, 0)
	pcf8574.SetFlag(0x00)
	pcf8574.Set(0xFF)
	println("run . . . ")
	go Blinking()
	for {

		/*	ikb1z.Servo(10, 0)
			ikb1z.Servo(15, 180)
			time.Sleep(time.Second * 1)
			ikb1z.Servo(10, 180)
			ikb1z.Servo(15, 0) */
		time.Sleep(time.Second * 1)

		// ikb1z.Servo(10, 0)
		// time.Sleep(time.Second * 2)
		// ikb1z.Servo(10, 180)

		pcf8574.Read(7)
		print("IR")
		print(pcf8574.Get())
		print("\n")

		// pcf8574.Write(1, 1).Write(1, 0).Write(0, 1)
	}

	//}
}

func Blinking() {
	for {
		machine.LED.High()
		time.Sleep(time.Second)
		machine.LED.Low()
		time.Sleep(time.Second)
	}
}
