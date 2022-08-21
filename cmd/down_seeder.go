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

// downSeederCmd represents the downSeeder command
var downSeederCmd = &cobra.Command{
	Use:   "seeder",
	Short: "A brief description of your command",
	Long:  "A brief description of your command",
	Run: func(cmd *cobra.Command, args []string) {
		files, err := ioutil.ReadDir("db/seeder")
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(0)
		}
		upFileName := []string{}

		if tableName != "" {
			for _, file := range files {
				tbls := strings.Split(tableName, " ")
				// filter only up file
				fl_up := strings.Split(file.Name(), ".down.")
				// find original table name
				original := strings.Split(fl_up[0], "_seeder_")[1]
				for _, g := range tbls {
					if g == original {
						upFileName = append(upFileName, file.Name())
					}
				}
			}
		}

		if tableName == "" {
			for _, file := range files {
				// filter only up file
				if len(strings.Split(file.Name(), ".down.")) > 1 {
					upFileName = append(upFileName, file.Name())
				}
			}

		}

		for _, fil := range upFileName {
			query, e := os.ReadFile(fmt.Sprintf("db/seeder/%s", fil))
			if e != nil {
				fmt.Println(e.Error())
				os.Exit(0)
			}
			ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
			defer cancel()
			conn := config.Connection(dbConnection)
			_, err := conn.DB.ExecContext(ctx, string(query))
			if err != nil {
				fmt.Println(err.Error())
				os.Exit(0)
			}
			msg := fmt.Sprintf("%s success %sdown %s", string(helper.GREEN), string(helper.WHITE), fil)
			fmt.Println(msg)
		}
	},
}

func init() {
	downCmd.AddCommand(downSeederCmd)
	downCmd.PersistentFlags().StringVarP(&tableName, "table", "t", "", "Table name")

	downCmd.PersistentFlags().StringVarP(&dbConnection, "db", "", "", "Database connection")
	downCmd.MarkFlagRequired("db")
}
