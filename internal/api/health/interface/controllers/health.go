package controllers

import (
	"aws-lambda-go/internal/common/logger"
	"aws-lambda-go/internal/common/responses"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// HealthController estructura para manejar la ruta de Health
type HealthController struct {
}

// NewHealthController constructor para HealthController
func NewHealthController() *HealthController {
	return &HealthController{}
}

// GetHealth
func (c *HealthController) GetHealth(ctx *gin.Context) {

	entry := logger.FromContext(ctx.Request.Context())

	entry.Info("HealthController.GetHealth")

	customResponse := responses.Response{
		Status: http.StatusOK,
		Data: gin.H{
			"status":    "ok",
			"message":   "El servicio está en línea y funcionando correctamente.",
			"timestamp": time.Now().Unix(),
		},
		Metadata: gin.H{
			"trace_id":  "d316a340-9c0a-419c-ad25-b7fefcdda3ce",
			"caller_id": "000000",
		},
		Errors: nil,
	}

	// Se almacena el objeto para que el middleware lo procese
	ctx.JSON(http.StatusOK, customResponse)
}
