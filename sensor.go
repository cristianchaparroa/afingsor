package afingsor

import (
	"unsafe"

	"fmt"
	"github.com/tarm/serial"
	"log"
)

// Documentation
// http://www.fut-electronics.com/wp-content/plugins/fe_downloads/Uploads/Advanced_fingerprint_module_for_arduino.pdf

// ISensor set the beahivor to comunicate with the fingerprint sensor
type ISensor interface {
	CreateModel() uint8
	DeleteModel(id uint16) uint8
	EmptyDatabase() uint8
	GetImage() uint8
	GetModel() uint8
	GetReply(packet []uint8, timeout uint16) uint8
	Image2tz() uint8
	LoadModel(id uint16) uint8
	StoreModel(id uint16) bool
	Setup(*serial.Config)
	VerifyPass() bool
	//ptr to packet should be returned
	writePacket(addr uint32, packettype uint8, len uint16, packet *[]uint8)
}

// Sensor provides the behavior to communicate pc with the fingerprint sensor
// through serial port
type Sensor struct {
	Password uint32
	Address  uint32
	port     *serial.Port
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

// GetReply handle the response from the sensor
func (s *Sensor) GetReply(packet []uint8, timeout uint16) uint8 {
	//var reply [20]int
	//idx := 0

	for {
		r := s.Read()
		fmt.Printf("what read?:  %q", r)
		//reply[idx] = s.Read()
		//idx++
		//if idx >= 9 {
		//	if (reply[0] != (FingerPrintStartCode >> 8)) || (reply[1] != FingerPrintStartCode & 0xFF) {
		//		return FingerPrintBadPacket
		//	}
		//	packetType := reply[6]
		//	len := reply[7]
		//	len = len << 8
		//	len |= reply[8]
		//	len -= 2
		//
		//	packet[0] = packetType
		//	for i := int(0); i < len; i++ {
		//		packet[1 + i] = reply[9 + 1]
		//	}
		//	return len
		//}
	}

}

// Image2tz ...
func (s *Sensor) Image2tz() int {
	return 0
}

// LoadModel ...
func (s *Sensor) LoadModel(id uint16) uint8 {
	return 0
}

// StoreModel ...
func (s *Sensor) StoreModel(id uint16) bool {
	packet := []uint8{FingerPrintStore, 0x01, uint8(id >> 8), uint8(id & 0xFF)}
	s.writePacket(s.Address, FingerPrintCommandPacket, uint16(unsafe.Sizeof(packet)+2), &packet)
	len := s.GetReply(packet, 0)
	if len != 1 && packet[0] != FingerPrintACKPacket {
		return true
	}
	return false
}

// VerifyPass ..
func (s *Sensor) VerifyPass() bool {
	fmt.Println("Veryfing pass.....")
	packet := []uint8{FingerPrintVerifyPassword, uint8(s.Password >> 24),
		uint8(s.Password >> 16), uint8(s.Password >> 8), uint8(s.Password)}

	s.writePacket(s.Address, FingerPrintCommandPacket, 7, &packet)
	len := s.GetReply(packet, 0)

	if len == 1 && (packet[0] == FingerPrintACKPacket) && (packet[1] == FingerPrintOk) {
		return true
	}
	return false
}

//Setup allows to configure the sensor conexion
func (s *Sensor) Init(c *serial.Config) {
	s.Address = 0xFFFFFFFF
	var err error
	s.port, err = serial.OpenPort(c)
	if err != nil {
		panic(err)
	}
}

// writePacket
func (s *Sensor) writePacket(addr uint32, packettype uint8, len uint16, packet *[]uint8) {
	s.Write(FingerPrintStartCode >> 8)
	s.Write(FingerPrintStartCode)
	s.Write(s.Address >> 24)
	s.Write(s.Address >> 16)
	s.Write(s.Address >> 8)
	s.Write(s.Address)
	s.Write(packettype)
	s.Write(len >> 8)
	s.Write(len)

	sum := (len >> 8) + (len & 0xFF) + uint16(packettype)
	for i := uint16(0); i < len-2; i++ {
		s.Write((*packet)[i])
		val := (*packet)[i]
		sum += uint16(val)
	}
	s.Write(sum >> 8)
	s.Write(sum)
}

// Write is a wrapper to write a message via serial port
func (s *Sensor) Write(message interface{}) {
	b, err := GetBytes(message)
	if err != nil {
		log.Fatal(err)
	}
	_, err = s.port.Write(b)
	if err != nil {
		log.Fatal(err)
	}

}

// Read is a  wrapper function to read the content in the serial port
func (s *Sensor) Read() []byte {
	buffer := make([]byte, 2048)
	n, err := s.port.Read(buffer)
	if err != nil {
		return nil
	}
	return buffer[:n]
}
