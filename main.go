package main

import (
	"fmt"

	"github.com/jung-kurt/gofpdf"
	"github.com/slidefy/handlers"
	"github.com/slidefy/helpers"
)

var pageSize gofpdf.InitType

func main() {
	pageSize.OrientationStr = "Landscape"
	pageSize.UnitStr = "cm"
	pageSize.FontDirStr = ""
	pageSize.Size = gofpdf.SizeType{
		Wd: 9.14,
		Ht: 16.26,
	}

	pdf := gofpdf.NewCustom(&pageSize)
	p, err := handlers.LoadJSONFile("./example/presentation.json")
	helpers.ErrorHandler(err, "Something happed on create json map")

	for _, obj := range p {
		output := "./tmp/slide" + string(obj.Title) + ".png"
		fmt.Println("Generating:", obj.Title)
		handlers.DrawTitleAndDesc(
			"./assets/images/bg.png",
			"./assets/fonts/BreeSerif-Regular.ttf",
			obj.Title,
			obj.Desc,
			output)

		pdf.AddPage()
		pdf.Image(output, 0, 0, 16.26, 9.14, false, "", 0, "")
	}
	err = pdf.OutputFileAndClose("./example/my-presentation.pdf")
	helpers.ErrorHandler(err, "Something happed on create pdf file")

}
