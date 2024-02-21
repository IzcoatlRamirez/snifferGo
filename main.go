package main

import (
	"fmt"
	"strconv"

	actuador "github.com/IzcoatlRam/sniffer-go/actuador"
	entorno "github.com/IzcoatlRam/sniffer-go/entorno"
	sensor "github.com/IzcoatlRam/sniffer-go/sensor"
	sniffer "github.com/IzcoatlRam/sniffer-go/sniffer"
	"github.com/fatih/color"
)

const simulaciones = 10

func main() {
	var PuntuajeGlobal = 0
	var Puntuajes []int
	for i := 0; i < simulaciones; i++ {
		e1 := entorno.NewEntorno(10, 0.50)
		s1 := sensor.NewSensorSimple(e1)
		a1 := actuador.NewActuadorSimple(e1)
		sniffer1 := sniffer.NewSniffer(s1, a1, e1)
		sniffer1.Run()
		Puntuajes = append(Puntuajes, sniffer1.GetPuntuaje())
		PuntuajeGlobal = PuntuajeGlobal + sniffer1.Actuador.GetCont()
	}
	for i := 0; i < simulaciones; i++ {
		fmt.Printf("Actuador #%s Puntuaje: %s\n", color.HiBlueString(strconv.Itoa(i+1)), color.GreenString(strconv.Itoa(Puntuajes[i])))
	}
	fmt.Printf("\nPuntuaje GLobal: %s\n", color.HiYellowString(strconv.Itoa(PuntuajeGlobal)))
	fmt.Printf("Puntuaje Promedio: %s", color.GreenString(strconv.Itoa(PuntuajeGlobal/simulaciones)))
}
