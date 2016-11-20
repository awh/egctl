package energenie

import (
	"time"

	log "github.com/Sirupsen/logrus"
	"github.com/stianeikeland/go-rpio"
)

var onCodes = map[string][4]rpio.State{
	"all":   [4]rpio.State{rpio.High, rpio.Low, rpio.High, rpio.High},
	"one":   [4]rpio.State{rpio.High, rpio.High, rpio.High, rpio.High},
	"two":   [4]rpio.State{rpio.High, rpio.High, rpio.High, rpio.Low},
	"three": [4]rpio.State{rpio.High, rpio.High, rpio.Low, rpio.High},
	"four":  [4]rpio.State{rpio.High, rpio.High, rpio.Low, rpio.Low},
}

var offCodes = map[string][4]rpio.State{
	"all":   [4]rpio.State{rpio.Low, rpio.Low, rpio.High, rpio.High},
	"one":   [4]rpio.State{rpio.Low, rpio.High, rpio.High, rpio.High},
	"two":   [4]rpio.State{rpio.Low, rpio.High, rpio.High, rpio.Low},
	"three": [4]rpio.State{rpio.Low, rpio.High, rpio.Low, rpio.High},
	"four":  [4]rpio.State{rpio.Low, rpio.High, rpio.Low, rpio.Low},
}

func execute(codes map[string][4]rpio.State, sockets ...string) error {
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

	for _, socket := range sockets {
		if code, ok := codes[socket]; ok {
			// Configure modulator
			d3.Write(code[0])
			d2.Write(code[1])
			d1.Write(code[2])
			d0.Write(code[3])

			// Wait for encoder to settle
			time.Sleep(100 * time.Millisecond)

			// Enable transmitter for a short period
			modulator.High()
			time.Sleep(250 * time.Millisecond)
			modulator.Low()
		} else {
			log.Warnf("Skipping bad socket name: %s", socket)
		}
	}

	return nil
}

func On(sockets ...string) {
	execute(onCodes, sockets...)
}

func Off(sockets ...string) {
	execute(offCodes, sockets...)
}
