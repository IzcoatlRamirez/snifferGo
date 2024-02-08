package sniffer

import (
	"github.com/IzcoatlRam/sniffer-go/actuador"
	"github.com/IzcoatlRam/sniffer-go/entorno"
	"github.com/IzcoatlRam/sniffer-go/sensor"
)

type Sniffer struct {
	Env      *entorno.Entorno
	Sensor   *sensor.SensorSimple
	Actuador *actuador.ActuadorSimple
}

// NewSniffer crea una nueva instancia de Sniffer.
func NewSniffer(sensor *sensor.SensorSimple, actuador *actuador.ActuadorSimple, env *entorno.Entorno) *Sniffer {
	return &Sniffer{
		Sensor:   sensor,
		Actuador: actuador,
		Env:      env,
	}
}
