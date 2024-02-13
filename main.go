package main

import (
	// "fmt"
	actuador "github.com/IzcoatlRam/sniffer-go/actuador"
	entorno "github.com/IzcoatlRam/sniffer-go/entorno"
	sensor "github.com/IzcoatlRam/sniffer-go/sensor"
	sniffer "github.com/IzcoatlRam/sniffer-go/sniffer"
)

func main() {
    e1 := entorno.NewEntorno(10, 0.50)
    s1 := sensor.NewSensorSimple(e1)
    a1 := actuador.NewActuadorSimple(e1)
    sniffer1 := sniffer.NewSniffer(s1, a1, e1)
    sniffer1.Run()

    // fmt.Println()

    // e2 := entorno.NewEntorno(5, 0.5)
    // s2 := sensor.NewSensorSimple(e2)
    // a2 := actuador.NewActuadorSimple(e2)
    // sniffer2 := sniffer.NewSniffer(s2, a2, e2)
    // sniffer2.Run()
}
