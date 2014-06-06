package gpio

import (
	"github.com/hybridgroup/gobot"
)

type AnalogSensorDriver struct {
	gobot.Driver
	Adaptor AnalogReader
}

func NewAnalogSensorDriver(a AnalogReader, name string, pin string) *AnalogSensorDriver {
	return &AnalogSensorDriver{
		Driver: gobot.Driver{
			Name: name,
			Pin:  pin,
			Commands: []string{
				"ReadC",
			},
		},
		Adaptor: a,
	}
}

func (a *AnalogSensorDriver) Start() bool { return true }
func (a *AnalogSensorDriver) Init() bool  { return true }
func (a *AnalogSensorDriver) Halt() bool  { return true }

func (a *AnalogSensorDriver) Read() int {
	return a.Adaptor.AnalogRead(a.Pin)
}
