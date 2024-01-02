package main

import (
	"github.com/gin-contrib/cors"
	errors "github.com/imyashkale/microforge/pkg/errors"
)

// middlewares
func (app *Application) middlewares() {
	// CORS
	app.router.Use(cors.Default())
	app.router.Use(errors.Handler(app.logger))
}
