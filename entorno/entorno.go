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
	DirtCount       int
	CleanCellCount  int
}

// NewEntorno crea una nueva instancia de Entorno.
func NewEntorno(Dimension int, probabilityDirt float64) *Entorno {
	e := &Entorno{
		Dimension:       Dimension,
		probabilityDirt: probabilityDirt,
		CleanCellCount:  0,
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
	e.DirtCount = 0
	for i := 0; i < e.Dimension; i++ {
		var row []bool
		for j := 0; j < e.Dimension; j++ {
			random := rand.Float64()
			if random < e.probabilityDirt {
				row = append(row, true)
				e.DirtCount++
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
	e.CleanCellCount++
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

// Regresa el contador de celdas limpiadas
func (e *Entorno) GetCleanCount() int {
	return e.CleanCellCount
}
