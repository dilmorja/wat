package main

import(
	"html/template"
	"io"
	"net/http"
	"path/filepath"
	"flag"

	"github.com/labstack/echo/v4"
)

var(
	FPORT = flag.String("port", "8080", "port to listen on")
)

func init() { flag.Parse() }

type T struct {
	*template.Template
}

func (t *T) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.ExecuteTemplate(w, name, data)
}

func main() {
	app := echo.New()

	renderer := &T{template.Must(template.ParseGlob(filepath.Join("public", "views", "*.html")))}

	app.Static("/static", filepath.Join("public", "static"))
	app.File("/bin/go.wasm", filepath.Join("bin", "go.wasm"))

	app.Renderer = renderer

	app.GET("/", func(c echo.Context) error {
		return c.Render(http.StatusOK, "Home", map[string]interface{}{
			"name": "tester",
		})
	})

	app.Start(*FPORT)
}