package main

import (
	"da1ch1a/go-todo/pkg/infra"
	"da1ch1a/go-todo/pkg/presentation/handlers"
	"da1ch1a/go-todo/pkg/presentation/registry"
	"da1ch1a/go-todo/pkg/presentation/routers"
	"io"

	"text/template"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/rs/zerolog/log"
)

func main() {
	envPath := "config/env/.env"
	err := godotenv.Load(envPath)
	if err != nil {
		log.Fatal().Msg("Error loading .env file")
	}

	// 初期化
	r := initializeRegistry()
	e := routers.NewRouter(r)

	// エラーハンドラの設定
	e.HTTPErrorHandler = handlers.ErrorHandler

	// access log ドキュメント通り
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "method=${method}, uri=${uri}, status=${status}\n",
	}))

	e.Logger.Fatal(e.Start(":1323"))
}

// TODO フロントがある程度形になったら消す
// echo.Rendererインターフェースの実装
type Template struct {
	templates *template.Template
}

// TODO フロントがある程度形になったら消す
func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, "layout.html", data)
}

func initializeRegistry() *registry.Registry {
	// DBの初期化
	db := infra.InitDb()

	return registry.BuildRegistry(db)
}
