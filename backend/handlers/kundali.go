package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/love0107/astro-mandir/internal/kundali"
	"github.com/love0107/astro-mandir/service"
)

type KundaliHandler struct {
	service *service.KundaliService
}

func NewKundaliHandler() *KundaliHandler {
	return &KundaliHandler{
		service: service.NewKundaliService(),
	}
}

func (h *KundaliHandler) GenerateKundali(c *gin.Context) {
	var req kundali.KundaliRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": "Request sahi nahi hai"})
		return
	}

	if req.Name == "" || req.DOB == "" {
		c.JSON(400, gin.H{"error": "Naam aur janm tithi zaroori hai"})
		return
	}

	data, err := h.service.Generate(c.Request.Context(), req)
	if err != nil {
		c.JSON(500, gin.H{"error": "Kundali nahi ban payi"})
		return
	}

	c.JSON(200, data)
}
