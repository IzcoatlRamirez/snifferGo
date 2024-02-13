package main

import (
	// "fmt"
	"github.com/IzcoatlRam/sniffer-go/actuador"
	"github.com/IzcoatlRam/sniffer-go/entorno"
	"github.com/IzcoatlRam/sniffer-go/sensor"
	"github.com/IzcoatlRam/sniffer-go/sniffer"
)

func main() {

	env := entorno.NewEntorno(10, 0.5)

	sensor := sensor.NewSensorSimple(env)
	actuador := actuador.NewActuadorSimple(env)

	sniffer := sniffer.NewSniffer(sensor, actuador, env)

	// sniffer.Env.PrintMatrix()
	// fmt.Println(sniffer.Sensor.DetectDirt([2]int{0, 1}))
	sniffer.Run()

	//
}
