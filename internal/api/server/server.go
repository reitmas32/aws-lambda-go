package server

import (
	"aws-lambda-go/internal/api/health"
	"aws-lambda-go/internal/api/router"
	"aws-lambda-go/internal/common/middlewares"
	"aws-lambda-go/internal/core/settings"
	"fmt"

	"github.com/aws/aws-lambda-go/lambda"
	ginadapter "github.com/awslabs/aws-lambda-go-api-proxy/gin"
	"github.com/gin-gonic/gin"
)

var ginLambda *ginadapter.GinLambda

func Run() {

	app := setUpRouter()

	app.Run(fmt.Sprintf(":%d", settings.Settings.PORT))
}

func RunLambda() {

	r := setUpRouter()

	// Adaptar Gin a Lambda
	ginLambda = ginadapter.New(r)

	// Iniciar Lambda
	lambda.Start(ginLambda.Proxy)
}

func setUpRouter() *gin.Engine {

	app := router.NewRouter()

	app.Use(middlewares.TraceMiddleware())
	//app.Use(middlewares.CatcherMiddleware)
	app.Use(middlewares.LoggerMiddleware())

	health.SetupHealthModule(app)

	return app
}
