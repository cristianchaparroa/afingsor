package main

import (
	"github.com/tarm/serial"

	"fmt"
	"github.com/cristianchaparroa/afingsor"
)

func main() {
	c := &serial.Config{Name: "/dev/cu.usbmodem1421", Baud: 9600}
	s := afingsor.Sensor{}
	s.Init(c)

	isVerify := s.VerifyPass()

	fmt.Printf("Found fingerprint seonsro  %s", isVerify)
}
