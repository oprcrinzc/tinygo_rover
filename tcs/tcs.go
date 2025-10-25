package tcs3472

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

type Device struct {
	bus  machine.I2C
	addr uint8
}

func New(i2c machine.I2C) Device {
	return Device{bus: i2c, addr: AddrTCS3472}
}

func (d *Device) write8(reg, val byte) error {
	buf := []byte{commandBit | reg, val}
	return d.bus.Tx(uint16(d.addr), buf, nil)
}

func (d *Device) read8(reg byte) (byte, error) {
	w := []byte{commandBit | reg}
	r := make([]byte, 1)
	err := d.bus.Tx(uint16(d.addr), w, r)
	return r[0], err
}

func (d *Device) read16(reg byte) (uint16, error) {
	w := []byte{commandBit | reg}
	r := make([]byte, 2)
	err := d.bus.Tx(uint16(d.addr), w, r)
	return uint16(r[1])<<8 | uint16(r[0]), err
}

func (d *Device) Enable() error {
	if err := d.write8(regEnable, enablePON); err != nil {
		return err
	}
	time.Sleep(3 * time.Millisecond)
	return d.write8(regEnable, enablePON|enableAEN)
}

func (d *Device) SetIntegrationTime(val byte) error {
	return d.write8(regATime, val)
}

func (d *Device) SetGain(val byte) error {
	return d.write8(regControl, val)
}

func (d *Device) ReadRawData() (clear, red, green, blue uint16, err error) {
	clear, err = d.read16(regCDATAL)
	if err != nil {
		return clear, red, green, blue, err
	}
	red, err = d.read16(regRDATAL)
	if err != nil {
		return clear, red, green, blue, err
	}
	green, err = d.read16(regGDATAL)
	if err != nil {
		return clear, red, green, blue, err
	}
	blue, err = d.read16(regBDATAL)
	return clear, red, green, blue, err
}
