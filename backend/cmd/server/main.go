package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/love0107/astro-mandir/db"
	"github.com/love0107/astro-mandir/handlers"
)

func main() {
	// Init DB
	db.InitDB()

	// Init handlers
	panchaangHandler := handlers.NewPanchaangHandler()

	// Router
	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "AstroMandir chal raha hai! 🙏"})
	})

	// API routes
	api := r.Group("/api")
	{
		api.GET("/today", panchaangHandler.GetToday)
	}

	log.Println("Server starting on :8080")
	r.Run(":8080")
}
