package game

import "fmt"

type GameState struct {
	// [rows][cols]
	Board [][] bool
	Running bool
	TotalCols int
	TotalRows int
}

// Create and initialize a new game state, i.e. starting a new game or clearing
func NewGameState() *GameState {
	const cols = 12
	const rows = 12
	row := []bool{rows: false}
	a := make([][]bool, cols)
	for i := range a {
		a[i] = row
	}	
	return &GameState{
		// Initialization for game state
		Board: a,
		Running: false,
		TotalCols: cols,
		TotalRows: rows,
	}
}

func (g *GameState) countLiveNeighbors(curRow int, curCol int) int {
	count := 0
	// for each cell, check col-1, col, col+1 for row-1, row, row+1 as long as not 0 and greater than total
	for r:= curRow - 1; r <= curRow + 1; r++ {
		for c:= curCol - 1; c <= curCol + 1; c++ {
			if r < 0 || c < 0 {
				continue
			}
			if r > g.TotalRows || c > g.TotalCols {
				continue
			}
			if g.Board[r][c] {
				count++
			}
		}
	} 

	return count
}

// Update game state
func (g *GameState) Step() {
	// update logic
}

func (g *GameState) IsRunning() bool {
	return g.Running
}

func (g *GameState) StartGame() {
	g.Running = true
	fmt.Println("hello i'm starting now")
}

func (g *GameState) PauseGame() {
    g.Running = false
	fmt.Println("hello I'm pausing now")
}