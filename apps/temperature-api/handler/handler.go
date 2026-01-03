package handler

import (
	"net/http"
	"temperature-api/service"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	service service.TemperatureInterface
}

func NewHandler(service service.TemperatureInterface) *Handler {
	return &Handler{
		service: service,
	}
}

func (h *Handler) RegisterRoutes(router *gin.Engine) {
	group := router.Group("/temperature")

	group.GET("", h.GetTemperature)
	group.GET("/:sensorId", h.GetTemperature)
}

func (h *Handler) GetTemperature(c *gin.Context) {
	sensorId := c.Param("sensorId")
	locationId := c.Query("location")

	resp, err := h.service.GetTemperature(sensorId, locationId)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})

		return
	}

	c.JSON(http.StatusOK, resp)
}
