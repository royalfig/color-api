package utils

import (
	"encoding/csv"
	"fmt"
	"os"

	"github.com/royalfig/color-name-api/models"
)

func LoadColorsFromCSV(filePath string) ([]models.Color, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, fmt.Errorf("error opening file: %w", err)
	}
	defer file.Close()

	reader := csv.NewReader(file)

	records, err := reader.ReadAll()
	if err != nil {
		return nil, fmt.Errorf("error reading file: %w", err)
	}

	var colors []models.Color
	for i, record := range records {
		if i == 0 {
			continue
		}

		colors = append(colors, models.Color{
			Name:       record[0],
			MatchColor: record[1],
			GoodName:   record[2],
		})

	}

	return colors, nil
}
