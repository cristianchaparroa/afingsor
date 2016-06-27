package afingsor

import (
	"github.com/tarm/serial"
)

// ISensor set the beahivor to comunicate with the fingerprint sensor
type ISensor interface {
	CreateModel() uint8
	DeleteModel(id uint16) uint8
	EmptyDatabase() uint8
	GetImage() uint8
	GetModel() uint8
	GetReply(packet []uint8, timeout uint16)
	Image2tz() uint8
	LoadModel(id uint16) uint8
	StoreModel(id uint16) uint8
	VerifyPass() bool
	Setup(*serial.Config)
	writePacket(addr uint32, packettype uint8, len uint16, packet *uint8)
}

// Sensor provides the beahivor to comunicate with the fingerprint sensor
type Sensor struct {
	Password uint32
	Address  uint32
	serial   *serial.Port
	config   *serial.Config
}

// CreateModel ...
func (s *Sensor) CreateModel() uint8 {
	return 0
}

// DeleteModel ...
func (s *Sensor) DeleteModel(id uint16) uint8 {
	return 0
}

// EmptyDatabase ...
func (s *Sensor) EmptyDatabase() uint8 {
	return 0
}

// GetImage ...
func (s *Sensor) GetImage() uint8 {
	return 0
}

// GetModel transfer a fingerprint template from char buffer 1 to host computer
func (s *Sensor) GetModel() uint8 {
	return 0
}

// GetReply ...
func (s *Sensor) GetReply(packet []uint8, timeout uint16) {}

// Image2tz ...
func (s *Sensor) Image2tz() uint8 {
	return 0
}

// LoadModel ...
func (s *Sensor) LoadModel(id uint16) uint8 {
	return 0
}

// StoreModel ...
func (s *Sensor) StoreModel(id uint16) uint8 {
	return 0
}

// VerifyPass ..
func (s *Sensor) VerifyPass() bool {
	return false
}

// Setup ...
func (s *Sensor) Setup(*serial.Config) {}

//
func (s *Sensor) writePacket(addr uint32, packettype uint8, len uint16, packet *uint8) {

}
