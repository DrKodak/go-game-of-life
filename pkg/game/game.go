package game

type GameState struct {
	Board [][] bool
	Running bool
	Cols int
	Rows int
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
		Cols: cols,
		Rows: rows,
	}
}

// Update game state
func (g *GameState) Step() {
	// update logic
}

func (g *GameState) IsRunning() bool {
	return g.Running
}

func (g *GameState) PauseGame() {
    g.Running = false
}