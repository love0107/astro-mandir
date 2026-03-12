package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/love0107/astro-mandir/db"
	"github.com/love0107/astro-mandir/handlers"
)

func main() {
	// Load .env file
	err := godotenv.Load()
	if err != nil {
		log.Println(".env file nahi mila — environment variables use karega")
	}

	// Now you can access anywhere in code like this
	// os.Getenv("PROKERALA_CLIENT_ID")

	db.InitDB()

	panchaangHandler := handlers.NewPanchaangHandler()
	bhajanHandler := handlers.NewBhajanHandler()
	rashifalHandler := handlers.NewRashifalHandler()
	kundaliHandler := handlers.NewKundaliHandler()

	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "AstroMandir chal raha hai! 🙏",
		})
	})

	api := r.Group("/api")
	{
		api.GET("/today", panchaangHandler.GetToday)
		api.GET("/bhajan", bhajanHandler.GetTodayBhajan)
		api.GET("/rashifal/:rashi", rashifalHandler.GetRashifal)
		api.POST("/kundali", kundaliHandler.GenerateKundali)
	}

	log.Println(" Server starting on :8080")
	r.Run(":" + getPort())
}

func getPort() string {
	port := os.Getenv("PORT")
	if port == "" {
		return "8080"
	}
	return port
}
