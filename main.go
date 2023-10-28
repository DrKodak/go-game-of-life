package main

import (
	"io"
	"log"
	"net/http"
	"text/template"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func TimeLoggingMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		start := time.Now()
		err := next(c)
		duration := time.Since(start)
		log.Printf("Processed request in %s", duration)
		return err
	}
}

func seq(end int) []int {
	var seq []int
	start := 0
	for i := start; i < end; i++ {
		seq = append(seq, i)
	}
	return seq
}

type Template struct {
	templates *template.Template
}

// TemplateRenderer is a custom html/template renderer for Echo
type TemplateRenderer struct {
	templates *template.Template
}

// Data needed for the grid structure
type GridData struct {
	Cols int
	CellCount int
}

func (t *TemplateRenderer) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	// Add global methods if data is a map
	if viewContext, isMap := data.(map[string]interface{}); isMap {
		viewContext["reverse"] = c.Echo().Reverse;
	}
	return t.templates.ExecuteTemplate(w, name, data)
}

func Index(c echo.Context) error {
	return c.Render(http.StatusOK, "index", "index ")
}

func Foobar(c echo.Context) error {
	return c.Render(http.StatusOK, "index.html", map[string]interface{}{
		"name":"Dolly!",
	})
}

func ColCount(c echo.Context) error {
	data := GridData {
		Cols: 10,
		CellCount: 100,
	}
	return c.Render(http.StatusOK, "index.html", data)
}

func main() {
	e := echo.New()
	e.Static("/static", "static")

	// Root level middleware
	e.Use(middleware.Logger())
	// e.Use(middleware.Recover())

	// Templates
	renderer := &TemplateRenderer{
		templates: template.Must(template.New("index.html").Funcs(template.FuncMap{
			"seq": seq,
		}).ParseFiles("templates/index.html")),
	}
	e.Renderer = renderer

	e.GET("/", ColCount)
	// e.GET("/", Index)

	e.POST("/pickle", func(c echo.Context) error {
		return c.HTML(http.StatusOK, `<div
        class="w-9 h-9 bg-gray-600 border border-gray-300 flex items-center justify-center"
        hx-post="/notpickle"
        hx-swap="outerHTML"
      >
	  🥒
      </div>`)
	})

	e.POST("/noPickle", func(c echo.Context) error {
		return c.HTML(http.StatusOK, `<div
        class="w-9 h-9 bg-gray-600 border border-gray-300 flex items-center justify-center"
        hx-post="/pickle"
        hx-swap="outerHTML"
      >
      </div>`)
	})

	e.POST("/clicked", func(c echo.Context) error {
		return c.HTML(http.StatusOK, `<p>You clicked me!</p>`)
	})

	e.Logger.Fatal(e.Start(":1323"))
}