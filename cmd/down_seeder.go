/*
Copyright © 2022 DanangKonang danangkonang21@gmail.com
*/
package cmd

import (
	"context"
	"fmt"
	"os"
	"sort"
	"strings"
	"time"

	"github.com/danangkonang/goql/config"
	"github.com/danangkonang/goql/helper"
	"github.com/spf13/cobra"
)

// downSeederCmd represents the downSeeder command
var downSeederCmd = &cobra.Command{
	Use:  "seeder",
	Long: "Down seeder file",
	Run: func(cmd *cobra.Command, args []string) {
		if dirName != "" {
			dirName = fmt.Sprintf("%s/", strings.TrimRight(dirName, "/"))
		} else {
			dirName = "seeder/"
		}

		files, err := os.ReadDir(dirName)
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(0)
		}
		sort.Slice(files, func(i, j int) bool {
			return files[i].Name() > files[j].Name()
		})

		downFileName := []string{}

		if tableName != "" {
			for _, file := range files {
				tbls := strings.Split(tableName, " ")
				// filter only up file
				fl_up := strings.Split(file.Name(), ".down.")
				// find original table name
				original := strings.Split(fl_up[0], "_seeder_")[1]
				for _, g := range tbls {
					if g == original {
						downFileName = append(downFileName, file.Name())
					}
				}
			}
		}

		if tableName == "" {
			for _, file := range files {
				// filter only up file
				if len(strings.Split(file.Name(), ".down.")) > 1 {
					downFileName = append(downFileName, file.Name())
				}
			}

		}

		for _, fil := range downFileName {
			query, e := os.ReadFile(fmt.Sprintf("%s%s", dirName, fil))
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
			msg := fmt.Sprintf("%s success %s down %s", string(helper.GREEN), string(helper.WHITE), fil)
			fmt.Println(msg)
		}
	},
}

func init() {
	downCmd.AddCommand(downSeederCmd)
	downSeederCmd.PersistentFlags().StringVarP(&tableName, "table", "t", "", "Table Name")
	downSeederCmd.PersistentFlags().StringVarP(&dbConnection, "db", "", "", "database connection")
	downSeederCmd.MarkFlagRequired("db")
}
