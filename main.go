package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/royalfig/color-name-api/handlers"
	"github.com/royalfig/color-name-api/services"
)

func main() {
	if err := services.LoadColors("data/colornames.csv"); err != nil {
		log.Fatalf("failed to load colors: %v", err)
	}

	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowOrigins: []string{"http://colorpalette.pro", "http://localhost:5173"},
		AllowMethods: []string{"GET"},
		AllowHeaders: []string{"Origin", "Content-Type"},
		ExposeHeaders: []string{"Content-Length"},
		MaxAge: 12 * time.Hour
	}))

	r.GET("/color/:hex", handlers.GetColorName)
	r.GET("/palette/:hex", handlers.GetPaletteNames)

	log.Println("Starting server on :8080...")
	r.Run(":8080")
}
