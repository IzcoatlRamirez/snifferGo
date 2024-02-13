package sniffer

import (
	"container/list"
	"fmt"
	"github.com/IzcoatlRam/sniffer-go/actuador"
	"github.com/IzcoatlRam/sniffer-go/entorno"
	"github.com/IzcoatlRam/sniffer-go/sensor"
	"github.com/eiannone/keyboard"
	"math/rand"
	"os"
	"os/exec"
	"time"
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

func (s *Sniffer) Start() {
	Exit := false
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
		s.Env.PrintMatrix()
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
			s.Actuador.KeyUp()
		case keyboard.KeyArrowDown:
			fmt.Println("Abajo")
			s.Actuador.KeyDown()
		case keyboard.KeyArrowRight:
			fmt.Println("Derecha")
			s.Actuador.KeyRight()
		case keyboard.KeyArrowLeft:
			fmt.Println("Izquierda")
			s.Actuador.KeyLeft()
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

func (s *Sniffer) Run() {
	rand.Seed(time.Now().UnixNano())
	Exit := false

	for !Exit {
		fmt.Print("Movimientos: ", s.Actuador.Movements, "/", s.Actuador.LimMoves,"\n")
		s.Env.PrintMatrix()
		time.Sleep(1 * time.Second)
		clearScreen()

		currentPos := s.Env.GetPositionAgent()
		if s.Sensor.DetectDirt(currentPos) {
			s.Env.CleanCell(currentPos)
		}

		availableMoves := s.Sensor.GetMoves()
		availableMoves = filterMoves(availableMoves, s.Actuador.Memory)
		if len(availableMoves) > 0 {
			moveFound := false
			for _, move := range availableMoves {
				if s.Sensor.DetectDirt(move) {
					s.Actuador.MoveAgent(move)
					s.Env.CleanCell(move)
					fmt.Println("Moviendo a ", move)
					moveFound = true
					break
				}
			}

			if !moveFound {
				// No se encontró suciedad en las posiciones disponibles, moverse aleatoriamente
				index := rand.Intn(len(availableMoves))
				destino := availableMoves[index]
				s.Actuador.MoveAgent(destino)
				fmt.Println("Moviendo a ", destino)
			}
		}

		// Condición de paro
		if s.Actuador.Movements == s.Actuador.LimMoves {
			Exit = true
			s.Env.PrintMatrix()
		}
	}
}

// Función para filtrar las posiciones disponibles excluyendo la posición de la memoria
func filterMoves(availableMoves [][2]int, memory *list.List) [][2]int {
	filteredMoves := make([][2]int, 0)

	for _, move := range availableMoves {
		if !positionInMemory(move, memory) {
			filteredMoves = append(filteredMoves, move)
		}
	}

	return filteredMoves
}

// Función para verificar si una posición está en la memoria
func positionInMemory(pos [2]int, memory *list.List) bool {
	for e := memory.Front(); e != nil; e = e.Next() {
		if e.Value.([2]int) == pos {
			return true
		}
	}
	return false
}

