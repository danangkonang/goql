/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"context"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"time"

	"github.com/danangkonang/goql/config"
	"github.com/danangkonang/goql/helper"
	"github.com/spf13/cobra"
)

// downMigrationCmd represents the downMigration command
var downMigrationCmd = &cobra.Command{
	Use:   "migration",
	Short: "A brief description of your command",
	Long:  "A brief description of your command",
	Run: func(cmd *cobra.Command, args []string) {
		files, err := ioutil.ReadDir("db/migration")
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(0)
		}
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		for _, file := range files {
			filename := file.Name()
			rmExtension := strings.Split(filename, ".")
			rmMigration := strings.Split(rmExtension[0], "_migration_")
			originalname := rmMigration[1]
			var query string
			switch "psql" {
			case "mysql":
				query = fmt.Sprintf("TRUNCATE %s;", originalname)
			case "psql":
				query = fmt.Sprintf("TRUNCATE  %s RESTART IDENTITY;", originalname)
			}
			// fmt.Println("QUERY", query)
			conn := config.Connection()
			conn.DB.ExecContext(ctx, query)
			msg := fmt.Sprintf("%s success %s down %s", string(helper.GREEN), string(helper.WHITE), file.Name())
			fmt.Println(msg)
		}
	},
}

func init() {
	downCmd.AddCommand(downMigrationCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// downMigrationCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// downMigrationCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
