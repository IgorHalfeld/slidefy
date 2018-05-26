package handlers

import (
	"github.com/fogleman/gg"
	"github.com/halfeld/slidefy/helpers"
)

// DrawTitleAndDesc load title and description
func DrawTitleAndDesc(
	imagePath string,
	fontPath string,
	title string,
	desc string,
	output string) {

	const width = 1920
	const height = 1080

	image, err := gg.LoadImage(imagePath)
	helpers.ErrorHandler(err, "Pass a correct path to image")

	dc := gg.NewContext(width, height)
	dc.SetRGB255(108, 92, 231)

	err = dc.LoadFontFace(fontPath, 160)
	helpers.ErrorHandler(err, "Fail to load font to title")

	dc.DrawImage(image, 0, 0)

	dc.DrawStringAnchored(title, width/2, height/2.3, 0.5, 0.5)

	err = dc.LoadFontFace(fontPath, 67)
	helpers.ErrorHandler(err, "Fail to load font to description")

	dc.SetRGB255(99, 110, 114)
	dc.DrawStringAnchored(desc, width/2, height/1.8, 0.5, 0.5)

	dc.Clip()
	dc.SavePNG(output)
}
