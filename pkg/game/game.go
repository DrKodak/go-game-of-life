package game

import (
	"fmt"
	"math/rand"
)

type GameState struct {
	// [rows][cols]
	Board [][] bool
	NextBoard [][] bool
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
		NextBoard: a,
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
			if r >= g.TotalRows || c >= g.TotalCols {
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
	fmt.Print("hello I'm stepping\n")
	// update logic
	g.NextBoard = g.Board
	for r:= 0; r < g.TotalRows; r++ {
		for c:= 0; c < g.TotalCols; c++ {
			count := g.countLiveNeighbors(r, c)
			if count > 0 {
				if g.Board[r][c] {
					// if count < 2, dies
					if count < 2 {
						g.NextBoard[r][c] = false	
					}
					// if count > 3, dies
					if count > 3 {
						g.NextBoard[r][c] = false	
					}
				} else {
					// if dead and count > 3, cell lives
					if count > 3 {
						g.NextBoard[r][c] = true
					}
				}
			} // if count > 0
		} // end for cols
	} // end for rows
	// Updates the actual Board
	fmt.Print("done doing the stepping\n")
	g.Board = g.NextBoard
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

func randomBool() bool {
	return rand.Float32() < 0.5
}

func (g *GameState) RandomizeBoard() {
	fmt.Print("randomizing\n")
	for r:= 0; r < g.TotalRows; r++ {
		for c:= 0; c < g.TotalCols; c++ {
			g.Board[r][c] = randomBool()
		}
	}
}