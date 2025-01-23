package utils

import (
	"os"
	"reflect"
	"testing"

	"github.com/royalfig/color-name-api/models"
)

func TestLoadColorsFromCSV(t *testing.T) {
	testCSV := `name,hex,good_name
100 Mph,#c93f38,x
18th Century Green,#a59344,
1975 Earth Red,#7b463b,`

	expected := []models.Color{
		{Name: "100 Mph", MatchColor: "#c93f38", GoodName: "x"},
		{Name: "18th Century Green", MatchColor: "#a59344", GoodName: ""},
		{Name: "1975 Earth Red", MatchColor: "#7b463b", GoodName: ""},
	}

	tempFile := "test_colors.csv"
	err := os.WriteFile(tempFile, []byte(testCSV), 0644)
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(tempFile)

	actual, err := LoadColorsFromCSV(tempFile)
	if err != nil {
		t.Fatal(err)
	}

	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("expected %v; got %v", expected, actual)
	}
}
