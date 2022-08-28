/*
Copyright © 2022 DanangKonang danangkonang21@gmail.com

*/
package cmd

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/danangkonang/goql/helper"
	"github.com/spf13/cobra"
)

var tableName string
var dirName string

// migrationCmd represents the migration command
var migrationCmd = &cobra.Command{
	Use:   "migration",
	Short: "A brief description of your command",
	Long:  "A brief description of your command",
	Run: func(cmd *cobra.Command, args []string) {
		if dirName != "" {
			dirName = fmt.Sprintf("%s/", strings.TrimRight(dirName, "/"))
		}

		if tableName == "" {
			fmt.Println("table name can not empty")
			os.Exit(0)
		}

		files, err := ioutil.ReadDir(dirName + "migration")
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(0)
		}
		for _, file := range files {
			filename := file.Name()
			rmExtension := strings.Split(filename, ".")
			rmMigration := strings.Split(rmExtension[0], "_migration_")
			originalname := rmMigration[1]
			if tableName == originalname {
				fmt.Printf("table '%s' already exists", tableName)
				os.Exit(0)
			}
		}
		unix_name_down := helper.CreateName(len(files) + 1)
		file_name_down := unix_name_down + "_migration_" + tableName + ".down.sql"
		path_down := "migration/" + file_name_down
		file_down, err := os.Create(path_down)
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(0)
		}

		unix_name_up := helper.CreateName(len(files))
		file_name_up := unix_name_up + "_migration_" + tableName + ".up.sql"
		path_up := "migration/" + file_name_up
		file_up, err := os.Create(path_up)
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(0)
		}

		var query_up string
		var query_down string
		query_up += fmt.Sprintf("CREATE TABLE IF NOT EXISTS %s(\n", tableName)
		query_up += "\tname VARCHAR (225) NOT NULL,\n"
		query_up += "\tcreated_at INTEGER NOT NULL,\n"
		query_up += "\tupdated_at INTEGER NULL\n"
		query_up += ");\n"
		query_down += fmt.Sprintf("DROP TABLE IF EXISTS %s CASCADE;", tableName)

		file_up.WriteString(query_up)
		file_down.WriteString(query_down)
		defer file_up.Close()
		defer file_down.Close()
		fmt.Println(string(helper.GREEN), "success", string(helper.WHITE), "created", path_down)
		fmt.Println(string(helper.GREEN), "success", string(helper.WHITE), "created", path_up)
	},
}

func init() {
	createCmd.AddCommand(migrationCmd)
	migrationCmd.PersistentFlags().StringVarP(&tableName, "table", "t", "", "Table name")
}
