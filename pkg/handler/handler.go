package handler

import (
	"net/http"

	"github.com/DrKodak/go-game-of-life/pkg/game"
	"github.com/DrKodak/go-game-of-life/pkg/renderer"
	"github.com/labstack/echo/v4"
)

type GameHandler struct {
	State *game.GameState
}

// Create a new game handler using a game state
func NewGameHandler(gs *game.GameState) *GameHandler {
	return &GameHandler{
		State: gs,
	}
}

// Function to be called from the POST to pause the game
func (gh *GameHandler) PauseGame(c echo.Context) error {
	gh.State.PauseGame()
	return nil
}

func (gh *GameHandler) StepGame(c echo.Context) error {
	gh.State.Step()
	return nil
}

func (gh *GameHandler) UpdateGame(c echo.Context) error {
	// TODO Might be update instead of a step
	if (gh.State.IsRunning()) {

		gh.State.Step()
		
		// Render the new state
		html, err := renderer.RenderState(gh.State)
		if err != nil {
			return err
		}
		
		return c.HTML(http.StatusOK, string(html))
	}
	return nil
}

func (gh *GameHandler) RenderIndex(c echo.Context) error {
	// If I wanted to add some special initial state I could through a call to the Renderer
	return c.File("templates/index.html")
}