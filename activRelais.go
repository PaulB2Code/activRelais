package Relay

import (
	"fmt"
	"time"

	"github.com/stianeikeland/go-rpio"
)

const ()

func (r *Relay) Close() {
	// Unmap gpio memory when done
	rpio.Close()
}

type Relay struct {
	Status bool
	On     rpio.Pin
	Off    rpio.Pin
}

var relaisOn_Number = 13
var relaisOff_Number = 19

var timeDelay = 100 * time.Millisecond

func New() (Relay, error) {
	// Open and map memory to access gpio, check for errors
	if err := rpio.Open(); err != nil {
		fmt.Println(err)
		return Relay{}, err
	}
	relaisPinOn := rpio.Pin(relaisOn_Number)
	relaisPinOff := rpio.Pin(relaisOff_Number)
	relaisPinOn.Output()
	relaisPinOff.Output()
	return Relay{On: relaisPinOn, Off: relaisPinOff}, nil
}

func (r *Relay) POff() {
	r.Off.Low()
	r.On.Low()
	time.Sleep(timeDelay)
	r.On.High()
}

func (r *Relay) POn() {
	r.On.Low()
	r.Off.Low()
	time.Sleep(timeDelay)
	r.Off.High()
}
