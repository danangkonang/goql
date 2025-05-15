/*
Copyright © 2022 DanangKonang danangkonang21@gmail.com
*/
package cmd

import (
	"context"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/danangkonang/goql/config"
	"github.com/danangkonang/goql/helper"
	"github.com/spf13/cobra"
)

var dbConnection string

var upMigrationCmd = &cobra.Command{
	Use:  "migration",
	Long: "execute migration file",
	Run: func(cmd *cobra.Command, args []string) {
		if dirName != "" {
			dirName = fmt.Sprintf("%s/", strings.TrimRight(dirName, "/"))
		} else {
			dirName = "migration/"
		}

		files, err := os.ReadDir(dirName)
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(0)
		}

		upFileName := []string{}

		if tableName != "" {
			for _, file := range files {
				tbls := strings.Split(tableName, " ")
				// filter only up file
				fl_up := strings.Split(file.Name(), ".up.")
				if len(fl_up) > 1 {
					// find original table name
					original := strings.Split(fl_up[0], "_migration_")[1]
					for _, g := range tbls {
						if g == original {
							upFileName = append(upFileName, file.Name())
						}
					}
				}
			}
		}

		if tableName == "" {
			for _, file := range files {
				// filter only up file
				if len(strings.Split(file.Name(), ".up.")) > 1 {
					upFileName = append(upFileName, file.Name())
				}
			}
		}

		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		conn := config.Connection(dbConnection)
		tx, err := conn.DB.BeginTx(ctx, nil)
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(0)
		}

		for _, fil := range upFileName {
			query, e := os.ReadFile(fmt.Sprintf("%s%s", dirName, fil))
			if e != nil {
				fmt.Println(e.Error())
				os.Exit(0)
			}
			// ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
			// defer cancel()
			// conn := config.Connection(dbConnection)
			_, err := tx.ExecContext(ctx, string(query))
			if err != nil {
				tx.Rollback()
				fmt.Println(err.Error())
				os.Exit(0)
				// continue
			}
			msg := fmt.Sprintf("%s success %s up %s", string(helper.GREEN), string(helper.WHITE), fil)
			fmt.Println(msg)
		}

		tx.Commit()
	},
}

func init() {
	upCmd.AddCommand(upMigrationCmd)
	upMigrationCmd.PersistentFlags().StringVarP(&tableName, "table", "t", "", "table name (not file name)")
	upMigrationCmd.PersistentFlags().StringVarP(&dbConnection, "db", "", "", "database connection")
	upMigrationCmd.MarkFlagRequired("db")
}
