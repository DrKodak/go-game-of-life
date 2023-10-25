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

// func getUser(c echo.Context) error {
// 	// User ID From path `users/:id`
// 	id := c.Param("id")
// 	return c.String(http.StatusOK, id)
// }

// func show(c echo.Context) error {
// 	team := c.QueryParam("team")
// 	member := c.QueryParam("member")
// 	return c.String(http.StatusOK, "team: " + team + ", member: " + member)
// }

// func save(c echo.Context) error {
// 	// get name and email
// 	name := c.FormValue("name")
// 	email := c.FormValue("email")
// 	return c.String(http.StatusOK, "name: " + name + "email: " + email)
// }

func TimeLoggingMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		start := time.Now()
		err := next(c)
		duration := time.Since(start)
		log.Printf("Processed request in %s", duration)
		return err
	}
}

type Template struct {
	templates *template.Template
}

func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

func Hello(c echo.Context) error {
	return c.Render(http.StatusOK, "hello", "World")
}

func Index(c echo.Context) error {
	return c.Render(http.StatusOK, "index", "index....")
}

func main() {
	e := echo.New()
	e.Static("/static", "static")

	// Root level middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Templates
	t := &Template{
		// for now do that but the docs say basically "templates/*.html"
		templates: template.Must(template.ParseGlob("templates/*.html")),
	}
	e.Renderer = t;

	e.GET("/hello", Hello)
	e.GET("/", Index)

	// e.GET("/", func(c echo.Context) error {
	// 	return c.String(http.StatusOK, "Hello life!")
	// })

	e.POST("/clicked", func(c echo.Context) error {
		return c.HTML(http.StatusOK, `<p>You clicked me!</p>`)
	})
	// e.POST("/users", saveUser)
	// e.GET("/users/:id", getUser)
	// e.GET("/show", show)
	// e.POST("/save", save)
	// e.PUT("/users/:id", updateUser)

	// e.DELETE("users/:id", deleteUser)
	e.Logger.Fatal(e.Start(":1323"))
}