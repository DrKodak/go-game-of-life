package main

import (
	"html/template"
	"io"

	"github.com/DrKodak/go-game-of-life/pkg/game"
	"github.com/DrKodak/go-game-of-life/pkg/handler"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// TemplateRenderer is a custom html/template renderer for Echo Framework
type TemplateRenderer struct {
	templates *template.Template
}

// Render function to fit with the interface
func (t *TemplateRenderer) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

func main() {
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Static files
	e.Static("/static", "static")

	// Template Render
	renderer := &TemplateRenderer{
		templates: template.Must(template.ParseFiles("templates/index.html")),
	}
	e.Renderer = renderer

	// Initialize the game
	gs := game.NewGameState()

	// Set up the game state handler
	gh := handler.NewGameHandler(gs)

	// Define all routes
	e.GET("/", gh.RenderIndex)
	e.POST("/pause-game", gh.PauseGame)
	e.POST("/step-game", gh.StepGame)

	// Start Server
	e.Logger.Fatal(e.Start(":42069"))
}