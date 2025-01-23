package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/royalfig/color-name-api/services"
)

func GetColorName(c *gin.Context) {
	hex := c.Param("hex")
	color, err := services.FindColorByHex(hex)
	if err == nil {
		color.MatchType = "exact"
		color.OriginalColor = hex
		c.JSON(http.StatusOK, gin.H{"color": color})
		return
	}

	closestColor, err := services.FindClosestColor(hex)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "color not found"})
		return
	}
	closestColor.MatchType = "nearest"
	closestColor.OriginalColor = hex

	c.JSON(http.StatusOK, gin.H{"color": closestColor})
}

func GetPaletteNames(c *gin.Context) {
	hexColors := c.Param("hex")
	colors, err := services.ParsePalette(hexColors)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	paletteName := services.GeneratePaletteName(colors)

	c.JSON(http.StatusOK, gin.H{"palette_name": paletteName, "colors": colors})
}
