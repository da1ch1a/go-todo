package main

import (
	"da1ch1a/go-todo/pkg/infra"
	"da1ch1a/go-todo/pkg/presentation/registry"
	"da1ch1a/go-todo/pkg/presentation/routers"
	"io"
	"log"
	"text/template"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
)

func main() {
	envPath := "config/env/.env"
	err := godotenv.Load(envPath)
  if err != nil {
    log.Fatal("Error loading .env file")
  }

	// 初期化
	r := initializeRegistry()
	e := routers.NewRouter(r)

	e.Logger.Fatal(e.Start(":1323"))
}

// echo.Rendererインターフェースの実装
type Template struct {
	templates *template.Template
}

func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, "layout.html", data)
}

func initializeRegistry() *registry.Registry {
	// DBの初期化
	db := infra.InitDb()

	return registry.BuildRegistry(db)
}
