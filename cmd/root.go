/*
Copyright © 2022 DanangKonang danangkonang21@gmail.com

*/
package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:     "goql",
	Version: "0.2.0",
	Long:    "Simple auto migrate database",
	CompletionOptions: cobra.CompletionOptions{
		DisableDefaultCmd: true,
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.SetHelpCommand(&cobra.Command{
		Hidden: true,
	})
	rootCmd.PersistentFlags().StringVarP(&dirName, "dir", "", "", "directory location migration and seeder")
}
