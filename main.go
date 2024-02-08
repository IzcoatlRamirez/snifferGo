package main

import (
	"fmt"
	"os"
	"os/exec"
	"github.com/IzcoatlRam/sniffer-go/sniffer"
	"github.com/IzcoatlRam/sniffer-go/actuador"
	"github.com/IzcoatlRam/sniffer-go/entorno"
	"github.com/IzcoatlRam/sniffer-go/sensor"
	"github.com/eiannone/keyboard"
)

func main() {
	Exit := false

	env := entorno.NewEntorno(10, 0.1)

	sensor := sensor.NewSensorSimple(env)
	actuador := actuador.NewActuadorSimple(env)

	sniffer := sniffer.NewSniffer(sensor, actuador, env)

	err := keyboard.Open()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer func() {
		_ = keyboard.Close()
	}()

	for !Exit {
		fmt.Println()
		sniffer.Env.PrintMatrix()
		fmt.Println()
		
		fmt.Print("Presiona las flechas (↑, ↓, ←, →) para desplazarte \n(presione 'esc' para salir)\n")
		_, key, err := keyboard.GetKey()
		clearScreen()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		switch key {
		case keyboard.KeyArrowUp:
			fmt.Println("Arriba")
			sniffer.Actuador.KeyUp()
		case keyboard.KeyArrowDown:
			fmt.Println("Abajo")
			sniffer.Actuador.KeyDown()
		case keyboard.KeyArrowRight:
			fmt.Println("Derecha")
			sniffer.Actuador.KeyRight()
		case keyboard.KeyArrowLeft:
			fmt.Println("Izquierda")
			sniffer.Actuador.KeyLeft()
		case keyboard.KeyEsc, keyboard.KeyCtrlC, 'q':
			fmt.Println("Saliendo...")
			return
		default:
			fmt.Println("Dirección no reconocida")
		}
	}

}

func clearScreen() {
	cmd := exec.Command("cmd", "/c", "cls")
	cmd.Stdout = os.Stdout
	cmd.Run()
}
