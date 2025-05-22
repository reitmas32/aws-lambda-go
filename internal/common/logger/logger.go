package logger

import (
	"aws-lambda-go/internal/core/settings"
	"context"
	"fmt"
	"time"

	"github.com/mattn/go-colorable"
	"github.com/sirupsen/logrus"
)

const loggerKey = "logger"

// CustomFormatter define un formateador personalizado para logrus con colores
type CustomFormatter struct{}

func (f *CustomFormatter) Format(entry *logrus.Entry) ([]byte, error) {
	// Obtener campos personalizados
	requestID, _ := entry.Data["trace_id"].(string)
	sessionID, _ := entry.Data["caller_id"].(string)
	path, _ := entry.Data["path"].(string)

	// Colores ANSI para diferentes niveles
	var levelColor string
	switch entry.Level {
	case logrus.DebugLevel:
		levelColor = "\033[36m" // Cyan
	case logrus.InfoLevel:
		levelColor = "\033[32m" // Green
	case logrus.WarnLevel:
		levelColor = "\033[33m" // Yellow
	case logrus.ErrorLevel, logrus.FatalLevel, logrus.PanicLevel:
		levelColor = "\033[31m" // Red
	default:
		levelColor = "\033[0m" // Reset
	}

	// Crear el log con el formato especificado y colores
	log := fmt.Sprintf(
		"%s | %s%s\033[0m | %s | %s | %s | %s:%d - %s\n",
		time.Now().Format("2006-01-02 15:04:05"),
		levelColor, entry.Level.String(), // Nivel con color
		requestID,
		sessionID,
		path,
		entry.Caller.Function,
		entry.Caller.Line,
		entry.Message,
	)

	return []byte(log), nil
}

// globalLogger es el logger principal
var globalLogger *logrus.Logger

func init() {
	// Configuración inicial del logger
	globalLogger = logrus.New()
	globalLogger.SetFormatter(&CustomFormatter{})
	globalLogger.SetLevel(logrus.InfoLevel)
	globalLogger.SetReportCaller(true) // Habilita la información del caller (función y línea)

	// Soporte para colores en todas las plataformas
	globalLogger.SetOutput(colorable.NewColorableStdout())
}

// WithFields crea un logger con campos adicionales
func WithFields(fields map[string]interface{}) *logrus.Entry {
	return globalLogger.WithFields(fields)
}

// WithLogger agrega un logger personalizado al contexto
func WithLogger(ctx context.Context, entry *logrus.Entry) context.Context {
	return context.WithValue(ctx, loggerKey, entry)
}

// FromContext obtiene el logger contextualizado desde el contexto
func FromContext(ctx context.Context) *logrus.Entry {
	entry, ok := ctx.Value(loggerKey).(*logrus.Entry)

	if ok {
		return entry
	}

	// Si no hay logger contextualizado, devuelve el logger global
	return globalLogger.WithFields(logrus.Fields{})
}

// Info escribe un mensaje a nivel INFO
func Info(message string) {
	globalLogger.Info(message)
}

// Error escribe un mensaje a nivel ERROR
func Error(message string) {
	globalLogger.Error(message)
}

// Debug escribe un mensaje a nivel DEBUG
func Debug(message string) {
	globalLogger.Debug(message)
}

// Info escribe un mensaje a nivel INFO
func InfoDelicate(message string) {
	if settings.Settings.ENVIRONMENT != "production" {
		globalLogger.Info(message)
	}
}
