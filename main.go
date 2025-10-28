package main

import (
	"image/color"
	"machine"
	"time"

	//"tinygo.org/x/machine"

	dev "rover/dev"

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

/*
var tcs34725 *dev.Tcs34725 = dev.NewTcs34725(machine.I2CConfig{
	SCL:       machine.GPIO1,
	SDA:       machine.GPIO0,
	Frequency: 100e3,
}, 0)
*/

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

	//	tcs34725.Init()

	go Blinking()
	// go GetColor()
	// go ReadLine()
	//	go Servo()
	//	go Oled()

	// loop for life
	//

	/*const (
		AddrTCS3472 = 0x29
		commandBit  = 0x80
		regEnable   = 0x00
		regATime    = 0x01
		regControl  = 0x0F
		regCDATAL   = 0x14
		regRDATAL   = 0x16
		regGDATAL   = 0x18
		regBDATAL   = 0x1A
		enablePON   = 0x01
		enableAEN   = 0x02
	)
	tcs34725.Write8(regEnable, enablePON)
	time.Sleep(3 * time.Millisecond)
	tcs34725.Write8(regEnable, enablePON|enableAEN)
	tcs34725.Write8(regATime, 0xD5)
	tcs34725.Write8(regControl, 0x01)

	time.Sleep(200 * time.Millisecond)
	*/
	//	ikb1z.Motor(1, 100)

	speed_1 := 100
	speed_2 := 100

	for {

		ikb1z.Motor(2, int8(speed_1))
		time.Sleep(time.Millisecond * 20)
		ikb1z.Motor(1, int8(speed_2))

		pcf8574.Read(7)
		lines := pcf8574.Get()
		lines = lines >> 3
		if lines == 0b0 {
			speed_1 = 0
			speed_2 = 0
		} else if lines == 0b01000 {
		} else {
			speed_1 = 0
			speed_2 = 0
		}

		//	println("Hi")
		/*
			r, g, b, c := tcs34725.GetRGBC()
			println("C:", c, "R:", r, "G:", g, "B:", b)
			time.Sleep(500 * time.Millisecond)
		*/
		time.Sleep(time.Millisecond * 200)
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

/*
func GetColor() {
	for {
		println("Hi")
		r, g, b, c := tcs34725.GetRGBC()
		println("C:", c, "R:", r, "G:", g, "B:", b)
		time.Sleep(500 * time.Millisecond)

	}
}*/

// tracking line whith pcf8574 pin 7, 4
func LineTrack() {
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
