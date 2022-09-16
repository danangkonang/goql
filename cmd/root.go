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
	Version: "0.1.5",
	Short:   "Simple auto migrate database",
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
	rootCmd.PersistentFlags().StringVarP(&dirName, "dir", "", "", "Directory location migration and seeder")
}
