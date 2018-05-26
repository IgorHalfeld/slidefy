package handlers

import (
	"fmt"

	"github.com/halfeld/slidefy/helpers"
	"github.com/jung-kurt/gofpdf"
)

var pageSize gofpdf.InitType

// CreatePdfFile with images
func CreatePdfFile(jsonfile string, pdfFile string) {
	pageSize.OrientationStr = "Landscape"
	pageSize.UnitStr = "cm"
	pageSize.FontDirStr = ""
	pageSize.Size = gofpdf.SizeType{
		Wd: 9.14,
		Ht: 16.26,
	}

	pdf := gofpdf.NewCustom(&pageSize)
	p, err := LoadJSONFile(jsonfile)
	helpers.ErrorHandler(err, "Something happed on create json map")

	for _, obj := range p {
		output := "./slide" + string(obj.Title) + ".png"
		fmt.Println("Generating:", obj.Title)
		DrawTitleAndDesc(
			"../assets/images/bg.png",
			"../assets/fonts/BreeSerif-Regular.ttf",
			obj.Title,
			obj.Desc,
			output)

		pdf.AddPage()
		pdf.Image(output, 0, 0, 16.26, 9.14, false, "", 0, "")
	}
	err = pdf.OutputFileAndClose(pdfFile)
	helpers.ErrorHandler(err, "Something happed on create pdf file")
	fmt.Println("")
	fmt.Println("Presentation created successfully.")
}
