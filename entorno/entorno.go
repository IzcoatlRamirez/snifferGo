package entorno
import (
	"fmt"
	"math/rand"
	"github.com/fatih/color"
)
// Entorno representa el entorno del sniffer.
type Entorno struct {
	Dimension       int
	probabilityDirt float64
	positionAgent   [2]int
	matrix          [][]bool
}

// NewEntorno crea una nueva instancia de Entorno.
func NewEntorno(Dimension int, probabilityDirt float64) *Entorno {
	e := &Entorno{
 		Dimension:       Dimension,
		probabilityDirt: probabilityDirt,
	}
	e.positionAgent = [2]int{0, 0}
	e.generateDirt()
	return e
}

// PrintMatrix imprime la matriz del entorno en la consola.
func (e *Entorno) PrintMatrix() {
	for i := 0; i < e.Dimension; i++ {
		for j := 0; j < e.Dimension; j++ {
			if i == e.positionAgent[0] && j == e.positionAgent[1] {
				fmt.Printf(" %s ", color.BlueString("S"))
			} else if e.matrix[i][j] {
				fmt.Print(" X ")
			} else {
				fmt.Print(" O ")
			}
		}
		fmt.Println()
	}
}

// generateDirt genera suciedad en el entorno según la probabilidad.
func (e *Entorno) generateDirt() {
	for i := 0; i < e.Dimension; i++ {
		var row []bool
		for j := 0; j < e.Dimension; j++ {
			random := rand.Float64()
			if random < e.probabilityDirt {
				row = append(row, true)
			} else {
				row = append(row, false)
			}
		}
		e.matrix = append(e.matrix, row)
	}
}

// CleanCell limpia una celda específica en el entorno.
func (e *Entorno) CleanCell(position [2]int) {
	e.matrix[position[0]][position[1]] = false
}

// IsDirty verifica si una posición específica en el entorno está sucia.
func (e *Entorno) IsDirty(position [2]int) bool {
	return e.matrix[position[0]][position[1]]
}

// GetPositionAgent devuelve la posición actual del agente en el entorno.
func (e *Entorno) GetPositionAgent() [2]int {
	return e.positionAgent
}

// MoveAgent mueve al agente a una nueva posición en el entorno.
func (e *Entorno) MoveAgent(position [2]int) {
	e.positionAgent = position
}

func (e *Entorno) GetAvailableMoves() [][2]int{
	currentPos := e.GetPositionAgent()
	moves := make([][2]int, 0)
	
	//ezquina superior izquierda
	if currentPos[0]==0 && currentPos[1] == 0 {
		moves = append(moves, [2]int{currentPos[0]+1, currentPos[1]})
		moves = append(moves, [2]int{currentPos[0], currentPos[1]+1})
		fmt.Println("Esquina superior izquierda")
	}

	
	//ezquina superior derecha
	if currentPos[0]==0 && currentPos[1] == e.Dimension-1 {
		moves = append(moves, [2]int{currentPos[0]+1, currentPos[1]})
		moves = append(moves, [2]int{currentPos[0], currentPos[1]-1})
		fmt.Println("Esquina superior derecha")
	}


	//ezquina inferior izquierda
	if currentPos[0]==e.Dimension-1 && currentPos[1] == 0 {
		moves = append(moves, [2]int{currentPos[0]-1, currentPos[1]})
		moves = append(moves, [2]int{currentPos[0], currentPos[1]+1})
		fmt.Println("Esquina inferior izquierda")
	}

	//ezquina inferior derecha
	if currentPos[0]==e.Dimension-1 && currentPos[1] == e.Dimension-1 {
		moves = append(moves, [2]int{currentPos[0]-1, currentPos[1]})
		moves = append(moves, [2]int{currentPos[0], currentPos[1]-1})
		fmt.Println("Esquina inferior derecha")
	}

	//borde superior
	if currentPos [0] == 0 && currentPos [1] != 0 && currentPos [1] != e.Dimension-1{
		moves = append(moves, [2]int{currentPos[0]+1, currentPos[1]})
		moves = append(moves, [2]int{currentPos[0], currentPos[1]+1})
		moves = append(moves, [2]int{currentPos[0], currentPos[1]-1})
		fmt.Println("Borde superior")
	}

	//borde inferior
	if currentPos [0] == e.Dimension-1 && currentPos [1] != 0 && currentPos [1] != e.Dimension-1{
		moves = append(moves, [2]int{currentPos[0]-1, currentPos[1]})
		moves = append(moves, [2]int{currentPos[0], currentPos[1]+1})
		moves = append(moves, [2]int{currentPos[0], currentPos[1]-1})
		fmt.Println("Borde inferior")
	}

	//borde izquierdo
	if currentPos [1] == 0 && currentPos [0] != 0 && currentPos [0] != e.Dimension-1{
		moves = append(moves, [2]int{currentPos[0]+1, currentPos[1]})
		moves = append(moves, [2]int{currentPos[0]-1, currentPos[1]})
		moves = append(moves, [2]int{currentPos[0], currentPos[1]+1})
		fmt.Println("Borde izquierdo")
	}

	//borde derecho
	if currentPos [1] == e.Dimension-1 && currentPos [0] != 0 && currentPos [0] != e.Dimension-1{
		moves = append(moves, [2]int{currentPos[0]+1, currentPos[1]})
		moves = append(moves, [2]int{currentPos[0]-1, currentPos[1]})
		moves = append(moves, [2]int{currentPos[0], currentPos[1]-1})
		fmt.Println("Borde derecho")
	}

	//centro
	if currentPos [0] != 0 && currentPos [0] != e.Dimension-1 && currentPos [1] != 0 && currentPos [1] != e.Dimension-1{
		moves = append(moves, [2]int{currentPos[0]+1, currentPos[1]})
		moves = append(moves, [2]int{currentPos[0]-1, currentPos[1]})
		moves = append(moves, [2]int{currentPos[0], currentPos[1]+1})
		moves = append(moves, [2]int{currentPos[0], currentPos[1]-1})
		fmt.Println("Centro")
	}

	return moves
}

