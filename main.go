package main

import (
	"da1ch1a/go-todo/pkg/presentation/routers"
	"io"
	"text/template"

	"github.com/labstack/echo/v4"
)

func main() {
	// テンプレートのプレコンパイル
	t := &Template{
		templates: template.Must(template.ParseGlob("pkg/presentation/views/*.html")),
	}

	e := echo.New()
	e.Renderer = t
	routers.Router(e)

	e.Logger.Fatal(e.Start(":1323"))
}

// echo.Rendererインターフェースの実装
type Template struct {
	templates *template.Template
}

func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}
