package main

import (
	"context"

	"github.com/gin-gonic/gin"
	"github.com/imyashkale/microforge/pkg/configs"
	"github.com/imyashkale/microforge/pkg/dynamodb"
	"github.com/imyashkale/microforge/pkg/log"
)

var (
	VERSION = "1.0.0"
)

func main() {

	app := &Application{
		logger:    log.New(),
		dyanamodb: dynamodb.New(),
		config:    configs.New(),
		router:    gin.New(),
	}

	// setting the application version
	app.logger.With(context.TODO(), "version", VERSION)

	// loading configuration
	if err := app.config.Load(app.logger); err != nil {
		app.logger.Errorf("loading configuration failed %s", err.Error())
		return
	}

	// middlewares are attached
	app.middlewares()

	// routes are attached
	app.routes()

	// starting the application
	app.Run()

}
