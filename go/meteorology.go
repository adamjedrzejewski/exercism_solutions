package meteorology

import "fmt"

type TemperatureUnit int

const (
	Celsius    TemperatureUnit = 0
	Fahrenheit TemperatureUnit = 1
)

func (t TemperatureUnit) String() string {
	switch t {
	case Celsius:
		return "°C"
	case Fahrenheit:
		return "°F"
	default:
		panic("No such temperature unit")
	}
}

type Temperature struct {
	degree int
	unit   TemperatureUnit
}

func (t Temperature) String() string {
	return fmt.Sprintf("%d %s",
		t.degree,
		t.unit.String())
}

type SpeedUnit int

const (
	KmPerHour    SpeedUnit = 0
	MilesPerHour SpeedUnit = 1
)

func (s SpeedUnit) String() string {
	switch s {
	case KmPerHour:
		return "km/h"
	case MilesPerHour:
		return "mph"
	default:
		panic("No such speed unit")
	}
}

type Speed struct {
	magnitude int
	unit      SpeedUnit
}

func (s Speed) String() string {
	return fmt.Sprintf("%d %s",
		s.magnitude,
		s.unit.String())
}

type MeteorologyData struct {
	location      string
	temperature   Temperature
	windDirection string
	windSpeed     Speed
	humidity      int
}

func (m MeteorologyData) String() string {
	return fmt.Sprintf("%s: %s, Wind %s at %s, %d%% Humidity",
		m.location,
		m.temperature,
		m.windDirection,
		m.windSpeed,
		m.humidity)
}

// Add a String method to MeteorologyData
