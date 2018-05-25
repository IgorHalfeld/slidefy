package handlers

import (
	"encoding/json"
	"io/ioutil"

	"github.com/slidefy/helpers"
	"github.com/slidefy/models"
)

// LoadJSONFile load presention readmap
func LoadJSONFile(jsonPath string) ([]models.Scratch, error) {
	file, err := ioutil.ReadFile(jsonPath)
	helpers.ErrorHandler(err, "Fail to load presentation json")

	var scratch []models.Scratch
	err = json.Unmarshal(file, &scratch)
	helpers.ErrorHandler(err, "Fail to load presentation json")

	return scratch, err
}
