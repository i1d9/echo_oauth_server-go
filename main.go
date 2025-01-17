package main

import (
	"html/template"
	"io"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
)

// Template structure to implement Render interface
type Template struct {
	templates *template.Template
}

// Render method for Template struct
func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

func main() {
	// Initialize Echo
	e := echo.New()

	// Load templates
	t := &Template{
		templates: template.Must(template.ParseGlob("templates/*.html")),
	}
	e.Renderer = t

	// Routes
	e.GET("/", func(c echo.Context) error {
		data := map[string]interface{}{
			"title": "Home Page",
			"message": "Welcome to the Echo Framework!",
		}
		return c.Render(http.StatusOK, "index.html", data)
	})


	
	e.GET("/login", func(c echo.Context) error {
		data := map[string]interface{}{
			"title": "About Page",
			"message": "This is the about page.",
		}
		return c.Render(http.StatusOK, "about.html", data)
	})


	e.GET("/register", func(c echo.Context) error {
		data := map[string]interface{}{
			"title": "About Page",
			"message": "This is the about page.",
		}
		return c.Render(http.StatusOK, "about.html", data)
	})


	// Start the server
	log.Fatal(e.Start(":8080"))
}
