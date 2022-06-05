package serial_test

import (
	"fmt"
	"math/rand"
	"testing"
	"time"

	coreserial "github.com/go-sensors/core/serial"
	"github.com/go-sensors/serial"
	"github.com/stretchr/testify/assert"
)

func Test_NewSerialPort_returns_configured_serial_port(t *testing.T) {
	// Arrange
	validParities := []coreserial.Parity{
		coreserial.ParityNone,
		coreserial.ParityOdd,
		coreserial.ParityEven,
		coreserial.ParityMark,
		coreserial.ParitySpace,
	}

	validStopBits := []coreserial.StopBits{
		coreserial.Stop1,
		coreserial.Stop1Half,
		coreserial.Stop2,
	}

	validConfigs := []coreserial.SerialPortConfig{}
	for _, parity := range validParities {
		for _, stopBits := range validStopBits {
			config := coreserial.SerialPortConfig{
				Baud:        rand.Int(),
				ReadTimeout: time.Duration(rand.Int63()),
				Size:        byte(rand.Intn(255)),
				Parity:      parity,
				StopBits:    stopBits,
			}
			validConfigs = append(validConfigs, config)
		}
	}

	for _, config := range validConfigs {
		t.Run(fmt.Sprintf("Valid config %v", config), func(t *testing.T) {
			// Arrange
			expectedName := fmt.Sprintf("some_device_name_or_path-%d", rand.Uint64())

			// Act
			actual, err := serial.NewSerialPort(expectedName, &config)

			// Assert
			assert.NoError(t, err)
			assert.Equal(t, expectedName, actual.Config.Name)
			assert.Equal(t, config.Baud, actual.Config.Baud)
			assert.Equal(t, config.ReadTimeout, actual.Config.ReadTimeout)
			assert.Equal(t, config.Size, actual.Config.Size)
			assert.Equal(t, config.Parity, coreserial.Parity(actual.Config.Parity))
			assert.Equal(t, config.StopBits, coreserial.StopBits(actual.Config.StopBits))
		})
	}
}

func Test_NewSerialPort_fails_with_invalid_parity(t *testing.T) {
	// Arrange
	name := fmt.Sprintf("some_device_name_or_path-%d", rand.Uint64())
	config := coreserial.SerialPortConfig{
		Baud:        rand.Int(),
		ReadTimeout: time.Duration(rand.Int63()),
		Size:        byte(rand.Intn(255)),
		Parity:      coreserial.Parity('X'),
		StopBits:    coreserial.Stop1,
	}

	// Act
	actual, err := serial.NewSerialPort(name, &config)

	// Assert
	assert.ErrorContains(t, err, "failed to translate parity")
	assert.Nil(t, actual)
}

func Test_NewSerialPort_fails_with_invalid_stop_bits(t *testing.T) {
	// Arrange
	name := fmt.Sprintf("some_device_name_or_path-%d", rand.Uint64())
	config := coreserial.SerialPortConfig{
		Baud:        rand.Int(),
		ReadTimeout: time.Duration(rand.Int63()),
		Size:        byte(rand.Intn(255)),
		Parity:      coreserial.ParityNone,
		StopBits:    coreserial.StopBits(3),
	}

	// Act
	actual, err := serial.NewSerialPort(name, &config)

	// Assert
	assert.ErrorContains(t, err, "failed to translate stop bits")
	assert.Nil(t, actual)
}
