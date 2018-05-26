package handlers

import (
	"fmt"

	"github.com/halfeld/slidefy/helpers"
	"github.com/halfeld/slidefy/models"
	"github.com/jung-kurt/gofpdf"
)

var pageSize gofpdf.InitType

// CreatePdfFile with images
func CreatePdfFile(options models.CLIOptions) {
	pageSize.OrientationStr = "Landscape"
	pageSize.UnitStr = "cm"
	pageSize.FontDirStr = ""
	pageSize.Size = gofpdf.SizeType{
		Wd: 9.14,
		Ht: 16.26,
	}

	pdf := gofpdf.NewCustom(&pageSize)
	p, err := LoadJSONFile(options.JSONFile)
	helpers.ErrorHandler(err, "Something happed on create json map")

	var bg string
	if options.BgFile == "" {
		bg = "../assets/images/bg.png"
	} else {
		bg = options.BgFile
	}

	for _, obj := range p {
		output := "./slide" + string(obj.Title) + ".png"
		fmt.Println("Generating:", obj.Title)
		DrawTitleAndDesc(
			bg,
			"../assets/fonts/BreeSerif-Regular.ttf",
			obj.Title,
			obj.Desc,
			output)

		pdf.AddPage()
		pdf.Image(output, 0, 0, 16.26, 9.14, false, "", 0, "")
	}
	err = pdf.OutputFileAndClose(options.PDFFile)
	helpers.ErrorHandler(err, "Something happed on create pdf file")
	fmt.Println("")
	fmt.Println("Presentation created successfully.")
}
