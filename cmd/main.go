package main

import (
	"github.com/spf13/cobra"
	"golang-restful-api-framework/internal/router"
	"golang-restful-api-framework/internal/setup"
	"strconv"
)

func InitCofig() string {
	var fileName string
	var rootCmd = &cobra.Command{
		Use: "CMD",
		Run: func(cmd *cobra.Command, args []string) {
		},
	}
	rootCmd.Flags().StringVarP(
		&fileName,
		"file",
		"f",
		"",
		"configuration file")
	rootCmd.Execute()
	return fileName
}

func main() {

	fileName := InitCofig()
	i := setup.Init(fileName)
	router.LoadUser(i)
	router.LoadTask(i)
	i.Router.Engine.Run(":" + strconv.Itoa(i.Config.App.Port))
}
