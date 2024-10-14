package routers

import (
	"da1ch1a/go-todo/pkg/presentation/handlers"
	"da1ch1a/go-todo/pkg/presentation/registry"
	"html/template"
	"io"

	// "text/template"

	"github.com/labstack/echo/v4"
)

// echo.Rendererインターフェースの実装
type Template struct {
	templates *template.Template
}

func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, "layout.html", data)
}
// echo.Rendererインターフェースの実装ここまで


func NewRouter(r *registry.Registry) *echo.Echo {
	// テンプレートのプレコンパイル
	t := &Template{
		templates: template.Must(template.ParseGlob("pkg/presentation/views/*.html")),
	}

	e := echo.New()
	e.Renderer = t

	taskHandler := handlers.TaskHandler{}

	e.GET("/tasks", taskHandler.List)
	e.GET("/json", taskHandler.Json)

	return e
}
