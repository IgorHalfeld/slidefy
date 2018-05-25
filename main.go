package main

import (
	"fmt"

	"github.com/slidefy/handlers"
	"github.com/slidefy/helpers"
)

func main() {
	p, err := handlers.LoadJSONFile("./example/presentation.json")
	helpers.ErrorHandler(err, "Something happed on create json map")

	for _, obj := range p {
		output := "./tmp/slide" + string(obj.Title) + ".png"
		fmt.Println(obj.Title)
		handlers.DrawTitleAndDesc(
			"./assets/images/bg.png",
			"./assets/fonts/BreeSerif-Regular.ttf",
			obj.Title,
			obj.Desc,
			output)
	}
}
