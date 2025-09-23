package devices

import (
	"machine"
)

type slave interface {
	Send(data []byte)
	Recv(data, read []byte)
}

type (
	Ikb1z struct {
		i2c     *machine.I2C
		scl     machine.Pin
		sda     machine.Pin
		freq    uint32
		init    bool
		addr    uint16
		read    []byte
		payload []byte
	}
)

func (ikb Ikb1z) send() {
	// println(ikb.payload, ikb.read)
	//	var err error = new(error)
	// println("Hi motor 1")
	err := ikb.i2c.Tx(ikb.addr, ikb.payload, ikb.read)
	if err != nil {
		println("(ikb1z) send() : error=", err)
	}
	// println("Hi motor")
}

// m = 1, 2, 3, 4 motor pin
//
// (1, 100) forward
//
// (128, 227) reverse
func (ikb *Ikb1z) Motor(m uint8, speed int8) *Ikb1z {
	if !ikb.init {
		println("(ikb1z) Motor : ikb1z is not init")
		return ikb
	}
	var reg byte = 0x20

	switch m {
	case 1:
		reg |= 0x01
	case 2:
		reg |= 0x02
	case 3:
		reg |= 0x04
	case 4:
		reg |= 0x08
	}
	ikb.payload = []byte{reg, byte(speed)}
	ikb.send()

	return ikb
}

// m = servo pin
func (ikb *Ikb1z) Servo(m uint8, pos int16) *Ikb1z {
	if !ikb.init {
		println("(ikb1z) Servo : ikb1z is not init")
		return ikb
	}
	var reg byte = 0x40
	switch m {
	case 10:
		reg |= 0x01
	case 11:
		reg |= 0x02
	case 12:
		reg |= 0x04
	case 13:
		reg |= 0x08
	case 14:
		reg |= 0x10
	case 15:
		reg |= 0x20
	}
	ikb.payload = []byte{reg, byte(pos)}
	ikb.send()
	return ikb
}

// pins[0] = scl
//
// pins[1] = sda
func NewIkb1z(pins *PinPack, freq uint32, i2c *machine.I2C) *Ikb1z {
	var ikb *Ikb1z = new(Ikb1z)
	if &pins[0] != nil && &pins[0] != &pins[1] {
		ikb.scl = pins[0]
	}
	if &pins[1] != nil && &pins[1] != &pins[0] {
		ikb.sda = pins[1]
	}

	if freq != 0 {
		ikb.freq = freq
	}
	ikb.read = nil
	ikb.addr = 0x48
	if i2c != nil {
		ikb.i2c = machine.I2C0
		ikb.i2c.Configure(machine.I2CConfig{
			SCL:       machine.Pin(ikb.scl),
			SDA:       machine.Pin(ikb.sda),
			Frequency: ikb.freq,
		})
	} else {
		ikb.i2c = i2c
	}
	//	ikb.i2c = &i2c_
	ikb.init = true
	println("(ikb1z) Loaded")
	return ikb
}

/*
func main() {
	f := NewIkb1z(PinPack{21, 22}, 100e3)

}
*/
