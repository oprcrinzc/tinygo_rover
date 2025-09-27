package devices

import (
	"machine"
)

type (
	PinPack [2]machine.Pin
)

type Device_ interface {
	SetI2c(*machine.I2CConfig)
}

type (
	Device struct {
		i2c     *machine.I2C
		scl     machine.Pin
		sda     machine.Pin
		freq    uint32
		init    bool
		addr    uint16
		read    []byte
		payload []byte
		flag    uint8
	}
)

func (dev *Device) SetI2c(cfg machine.I2CConfig) {
	dev.i2c = machine.I2C1
	dev.i2c.Configure(cfg)
	dev.scl = cfg.SCL
	dev.sda = cfg.SDA
	dev.freq = cfg.Frequency
}

func NewPinPack(a, b machine.Pin) *PinPack {
	p_ := PinPack{}
	p_[0] = a
	p_[1] = b
	return &p_
}

func setI2cIfhave(dev *Device, cfg machine.I2CConfig) int {
	err := 0
	if cfg.SCL == 0 {
		err += 0x1
	}
	if cfg.SDA == 0 {
		err += 0x10
	}
	if cfg.SDA == cfg.SCL {
		err += 0x100
	}
	if cfg.Frequency == 0 {
		err += 0x1000
	}

	if err == 0 {
		dev.SetI2c(cfg)
	}
	return err
}
