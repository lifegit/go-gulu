package main

import (
	"github.com/lifegit/go-gulu/v2/nice/statisticsLine/app"
	"github.com/spf13/cobra"
	"log"
)

//go:generate go build -o main
func main() {
	var rootCmd = &cobra.Command{Use: "statisticsLine"}
	rootCmd.AddCommand(app.RestCmd)

	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
