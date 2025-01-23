package services

import (
	"fmt"
	"strings"
	"time"

	"github.com/royalfig/color-name-api/models"
	"github.com/royalfig/color-name-api/utils"
	"golang.org/x/exp/rand"
)

func ParsePalette(palette string) ([]models.Color, error) {
	if strings.TrimSpace(palette) == "" {
		return nil, fmt.Errorf("palette string is empty")
	}

	var colors []models.Color
	hexValues := strings.Split(palette, ",")

	for _, str := range hexValues {
		hex, err := utils.NormalizeHex(strings.TrimSpace(str))
		if err != nil {
			return nil, err
		}
		color, err := FindColorByHex(hex)
		if err == nil {
			color.MatchType = "exact"
			color.OriginalColor = hex
			colors = append(colors, color)
			continue
		}

		nearest, err := FindClosestColor(hex)

		if err != nil {
			return nil, err
		}
		nearest.MatchType = "nearest"
		nearest.OriginalColor = hex
		colors = append(colors, nearest)
	}

	return colors, nil
}

func GeneratePaletteName(colors []models.Color) string {
	rand.Seed(uint64(time.Now().UnixNano()))

	var allWords [][]string
	for _, color := range colors {
		words := utils.ExtractValidWords((color.Name))
		if len(words) > 0 {
			allWords = append(allWords, words)
		}
	}

	if len(allWords) < 2 {
		return colors[0].Name // Return the first color name if there are not enough words
	}

	fmt.Printf("allWords: %v\n", allWords)

	firstIndex := rand.Intn(len(allWords))
	secondIndex := rand.Intn(len(allWords))
	fmt.Print("firstIndex: ", firstIndex, " secondIndex: ", secondIndex, "\n")
	// Compare memory addresses to ensure we don't pick the same words
	for firstIndex == secondIndex {
		secondIndex = rand.Intn(len(allWords))
	}

	firstColorWords := allWords[firstIndex]
	secondColorWords := allWords[secondIndex]

	firstWord := firstColorWords[rand.Intn(len(firstColorWords))]
	secondWord := secondColorWords[rand.Intn(len(secondColorWords))]

	return fmt.Sprintf("%s %s", firstWord, secondWord)
}
