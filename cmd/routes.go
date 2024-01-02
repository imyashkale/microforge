package main

import (
	"github.com/imyashkale/microforge/internal/healthcheck"
	"github.com/imyashkale/microforge/internal/dummy"
)

// attaching handlers
func (app *Application) routes() {
	root := app.router.Group(app.config.Application.Endpoint)
	v1 := root.Group("/v1")
	{
		// Healthchecks
		healthcheck.RegisterHandlers(v1)

		// dummy handler
		dummy.RegisterDummyHandler(
			v1,
			dummy.NewService(
				dummy.NewRepository(app.config.Dynamodb, app.config.Dynamodb.TableName),
				app.logger,
			),
			app.logger,
		)

	}
}
