/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// upSeederCmd represents the upSeeder command
var upSeederCmd = &cobra.Command{
	Use:   "seeder",
	Short: "A brief description of your command",
	Long:  "A brief description of your command",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("upSeeder called")
	},
}

func init() {
	upCmd.AddCommand(upSeederCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// upSeederCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// upSeederCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
