package services

import (
	"errors"
	"math"

	"github.com/royalfig/color-name-api/models"
	"github.com/royalfig/color-name-api/utils"
)

var colorData []models.Color

func LoadColors(filePath string) error {
	var err error
	colorData, err = utils.LoadColorsFromCSV(filePath)
	return err
}

func FindColorByHex(hex string) (models.Color, error) {
	hex, err := utils.NormalizeHex(hex)
	if err != nil {
		return models.Color{}, err
	}

	for _, color := range colorData {
		if color.MatchColor == hex {
			return color, nil
		}
	}
	return models.Color{}, errors.New("color not found")
}

func FindClosestColor(hex string) (models.Color, error) {
	hex, err := utils.NormalizeHex(hex)
	if err != nil {
		return models.Color{}, err
	}

	r, g, b, err := utils.HexToRGB(hex)
	if err != nil {
		return models.Color{}, err
	}

	var closestColor models.Color
	minDistance := math.MaxFloat64

	for _, color := range colorData {
		cr, cg, cb, _ := utils.HexToRGB(color.MatchColor)
		distance := utils.CalculateColorDistance(r, g, b, cr, cg, cb)
		if distance < minDistance {
			minDistance = distance
			closestColor = color
		}
	}

	return closestColor, nil
}
