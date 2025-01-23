package utils

import (
	"fmt"
	"strconv"
	"strings"
)

func NormalizeHex(hex string) (string, error) {

	hex = strings.TrimPrefix(hex, "#")

	if len(hex) == 3 {
		hex = fmt.Sprintf("%c%c%c%c%c%c", hex[0], hex[0], hex[1], hex[1], hex[2], hex[2])
	}

	if len(hex) != 6 {
		return "", fmt.Errorf("invalid hex length: %d", len(hex))
	}

	return "#" + strings.ToLower(hex), nil
}

func HexToRGB(hex string) (int, int, int, error) {
	hex, err := NormalizeHex(hex)
	if err != nil {
		return 0, 0, 0, err
	}

	r, err := strconv.ParseInt(hex[1:3], 16, 0)
	if err != nil {
		return 0, 0, 0, fmt.Errorf("error parsing red: %w", err)
	}

	g, err := strconv.ParseInt(hex[3:5], 16, 0)
	if err != nil {
		return 0, 0, 0, fmt.Errorf("error parsing green: %w", err)
	}

	b, err := strconv.ParseInt(hex[5:7], 16, 0)
	if err != nil {
		return 0, 0, 0, fmt.Errorf("error parsing blue: %w", err)
	}

	return int(r), int(g), int(b), nil
}
