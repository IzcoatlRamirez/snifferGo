package actuador

import "github.com/IzcoatlRam/sniffer-go/entorno"
import "fmt"
import "container/list"
import "math/rand"

type ActuadorSimple struct {
	env *entorno.Entorno
	MaxPos int
	MinPos int
	LimMoves int
	Movements int
	Memory   *list.List
	Jumps	int
	JumpSize int
}

// NewActuadorSimple crea una nueva instancia de ActuadorSimple.
func NewActuadorSimple(env *entorno.Entorno) *ActuadorSimple {
	return &ActuadorSimple{
		env: env,
		MaxPos: env.Dimension,
		MinPos: 0,
		LimMoves: ((env.Dimension * env.Dimension)/3)*2,
		Movements: 0,
		Jumps: 0,
		Memory:   list.New(),
		JumpSize: env.Dimension,
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
    a.env.MoveAgent(pos)
    a.Movements++
}

func (a *ActuadorSimple) isCellVisited(pos [2]int) bool {
    for e := a.Memory.Front(); e != nil; e = e.Next() {
        if e.Value.([2]int) == pos {
            return true
        }
    }
    return false
}

func (a *ActuadorSimple) GetAvailableMoves() [][2]int{
	currentPos := a.env.GetPositionAgent()
	moves := make([][2]int, 0)
	
	//ezquina superior izquierda
	if currentPos[0]==0 && currentPos[1] == 0 {
		moves = append(moves, [2]int{currentPos[0]+1, currentPos[1]})
		moves = append(moves, [2]int{currentPos[0], currentPos[1]+1})
		fmt.Println("Esquina superior izquierda")
		a.cleanMemory()
	}

	
	//ezquina superior derecha
	if currentPos[0]==0 && currentPos[1] == a.env.Dimension-1 {
		moves = append(moves, [2]int{currentPos[0]+1, currentPos[1]})
		moves = append(moves, [2]int{currentPos[0], currentPos[1]-1})
		fmt.Println("Esquina superior derecha")
		a.cleanMemory()
	}


	//ezquina inferior izquierda
	if currentPos[0]==a.env.Dimension-1 && currentPos[1] == 0 {
		moves = append(moves, [2]int{currentPos[0]-1, currentPos[1]})
		moves = append(moves, [2]int{currentPos[0], currentPos[1]+1})
		fmt.Println("Esquina inferior izquierda")
		a.cleanMemory()
	}

	//ezquina inferior derecha
	if currentPos[0]==a.env.Dimension-1 && currentPos[1] == a.env.Dimension-1 {
		moves = append(moves, [2]int{currentPos[0]-1, currentPos[1]})
		moves = append(moves, [2]int{currentPos[0], currentPos[1]-1})
		fmt.Println("Esquina inferior derecha")
		a.cleanMemory()
	}

	//borde superior
	if currentPos [0] == 0 && currentPos [1] != 0 && currentPos [1] != a.env.Dimension-1{
		moves = append(moves, [2]int{currentPos[0]+1, currentPos[1]})
		moves = append(moves, [2]int{currentPos[0], currentPos[1]+1})
		moves = append(moves, [2]int{currentPos[0], currentPos[1]-1})
		fmt.Println("Borde superior")
		a.cleanMemory()
	}

	//borde inferior
	if currentPos [0] == a.env.Dimension-1 && currentPos [1] != 0 && currentPos [1] != a.env.Dimension-1{
		moves = append(moves, [2]int{currentPos[0]-1, currentPos[1]})
		moves = append(moves, [2]int{currentPos[0], currentPos[1]+1})
		moves = append(moves, [2]int{currentPos[0], currentPos[1]-1})
		fmt.Println("Borde inferior")
		a.cleanMemory()
	}

	//borde izquierdo
	if currentPos [1] == 0 && currentPos [0] != 0 && currentPos [0] != a.env.Dimension-1{
		moves = append(moves, [2]int{currentPos[0]+1, currentPos[1]})
		moves = append(moves, [2]int{currentPos[0]-1, currentPos[1]})
		moves = append(moves, [2]int{currentPos[0], currentPos[1]+1})
		fmt.Println("Borde izquierdo")
		a.cleanMemory()
	}

	//borde derecho
	if currentPos [1] == a.env.Dimension-1 && currentPos [0] != 0 && currentPos [0] != a.env.Dimension-1{
		moves = append(moves, [2]int{currentPos[0]+1, currentPos[1]})
		moves = append(moves, [2]int{currentPos[0]-1, currentPos[1]})
		moves = append(moves, [2]int{currentPos[0], currentPos[1]-1})
		fmt.Println("Borde derecho")
		a.cleanMemory()
	}

	//centro
	if currentPos [0] != 0 && currentPos [0] != a.env.Dimension-1 && currentPos [1] != 0 && currentPos [1] != a.env.Dimension-1{
		moves = append(moves, [2]int{currentPos[0]+1, currentPos[1]})
		moves = append(moves, [2]int{currentPos[0]-1, currentPos[1]})
		moves = append(moves, [2]int{currentPos[0], currentPos[1]+1})
		moves = append(moves, [2]int{currentPos[0], currentPos[1]-1})
		fmt.Println("Centro")
	}

	return moves
}

func (a *ActuadorSimple) cleanMemory() {
	a.Memory.Init()
}

// Función para generar una posición aleatoria con distancia específica desde la posición actual.
func generateRandomPosition(currentPos [2]int, distance int, env *entorno.Entorno) [2]int {
	// Generar posiciones aleatorias hasta encontrar una con la distancia deseada
	distance /= 2
	var newPos [2]int
	for {
		// Calcular diferencias en cada dimensión
		deltaX := rand.Intn(2*distance+1) - distance
		deltaY := rand.Intn(2*distance+1) - distance

		// Calcular la nueva posición
		newPos[0] = currentPos[0] + deltaX
		newPos[1] = currentPos[1] + deltaY

		// Verificar que la nueva posición esté dentro de los límites del entorno
		if newPos[0] >= 0 && newPos[0] < env.Dimension && newPos[1] >= 0 && newPos[1] < env.Dimension {
			break
		}
	}
	return newPos
}

func (a *ActuadorSimple) JumpAgent() {
	// Obtener la posición actual del agente
	currentPos := a.env.GetPositionAgent()

	// Generar una nueva posición aleatoria 
	newPos := generateRandomPosition(currentPos, a.JumpSize, a.env)

	// Mover al agente a la nueva posición
	a.env.MoveAgent(newPos)

	// Almacenar la nueva posición en la memoria
	if !a.isCellVisited(newPos) {
		a.Memory.PushBack(newPos)
	}
	// Verificar y eliminar elementos de la memoria si está llena
	if a.Memory.Len() > (a.env.Dimension*a.env.Dimension)/5 {
		elementsToRemove := a.Memory.Len() / 1
		for i := 0; i < elementsToRemove; i++ {
			a.Memory.Remove(a.Memory.Front())
		}
	}
	a.Movements++
}