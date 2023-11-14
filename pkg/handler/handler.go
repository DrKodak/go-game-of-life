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
// Needs to update server-side game state and update the button
func (gh *GameHandler) PauseGame(c echo.Context) error {
	gh.State.PauseGame()
	return c.HTML(http.StatusOK, `<button
	class="bg-purple-600 text-gray px-4 py-2 rounded shadow-md" hx-post="/start-game"
	hx-swap="outerHTML" >Start ▶️<button>`)
}

func (gh *GameHandler) StartGame(c echo.Context) error {
	gh.State.StartGame()
	return c.HTML(http.StatusOK, `<button
	class="bg-purple-600 text-gray px-4 py-2 rounded shadow-md" hx-post="/pause-game"
	hx-swap="outerHTML" >Pause ⏸️<button>`)
}

// Function to be called from POST
// Needs to update server-side game state by calling State's step function. Should this also call UpdateGame after?
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
	// return c.File("templates/index.html")
	return c.Render(http.StatusOK, "index", nil)
}