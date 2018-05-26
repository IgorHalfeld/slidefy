package main

import (
	"github.com/slidefy/handlers"
	"github.com/slidefy/helpers"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "slidefy",
	Short: "The way to create slides.",
	Long:  `The way to create slides when you're in a hurry, but you want them to look good.`,
	Run: func(cmd *cobra.Command, args []string) {
		JSONFile := cmd.Flag("json").Value.String()
		handlers.CreatePdfFile(JSONFile)
	},
}

func init() {
	rootCmd.PersistentFlags().String("json", "./file.json", "Path to json presentation file")
}

func main() {
	err := rootCmd.Execute()
	helpers.ErrorHandler(err, "Something happen with execution")
}
