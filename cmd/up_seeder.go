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

// upSeederCmd represents the upSeeder command
var upSeederCmd = &cobra.Command{
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
				fl_up := strings.Split(file.Name(), ".")
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
				upFileName = append(upFileName, file.Name())
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
			conn := config.Connection()
			_, err := conn.DB.ExecContext(ctx, string(query))
			if err != nil {
				fmt.Println(err.Error())
				os.Exit(0)
			}
			msg := fmt.Sprintf("%s success %s up %s", string(helper.GREEN), string(helper.WHITE), fil)
			fmt.Println(msg)
		}

	},
}

func init() {
	upCmd.AddCommand(upSeederCmd)
	upSeederCmd.PersistentFlags().StringVarP(&tableName, "table", "t", "", "For table name")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// upSeederCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// upSeederCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}