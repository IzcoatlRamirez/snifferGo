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
	"github.com/fatih/color"
	"strconv"
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
		ClearScreen()
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

func ClearScreen() {
	cmd := exec.Command("cmd", "/c", "cls")
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func (s *Sniffer) Run() {
	rand.Seed(time.Now().UnixNano())
	Exit := false

	for !Exit {
		s.Show()
		time.Sleep(150 * time.Millisecond)
		ClearScreen()

		currentPos := s.Env.GetPositionAgent()
		if s.Sensor.DetectDirt(currentPos) {
			s.Env.CleanCell(currentPos)
		}

		availableMoves := s.Actuador.GetAvailableMoves()
		availableMoves = filterMoves(availableMoves, s.Actuador.Memory)
		if len(availableMoves) > 0 {
			moveFound := false
			for _, move := range availableMoves {
				if s.Sensor.DetectDirt(move) {
					s.Actuador.MoveAgent(move)
					s.Env.CleanCell(move)
					// fmt.Println("Moviendo a ", move)
					moveFound = true
					break
				}
			}
			if !moveFound {
				s.Actuador.JumpAgent()
			}
		}
		// Condición de paro
		if s.Actuador.Movements == s.Actuador.LimMoves {
			Exit = true
			s.Show()
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

func (s *Sniffer) Show() {
	fmt.Printf("Movimientos: %s/%s\n", color.CyanString(strconv.Itoa(s.Actuador.Movements)), color.CyanString(strconv.Itoa(s.Actuador.LimMoves)))
	fmt.Printf("Celdas en memoria: %s\n", color.YellowString(strconv.Itoa(s.Actuador.Memory.Len())))
	fmt.Printf("Basura total: %s\n",color.MagentaString(strconv.Itoa(s.Env.DirtCount)))
	fmt.Printf("Basura limpiada: %s\n", color.GreenString(strconv.Itoa(s.Env.CleanCellCount)))
	fmt.Printf("Porcentaje de limpieza: %s%%\n\n", color.RedString("%.2f", (float64(s.Env.CleanCellCount)/float64(s.Env.DirtCount))*100))	
	s.Env.PrintMatrix()
}
