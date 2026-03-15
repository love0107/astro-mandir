package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/love0107/astro-mandir/service"
)

type PanchaangHandler struct {
	service *service.PanchaangService
}

func NewPanchaangHandler() *PanchaangHandler {
	return &PanchaangHandler{
		service: service.NewPanchaangService(),
	}
}

func (h *PanchaangHandler) GetToday(c *gin.Context) {
	// Read lat/lng from query params
	lat := c.Query("lat")
	lng := c.Query("lng")

	// Default to Lucknow if not provided
	if lat == "" || lng == "" {
		lat = "26.8467"
		lng = "80.9462"
	}

	data, err := h.service.GetToday(c.Request.Context(), lat, lng)
	if err != nil {
		c.JSON(500, gin.H{"error": "Panchang fetch nahi ho paya"})
		return
	}

	c.JSON(200, data)
}
