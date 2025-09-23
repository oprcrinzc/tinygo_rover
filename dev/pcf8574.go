package devices

import (
	"machine"
)

type (
	Pcf8574 struct {
		i2c     *machine.I2C
		scl     machine.Pin
		sda     machine.Pin
		freq    uint32
		addr    uint16
		read    []byte
		payload []byte
		flag    uint8
	}
)

func (p *Pcf8574) send() {
	err := p.i2c.Tx(p.addr, p.payload, p.read)
	if err != nil {
		println("(pcf8574) send() : error =", err)
	}
}

func (p *Pcf8574) Write(pin, value uint8) *Pcf8574 {
	if pin > 7 || pin < 0 {
		return p
	}
	if value == 1 {
		p.payload[0] |= (1 << pin)
	}
	if value == 0 {
		p.payload[0] ^= (1 << pin)
	}
	if (p.flag>>0)&0x01 == 1 {
		println("(pcf8574) Write() : payload=", p.payload[0])
	}
	return p
}

func (p *Pcf8574) Read(pin uint8) {
	err := p.i2c.ReadRegister(uint8(p.addr), 0x00, p.read)
	if err != nil {
		println("(Pcf8574) Read() : error = ", err)
	}
}

func (p *Pcf8574) Get() byte {
	return p.read[0]
}

func (p *Pcf8574) SetFlag(f uint8) {
	p.flag = f
}

func NewPcf8574(addr uint16, pins *PinPack, freq uint32, i2c *machine.I2C) *Pcf8574 {
	var p *Pcf8574 = new(Pcf8574)
	if &pins[0] != nil && &pins[0] != &pins[1] {
		p.scl = pins[0]
	}
	if &pins[1] != nil && &pins[1] != &pins[0] {
		p.sda = pins[1]
	}
	if freq != 0 {
		p.freq = freq
	}
	p.read = make([]byte, 1)
	p.payload = make([]byte, 1)
	p.payload[0] = 0x00
	p.addr = addr
	if i2c != nil {
		p.i2c = machine.I2C0
		p.i2c.Configure(machine.I2CConfig{
			SCL:       p.scl,
			SDA:       p.sda,
			Frequency: p.freq,
		})
	} else {
		p.i2c = i2c
	}
	println("(pcf8574) Loaded")

	return p
}
