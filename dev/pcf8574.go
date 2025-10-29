package devices

import (
	"machine"
)

type (
	Pcf8574 Device
)

func (p *Pcf8574) send() {
	err := p.i2c.Tx(p.addr, p.payload, nil)
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

func (p *Pcf8574) Set(value uint8) *Pcf8574 {
	if value >= 0 && value <= 255 {
		p.payload[0] = value
		p.send()
	}
	return p
}

func (p *Pcf8574) Read() {
	// err := p.i2c.ReadRegister(uint8(p.addr), 0x00, p.read)
	err := p.i2c.Tx(p.addr, nil, p.read)
	// err := p.i2c.ReadRegister(p.addr)
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

func NewPcf8574(addr uint16, cfg machine.I2CConfig, channel uint8) *Pcf8574 {
	var p *Pcf8574 = new(Pcf8574)

	p.read = make([]byte, 1)
	p.payload = make([]byte, 1)
	p.payload[0] = 0x00
	p.addr = addr

	err := setI2cIfhave((*Device)(p), cfg, channel)

	if err != 0 {
		println("(pcf8574) NewPcf8574() : set i2c config err ", err)
		return nil
	}

	println("(pcf8574) Loaded")

	return p
}

// before TEST
