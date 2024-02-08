package actuador

import "github.com/IzcoatlRam/sniffer-go/entorno"
import "fmt"

const (
	MaxPos = 9
	MinPos = 0
)

type ActuadorSimple struct {
	env *entorno.Entorno
}

// NewActuadorSimple crea una nueva instancia de ActuadorSimple.
func NewActuadorSimple(env *entorno.Entorno) *ActuadorSimple {
	return &ActuadorSimple{env: env}
}

// KeyUp implementa la función de movimiento hacia arriba para ActuadorSimple.
func (a *ActuadorSimple) KeyUp() {
	pos := a.env.GetPositionAgent()
	if pos[0] == MinPos {
		fmt.Println("Posición no válida")
	} else {
		pos[0]--
		a.env.MoveAgent(pos)
		fmt.Println(a.env.GetPositionAgent())

	}
}

// KeyDown implementa la función de movimiento hacia abajo para ActuadorSimple.
func (a *ActuadorSimple) KeyDown() {
	pos := a.env.GetPositionAgent()
	if pos[0] == MaxPos {
		fmt.Println("Posición no válida")
	} else {
		pos[0]++
		a.env.MoveAgent(pos)
		fmt.Println(a.env.GetPositionAgent())

	}
}

// KeyLeft implementa la función de movimiento hacia la izquierda para ActuadorSimple.
func (a *ActuadorSimple) KeyLeft() {
	pos := a.env.GetPositionAgent()
	if pos[1] == MinPos {
		fmt.Println("Posición no válida")
	} else {
		pos[1]--
		a.env.MoveAgent(pos)
		fmt.Println(a.env.GetPositionAgent())
	}
}

// KeyRight implementa la función de movimiento hacia la derecha para ActuadorSimple.
func (a *ActuadorSimple) KeyRight() {
	pos := a.env.GetPositionAgent()
	if pos[1] == MaxPos {
		fmt.Println("Posición no válida")
	} else {
		pos[1]++
		a.env.MoveAgent(pos)
		fmt.Println(a.env.GetPositionAgent())

	}
}