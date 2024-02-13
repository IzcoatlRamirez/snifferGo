package sensor

import "github.com/IzcoatlRam/sniffer-go/entorno"


// SensorSimple es una implementación concreta de la interfaz Sensor.
type SensorSimple struct {
	env *entorno.Entorno
}

// NewSensorSimple crea una nueva instancia de SensorSimple.
func NewSensorSimple(env *entorno.Entorno) *SensorSimple {
	return &SensorSimple{env: env}
}

// DetectDirt implementa la función de detección de suciedad para SensorSimple.
func (s *SensorSimple) DetectDirt(position [2]int) bool {
	return s.env.IsDirty(position)
}

func (s *SensorSimple) GetMoves() [][2]int {
	return s.env.GetAvailableMoves()
}	