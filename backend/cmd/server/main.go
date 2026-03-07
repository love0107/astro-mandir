package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/love0107/astro-mandir/db"
)

func main() {
	// Database initialize
	db.InitDB()

	// Router
	r := gin.Default()

	// Test route
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "AstroMandir chal raha hai! 🙏",
			"db":      "connected ✅",
		})
	})

	log.Println("Server starting on :8080")
	r.Run(":8080")
}
