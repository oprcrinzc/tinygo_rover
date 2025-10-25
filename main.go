package main

import (
	"image/color"
	"machine"
	"time"

	//"tinygo.org/x/machine"

	dev "rover/dev"
	tcs3472 "rover/tcs"

	"tinygo.org/x/drivers/ssd1306"
	//
	// "machine"
	//
)

var ikb1z *dev.Ikb1z = dev.NewIkb1z(machine.I2CConfig{
	SCL:       machine.GPIO15,
	SDA:       machine.GPIO14,
	Frequency: 400e3,
}, 1)

var pcf8574 *dev.Pcf8574 = dev.NewPcf8574(0x20, machine.I2CConfig{
	SCL:       machine.GPIO17,
	SDA:       machine.GPIO16,
	Frequency: 100e3,
}, 0)

func main() {
	// for {
	/*	ikb1z.Servo(10, 0)
		time.Sleep(time.Second * 3)
		ikb1z.Servo(10, 90)
		time.Sleep(time.Second * 3)
		ikb1z.Servo(10, 180)
	*/
	machine.LED.Configure(machine.PinConfig{Mode: machine.PinOutput})

	// 3v3 out for ikb1z
	machine.GPIO20.Configure(machine.PinConfig{Mode: machine.PinOutput})
	machine.GPIO20.High()

	ikb1z.Servo(10, 0)
	pcf8574.SetFlag(0x00)
	pcf8574.Set(0xFF)
	println("run . . . ")

	go Blinking()
	go ReadLine()
	go Servo()
	//	go Oled()

	// loop for life
	//

	i2c := *machine.I2C0
	i2c.Configure(machine.I2CConfig{Frequency: 400_000})

	sensor := tcs3472.New(i2c)
	sensor.Enable()
	sensor.SetIntegrationTime(0xFF) // longest integration time
	sensor.SetGain(0x01)            // 4x gain

	for {
		c, r, g, b, _ := sensor.ReadRawData()
		println("C:", c, "R:", r, "G:", g, "B:", b)
		time.Sleep(500 * time.Millisecond)
	}

	/*
		for {
			ikb1z.Motor(1, 100)
			time.Sleep(time.Second * 2)
			// ikb1z.Servo(10, 180)
			ikb1z.Motor(1, 0)
			time.Sleep(time.Second)
			// pcf8574.Write(1, 1).Write(1, 0).Write(0, 1)
			//}
		}
	*/
}

func Oled() {
	var ssd1306_i2c1 *machine.I2C = machine.I2C1
	ssd1306_i2c1.Configure(machine.I2CConfig{
		SCL: machine.GPIO19,
		SDA: machine.GPIO18,
	})
	var oled *ssd1306.Device = ssd1306.NewI2C(machine.I2C1)

	oled.Configure(ssd1306.Config{Address: 0x3c, Width: 128, Height: 64})
	oled.ClearBuffer()
	for {
		oled.ClearDisplay()
		oled.SetPixel(0, 0, color.RGBA{255, 255, 255, 255})
		// oled.FillRectangle(20, 20, 20, 20, color.RGBA{255, 255, 255, 255})

		oled.Display()
	}
}

func Servo() {
	for {
		ikb1z.Servo(10, 0)
		ikb1z.Servo(15, 180)
		ikb1z.Motor(1, 100)
		time.Sleep(time.Second * 1)
		ikb1z.Servo(10, 180)
		ikb1z.Servo(15, 0)
		time.Sleep(time.Second * 1)
	}
}

func ReadLine() {
	for {
		pcf8574.Read(7)
		print("IR")
		print(pcf8574.Get())
		print("\n")
		time.Sleep(time.Millisecond * 250)
	}
}

func Blinking() {
	for {
		machine.LED.High()
		time.Sleep(time.Second)
		machine.LED.Low()
		time.Sleep(time.Second)
	}
}
