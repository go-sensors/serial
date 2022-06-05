package serial

import (
	coreio "github.com/go-sensors/core/io"
	coreserial "github.com/go-sensors/core/serial"
	"github.com/pkg/errors"
	"github.com/tarm/serial"
)

type SerialPort struct {
	Config *serial.Config
}

func NewSerialPort(name string, config *coreserial.SerialPortConfig) (*SerialPort, error) {
	var parity serial.Parity
	switch config.Parity {
	case coreserial.ParityNone:
		parity = serial.ParityNone
	case coreserial.ParityOdd:
		parity = serial.ParityOdd
	case coreserial.ParityEven:
		parity = serial.ParityEven
	case coreserial.ParityMark:
		parity = serial.ParityMark
	case coreserial.ParitySpace:
		parity = serial.ParitySpace
	default:
		return nil, errors.Errorf("failed to translate parity %v to an implementation-specific value", config.Parity)
	}

	var stopBits serial.StopBits
	switch config.StopBits {
	case coreserial.Stop1:
		stopBits = serial.Stop1
	case coreserial.Stop1Half:
		stopBits = serial.Stop1Half
	case coreserial.Stop2:
		stopBits = serial.Stop2
	default:
		return nil, errors.Errorf("failed to translate stop bits %v to an implementation-specific value", config.StopBits)
	}

	serialPort := &SerialPort{
		&serial.Config{
			Name:        name,
			Baud:        config.Baud,
			ReadTimeout: config.ReadTimeout,
			Size:        config.Size,
			Parity:      parity,
			StopBits:    stopBits,
		},
	}
	return serialPort, nil
}

func (sp *SerialPort) Open() (coreio.Port, error) {
	port, err := serial.OpenPort(sp.Config)
	if err != nil {
		return nil, errors.Wrap(err, "failed to open port")
	}

	return port, nil
}
