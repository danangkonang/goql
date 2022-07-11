/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"context"
	"fmt"
	"io/ioutil"
	"os"
	"time"

	"github.com/danangkonang/goql/config"
	"github.com/spf13/cobra"
)

// upMigrationCmd represents the upMigration command
var upMigrationCmd = &cobra.Command{
	Use:   "migration",
	Short: "A brief description of your command",
	Long:  "A brief description of your command",
	Run: func(cmd *cobra.Command, args []string) {
		files, err := ioutil.ReadDir("db/migration")
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(0)
		}
		for _, file := range files {
			dat, e := os.ReadFile(fmt.Sprintf("db/migration/%s", file.Name()))
			if e != nil {
				fmt.Println(e)
				os.Exit(0)
			}
			ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
			defer cancel()
			conn := config.Connection()
			_, err := conn.DB.ExecContext(ctx, string(dat))
			if err != nil {
				fmt.Println(err)
				os.Exit(0)
			}
		}
	},
}

func init() {
	upCmd.AddCommand(upMigrationCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// upMigrationCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// upMigrationCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
