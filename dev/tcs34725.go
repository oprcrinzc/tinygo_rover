package devices

import (
	"machine"
	"time"
)

const (
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

type (
	Tcs34725 Device
)

func (d *Tcs34725) Write8(reg, val byte) error {
	buf := []byte{commandBit | reg, val}
	return d.i2c.Tx(uint16(d.addr), buf, nil)
}

func (d *Tcs34725) Read16(reg byte) (uint16, error) {
	err := d.i2c.Tx(uint16(d.addr), []byte{commandBit | reg}, nil)
	if err != nil {
		return 0, err
	}
	data := make([]byte, 2)
	err = d.i2c.Tx(uint16(d.addr), nil, data)
	if err != nil {
		return 0, err
	}
	return uint16(data[1])<<8 | uint16(data[0]), nil
}

func (d *Tcs34725) Init() {
	d.Write8(regEnable, enablePON)
	time.Sleep(3 * time.Millisecond)
	d.Write8(regEnable, enablePON|enableAEN)
	d.Write8(regATime, 0xD5)
	d.Write8(regControl, 0x01)
	time.Sleep(200 * time.Millisecond)
}

func (d *Tcs34725) GetRGBC() (r, g, b, c uint16) {
	c, _ = d.Read16(regCDATAL)
	r, _ = d.Read16(regRDATAL)
	g, _ = d.Read16(regGDATAL)
	b, _ = d.Read16(regBDATAL)
	return r, g, b, c
}

func NewTcs34725(i2cConfig machine.I2CConfig, channel uint8) *Tcs34725 {
	var tcs *Tcs34725 = new(Tcs34725)
	tcs.addr = 0x29
	tcs.payload = make([]byte, 1)
	tcs.payload[0] = 0x00
	tcs.read = make([]byte, 1)

	err := setI2cIfhave((*Device)(tcs), i2cConfig, channel)
	if err != 0 {
		println("(tcs34725) NewTcs34725() : set i2c config error ", err)
		return nil
	}
	print("(tcs34725) Loaded")
	return tcs
}
