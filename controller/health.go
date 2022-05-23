package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

var HealthVar HealthType = HealthType{message: "API server is Healthy", status: 200}

type HealthController interface {
	Health()
}

type HealthType struct {
	message string
	status  int
}

func (h *HealthType) Health(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"Message": h.message,
		"status":  h.status,
	})
}
