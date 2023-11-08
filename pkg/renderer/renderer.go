package renderer

import (
	"fmt"
	"html/template"

	"github.com/DrKodak/go-game-of-life/pkg/game"
)

func Seq(end int) []int {
	var seq []int
	start := 0
	for i := start; i < end; i++ {
		seq = append(seq, i)
	}
	return seq
}

func PrintGameState(g *game.GameState) {
	for row := 0; row < g.Rows; row++ {
		for col := 0; col < g.Cols; col++ {
			fmt.Print(g.Board[row][col], " ")
		}
		fmt.Print("\n")
	}
}

func RenderState(g *game.GameState) (template.HTML, error) {
	// Convert Game state to html
	println(g.IsRunning())
	return template.HTML("<div>Game State Render</div>"), nil 
}