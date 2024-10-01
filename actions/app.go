package actions

import (
	"os"
	"sync"

	"go_buffalo/models"

	"github.com/gobuffalo/buffalo"
	"github.com/gobuffalo/buffalo-pop/v3/pop/popmw"
	"github.com/gobuffalo/envy"
	"github.com/gobuffalo/middleware/contenttype"
	"github.com/gobuffalo/middleware/i18n"
	"github.com/gobuffalo/middleware/paramlogger"
	"github.com/gobuffalo/x/sessions"
	"github.com/rs/cors"
)

var ENV = envy.Get("GO_ENV", "development")

var (
	app     *buffalo.App
	appOnce sync.Once
	T       *i18n.Translator
)

func App() *buffalo.App {
	appOnce.Do(func() {
		os.Setenv("PORT", "8001")
		app = buffalo.New(buffalo.Options{
			Env:          ENV,
			SessionStore: sessions.Null{},
			PreWares: []buffalo.PreWare{
				cors.Default().Handler,
			},
			SessionName: "_go_buffalo_session",
		})

		app.Use(paramlogger.ParameterLogger)

		app.Use(contenttype.Set("application/json"))

		app.Use(popmw.Transaction(models.DB))

		app.GET("/", HomeHandler)

		app.GET("/application", ApplicationIndex)
		app.GET("/application/{id}", ApplicationShow)
	})

	return app
}
