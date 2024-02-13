package actuador

import "github.com/IzcoatlRam/sniffer-go/entorno"
import "fmt"
import "container/list"

type ActuadorSimple struct {
	env *entorno.Entorno
	MaxPos int
	MinPos int
	LimMoves int
	Movements int
	Memory   *list.List
}

// NewActuadorSimple crea una nueva instancia de ActuadorSimple.
func NewActuadorSimple(env *entorno.Entorno) *ActuadorSimple {
	return &ActuadorSimple{
		env: env,
		MaxPos: env.Dimension,
		MinPos: 0,
		LimMoves: env.Dimension * env.Dimension,
		Movements: 0,
		Memory:   list.New(),
		}
}

// KeyUp implementa la función de movimiento hacia arriba para ActuadorSimple.
func (a *ActuadorSimple) KeyUp() {
	pos := a.env.GetPositionAgent()
	if pos[0] == a.MinPos {
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
	if pos[0] == a.MaxPos {
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
	if pos[1] == a.MinPos {
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
	if pos[1] == a.MaxPos {
		fmt.Println("Posición no válida")
	} else {
		pos[1]++
		a.env.MoveAgent(pos)
		fmt.Println(a.env.GetPositionAgent())
	}
}

func (a *ActuadorSimple) MoveAgent(pos [2]int) {
	a.Memory.Init()
	a.Memory.PushBack(a.env.GetPositionAgent())
	a.env.MoveAgent(pos)
	a.Movements++
}