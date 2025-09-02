/*
Copyright © 2022 DanangKonang danangkonang21@gmail.com
*/
package cmd

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/danangkonang/goql/helper"
	"github.com/spf13/cobra"
)

var tableName string
var dirName string

// migrationCmd represents the migration command
var migrationCmd = &cobra.Command{
	Use:  "migration",
	Long: "Generate migration file",
	Run: func(cmd *cobra.Command, args []string) {

		if dirName != "" {
			dirName = fmt.Sprintf("%s/", strings.TrimRight(dirName, "/"))
		} else {
			dirName = "migration/"
		}
		if tableName == "" {
			fmt.Println("table name can not empty")
			os.Exit(0)
		}

		names := strings.Split(tableName, ",")
		for _, v := range names {
			files, err := os.ReadDir(dirName)
			if err != nil {
				os.MkdirAll(dirName, 0700)
			}

			maxSeq := 0
			for _, file := range files {
				filename := file.Name()
				rmExtension := strings.Split(filename, ".")
				rmMigration := strings.Split(rmExtension[0], "_migration_")
				originalname := rmMigration[1]
				if v == originalname {
					fmt.Printf("table '%s' already exists\n", v)
					os.Exit(0)
				}

				parts := strings.SplitN(file.Name(), "_", 2)
				if seq, err := strconv.Atoi(parts[0]); err == nil {
					if seq > maxSeq {
						maxSeq = seq
					}
				}
			}
			unix_name_up := helper.CreateNextName(maxSeq + 1)
			file_name_up := unix_name_up + "_migration_" + v + ".up.sql"
			path_up := dirName + file_name_up
			file_up, err := os.Create(path_up)
			if err != nil {
				fmt.Println(err.Error())
				os.Exit(0)
			}

			unix_name_down := helper.CreateNextName(maxSeq + 2)
			file_name_down := unix_name_down + "_migration_" + v + ".down.sql"
			path_down := dirName + file_name_down
			file_down, err := os.Create(path_down)
			if err != nil {
				fmt.Println(err.Error())
				os.Exit(0)
			}

			var query_up string
			var query_down string
			query_up += fmt.Sprintf("CREATE TABLE IF NOT EXISTS %s(\n", v)
			query_up += "\tid INT AUTO_INCREMENT PRIMARY KEY,\n"
			query_up += "\tname VARCHAR(225) NOT NULL,\n"
			query_up += "\tcreated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,\n"
			query_up += "\tupdated_at TIMESTAMP NULL\n"
			query_up += ");\n"
			query_down += fmt.Sprintf("DROP TABLE IF EXISTS %s CASCADE;", v)

			file_up.WriteString(query_up)
			file_down.WriteString(query_down)
			defer file_up.Close()
			defer file_down.Close()
			fmt.Println(string(helper.GREEN), "success", string(helper.WHITE), "created", path_up)
			fmt.Println(string(helper.GREEN), "success", string(helper.WHITE), "created", path_down)
		}
	},
}

func init() {
	migrationCmd.PersistentFlags().StringVarP(&tableName, "table", "t", "", "Name of tables to generate migrations for (comma-separated for multiple tables, e.g., users,products)")
	createCmd.AddCommand(migrationCmd)
}
