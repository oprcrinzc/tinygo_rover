package main

import (
	"fmt"
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

var tcs34725 *dev.Tcs34725 = dev.NewTcs34725(machine.I2CConfig{
	SCL:       machine.GPIO15,
	SDA:       machine.GPIO14,
	Frequency: 400e3,
}, 1)

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
	machine.GPIO18.Configure(machine.PinConfig{Mode: machine.PinOutput})
	machine.GPIO18.High()

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

		// speed_1, speed_2 = LineTrack()

		speed_1, speed_2 = StupidLineTrack()

		ikb1z.Motor(2, int8(speed_2))
		time.Sleep(time.Millisecond * 20)
		ikb1z.Motor(1, int8(speed_1))

		// time.Sleep(time.Second * 10)

		//break
		/*
			pcf8574.Read()
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
		*/
		//	println("Hi")
		/*
			r, g, b, c := tcs34725.GetRGBC()
			println("C:", c, "R:", r, "G:", g, "B:", b)
			time.Sleep(500 * time.Millisecond)
		*/
		//time.Sleep(time.Millisecond * 200)
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

/*
var (
	Kp float32 = 1
	Ki float32 = 1
	Kd float32 = 1

	I int64 = 0

	Perr  int   = 0
	Ptime int64 = time.Now().UnixMilli()
) */

// tracking line whith pcf8574 pin 6, 4
/*func LineTrack() (L, R int) {
	pcf8574.Read()
	l := pcf8574.Get()
	l = l >> 3
	want := int(0b01010)
	got := int(l)
	err := want - got
	Ntime := time.Now().UnixMilli()
	Dt := Ntime - Ptime
	I = I + int64(err)*(Dt)
	D := float32(int64(err-Perr) / (Dt))
	out := Kp*float32(err) + Ki*float32(I) + Kd*D
	time.Sleep(time.Millisecond * 100)

	// Left := (l >> 6) & 1
	// Right := (l >> 4) & 1
	Ptime = Ntime
	Perr = err
	return L, R
}*/

func StupidLineTrack() (L, R int) {
	pcf8574.Read()
	raw := pcf8574.Get() >> 3 // 5 bits
	err := int(raw) - int(27)

	p := err
	I += int64(err)
	d := err - Perr

	if err == 0 {
		I = 0
		L = 100
		R = 100
		return L, R
	}

	println("StupidLineTrack: ", raw, " ERR: ", err)
	speed := float32(p)*Kp + float32(I)*Ki + float32(d)*Kd

	R = int(100 + speed)
	L = int(100 - speed)
	Perr = err
	if R > 100 {
		R = 100
	}
	if R < -100 {
		R = -100
	}
	if L > 100 {
		L = 100
	}
	if L < -100 {
		L = -100
	}
	return L, R
}

var (
	Kp float32 = 15.0
	Ki float32 = 0.1 // 0.1
	Kd float32 = 5

	I     int64
	Perr  int
	Ptime int64
)

func LineTrack() (L, R int) {
	pcf8574.Read()
	raw := pcf8574.Get() >> 3 // 5 bits

	// Extract individual sensors (0–4)
	s0 := (raw >> 0) & 1
	s1 := (raw >> 1) & 1
	s2 := (raw >> 2) & 1
	s3 := (raw >> 3) & 1
	s4 := (raw >> 4) & 1

	sum := s0 + s1 + s2 + s3 + s4
	if sum == 0 {
		I = 0
		// no line detected — keep previous output or stop
		return 0, 0
	}

	position := (s0*0 + s1*1 + s2*2 + s3*3 + s4*4) / sum
	err := position - 10 // center = 2

	Ntime := time.Now().UnixMilli()
	Dt := Ntime - Ptime
	if Dt == 0 {
		Dt = 1
	}

	I += int64(err) * Dt
	D := float32(int64(int(err)-int(Perr))) / float32(Dt)
	out := Kp*float32(err) + Ki*float32(I) + Kd*D

	// Example motor mix:
	base := 100 // base speed
	R = int(float32(base) - out)
	L = int(float32(base) + out)

	Ptime = Ntime
	Perr = int(err)
	return L, R
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
		pcf8574.Read()
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

func detectColor(c, r, g, b uint16) string {
	/*
		if c < 100 {
			return "black"
		}*/

	R := float32(r) / float32(c) * 255.0
	G := float32(g) / float32(c) * 255.0
	B := float32(b) / float32(c) * 255.0

	fmt.Print("C:", c, "R:", R, "G:", G, "B:", B)

	if R >= 80 && R <= 110 && G >= 87 && G <= 92 && B >= 60 {
		if c < 1000 {
			return "Black"
		}
		return "White"
	}

	/*	if R < 5 && R > B && G > B && R > G /*&& B >= 1.2 && B <= 1.4 {
		return "Yellow"
	}*/
	if R > G && R > B {
		if R >= 160 {
			return "Red"
		}
		return "Yellow"

	}
	if G > R && G > B {
		return "Green"
	}
	if B > G && B > R {
		return "Blue"
	}

	return "..."
}
