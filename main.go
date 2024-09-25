package main

import (
	"github.com/spf13/cobra"
	"movies/cmd"
)

// @title 	Movies Service API
// @version	1.0
// @description A Movies service API in Go using Gin framework

// @host 	localhost:8080
// @BasePath /api/v1
func main() {

	rootCmd := &cobra.Command{
		Use:   "go api template",
		Short: "A brief description of your application",
	}

	rootCmd.AddCommand(
		cmd.API(),
		cmd.Migrations(),
	)

	cobra.CheckErr(rootCmd.Execute())
}
