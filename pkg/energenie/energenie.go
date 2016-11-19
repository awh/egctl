package energenie

import (
	"time"

	"github.com/stianeikeland/go-rpio"
)

func Execute(code [4]rpio.State) error {
	err := rpio.Open()
	if err != nil {
		return err
	}

	// Instantiate GPIO pins
	d0 := rpio.Pin(17)
	d1 := rpio.Pin(22)
	d2 := rpio.Pin(23)
	d3 := rpio.Pin(27)
	keying := rpio.Pin(24)
	modulator := rpio.Pin(25)

	// Configure for output
	d0.Output()
	d1.Output()
	d2.Output()
	d3.Output()
	keying.Output()
	modulator.Output()

	// Disable the modulator
	modulator.Low()
	// Set modulator to ASK
	keying.Low()

	// Configure modulator
	d3.Write(code[0])
	d2.Write(code[1])
	d1.Write(code[2])
	d0.Write(code[3])

	// Wait for encoder to settle
	time.Sleep(100 * time.Millisecond)

	// Enable transmitter
	modulator.High()
	time.Sleep(250 * time.Millisecond)
	modulator.Low()

	return nil
}
