package main

import (
	"fmt"
	"html/template"
	"io"

	"github.com/DrKodak/go-game-of-life/pkg/game"
	"github.com/DrKodak/go-game-of-life/pkg/handler"
	"github.com/DrKodak/go-game-of-life/pkg/renderer"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// TemplateRenderer is a custom html/template renderer for Echo Framework
type TemplateRenderer struct {
	templates *template.Template
}

// Render function to fit with the interface
func (t *TemplateRenderer) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	// add global methods if data is a map
	if viewContext, isMap := data.(map[string]interface{}); isMap {
		viewContext["reverse"] = c.Echo().Reverse
	}
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
	templateRenderer := &TemplateRenderer{
		templates: template.Must(template.New("index").Funcs(template.FuncMap{
			"Sequence": renderer.Seq,
		}).ParseFiles("templates/index.html")),
	}
	// templateRenderer := &TemplateRenderer{
	// 	templates: template.Must(
	// 	template.ParseFiles("templates/index.html")),
	// }
	e.Renderer = templateRenderer

	// Initialize the game
	gs := game.NewGameState()

	// Set up the game state handler
	gh := handler.NewGameHandler(gs)

	// Define all routes
	e.GET("/", gh.RenderIndex)
	e.POST("/pause-game", gh.PauseGame)
	e.POST("/step-game", gh.StepGame)

	fmt.Printf("%#v\n", templateRenderer.templates)
	// Start Server
	e.Logger.Fatal(e.Start(":42069"))
}