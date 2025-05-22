package main

import (
	"aws-lambda-go/internal/api/server"
	"aws-lambda-go/internal/core/settings"
	"fmt"
)

func main() {
	fmt.Println("aws-lambda-go v0.0.1")

	settings.LoadDotEnv()

	settings.LoadEnvs()

	if settings.Settings.TYPE_HANDLER == "API" {
		fmt.Println("Starting API server")
		fmt.Printf("Listening on port %d\n", settings.Settings.PORT)
		server.Run()
	} else {
		fmt.Println("Starting Lambda server")
		fmt.Printf("Listening on port %d\n", settings.Settings.PORT)
		server.RunLambda()
	}
}
