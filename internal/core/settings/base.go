package settings

import (
	"fmt"
	"log"
	"os"
	"reflect"
	"time"

	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	ENVIRONMENT  string `required:"false"`
	TYPE_HANDLER string `required:"false" default:"API"`
	PORT         int    `required:"false" default:"8080"`
	TIMESTAMP    string `required:"false"`
	APP_NAME     string `required:"false" default:"aws-lambda-go"`
}

var Settings Config
var EnvDir = ".envs"

func LoadDotEnv() {

	err := godotenv.Load(fmt.Sprintf("%s/.env.base", EnvDir))
	if err != nil {
		log.Printf("No %s file found, using system environment variables", fmt.Sprintf("%s/.env.base.base", EnvDir))
	}

	environment := os.Getenv("ENVIRONMENT")
	if environment == "" {
		log.Println("ENVIRONMENT is not set")
	}

	// Mapear el archivo .env.base correspondiente al entorno
	envFiles := map[string]string{
		"":            fmt.Sprintf("%s/.env", EnvDir),
		"local":       fmt.Sprintf("%s/.env.local", EnvDir),
		"development": fmt.Sprintf("%s/.env.dev", EnvDir),
		"production":  fmt.Sprintf("%s/.env.prod", EnvDir),
		"staging":     fmt.Sprintf("%s/.env.staging", EnvDir),
	}

	// Obtener el archivo de entorno correspondiente
	envFile, exists := envFiles[environment]
	if !exists {
		log.Printf("Environment '%s' is not supported. Must be one of: local, development, production, staging", environment)
	}

	// Cargar las variables desde el archivo correspondiente
	err = godotenv.Load(envFile)
	if err != nil {
		log.Printf("No %s file found, using system environment variables", envFile)
	} else {
		log.Printf("Loaded environment variables from %s", envFile)
	}
}

func LoadEnvs() {
	// Procesar las variables de entorno en la estructura Settings
	err := envconfig.Process("", &Settings)
	if err != nil {
		log.Fatalf("Error loading configuration: %v", err)
	}

	Settings.TIMESTAMP = time.Now().Format("2006-01-02 15:04:05")

	// Imprimir las Settings si el entorno es local o development
	if Settings.ENVIRONMENT == "local" || Settings.ENVIRONMENT == "development" {
		log.Println("Settings:")

		// Obtener el tipo y valor de la estructura Settings
		v := reflect.ValueOf(Settings)
		t := reflect.TypeOf(Settings)

		// Recorrer cada campo de la estructura
		for i := 0; i < v.NumField(); i++ {
			fieldName := t.Field(i).Name
			fieldValue := v.Field(i).Interface()
			fmt.Printf("\033[32m%s\033[0m: %v\n", fieldName, fieldValue)
		}
	}
}
