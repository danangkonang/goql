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

// var upTableName string

// upMigrationCmd represents the upMigration command
var upMigrationCmd = &cobra.Command{
	Use:   "migration",
	Short: "A brief description of your command",
	Long:  "A brief description of your command",
	Run: func(cmd *cobra.Command, args []string) {

		if tableName != "" {
			tbls := strings.Split(tableName, " ")
			for _, n := range tbls {
				query, e := os.ReadFile(fmt.Sprintf("db/migration/%s", n))
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
				msg := fmt.Sprintf("%s success %s up %s", string(helper.GREEN), string(helper.WHITE), n)
				fmt.Println(msg)
			}
		}

		if tableName == "" {
			files, err := ioutil.ReadDir("db/migration")
			if err != nil {
				fmt.Println(err.Error())
				os.Exit(0)
			}
			for _, file := range files {
				query, e := os.ReadFile(fmt.Sprintf("db/migration/%s", file.Name()))
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
				msg := fmt.Sprintf("%s success %s up %s", string(helper.GREEN), string(helper.WHITE), file.Name())
				fmt.Println(msg)
			}
		}
	},
}

func init() {
	upCmd.AddCommand(upMigrationCmd)
	upMigrationCmd.PersistentFlags().StringVarP(&tableName, "table", "t", "", "A File name to unzip and open in IDE")
	// createCmd.PersistentFlags().StringVarP(&upTableName, "table", "t", "", "A File name to unzip and open in IDE")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// upMigrationCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// upMigrationCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
