package game

type GameState struct {
	Board [][] bool
	Running bool
}

// Create and initialize a new game state, i.e. starting a new game or clearing
func NewGameState() *GameState {
	return &GameState{
		// Initialization for game state
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