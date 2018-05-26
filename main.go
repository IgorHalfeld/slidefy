package main

import (
	"github.com/halfeld/slidefy/handlers"
	"github.com/halfeld/slidefy/helpers"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "slidefy",
	Short: "The way to create slides.",
	Long:  `The way to create slides when you're in a hurry, but you want them to look good.`,
	Run: func(cmd *cobra.Command, args []string) {
		JSONFile := cmd.Flag("json").Value.String()
		PDFFile := cmd.Flag("pdfoutput").Value.String()
		bgFile := cmd.Flag("bg").Value.String()
		handlers.CreatePdfFile(JSONFile, PDFFile, bgFile)
	},
}

func init() {
	rootCmd.PersistentFlags().String("json", "./file.json", "Path to json presentation file")
	rootCmd.PersistentFlags().String("pdfoutput", "./file.pdf", "Path to pdf file")
	rootCmd.PersistentFlags().String("bg", "./bg.png", "Path to image file")
}

func main() {
	err := rootCmd.Execute()
	helpers.ErrorHandler(err, "Something happen with execution")
}
