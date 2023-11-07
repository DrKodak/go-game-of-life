package renderer

import (
	"html/template"

	"github.com/DrKodak/go-game-of-life/pkg/game"
)

func RenderState(g *game.GameState) (template.HTML, error) {
	// Convert Game state to html
	println(g.IsRunning())
	return template.HTML("<div>Game State Render</div>"), nil 
}