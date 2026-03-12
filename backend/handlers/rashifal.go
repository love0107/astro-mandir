package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/love0107/astro-mandir/service"
)

type RashifalHandler struct {
	service *service.RashifalService
}

func NewRashifalHandler() *RashifalHandler {
	return &RashifalHandler{
		service: service.NewRashifalService(),
	}
}

func (h *RashifalHandler) GetRashifal(c *gin.Context) {
	// Get rashi from URL — /api/rashifal/mesh
	rashi := c.Param("rashi")

	if rashi == "" {
		c.JSON(400, gin.H{"error": "Rashi batao — mesh, vrishabh, etc."})
		return
	}

	data, err := h.service.GetRashifal(c.Request.Context(), rashi)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, data)
}
