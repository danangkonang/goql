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
			switch os.Getenv("DB_DRIVER") {
			case "mysql":
				// query = fmt.Sprintf("TRUNCATE %s;", originalname)
				query = fmt.Sprintf("DROP TABLE %s;", originalname)
			case "postgres":
				// query = fmt.Sprintf("TRUNCATE %s RESTART IDENTITY;", originalname)
				// IF EXISTS
				query = fmt.Sprintf("DROP TABLE %s;", originalname)
			}
			conn := config.Connection()
			_, err := conn.DB.ExecContext(ctx, query)
			if err != nil {
				fmt.Println(err.Error())
				os.Exit(0)
			}
			// fmt.Println(query)
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
