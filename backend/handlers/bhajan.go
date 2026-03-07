package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/love0107/astro-mandir/service"
)

type BhajanHandler struct {
	service *service.BhajanService
}

func NewBhajanHandler() *BhajanHandler {
	return &BhajanHandler{
		service: service.NewBhajanService(),
	}
}

func (h *BhajanHandler) GetTodayBhajan(c *gin.Context) {
	// Optional — rashi for personalization
	rashi := c.Query("rashi")

	data, err := h.service.GetTodayBhajan(c.Request.Context(), rashi)
	if err != nil {
		c.JSON(500, gin.H{"error": "Bhajan fetch nahi ho paya"})
		return
	}

	c.JSON(200, data)
}