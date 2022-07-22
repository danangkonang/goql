/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

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

// migrationCmd represents the migration command
var migrationCmd = &cobra.Command{
	Use:   "migration",
	Short: "A brief description of your command",
	Long:  "A brief description of your command",
	Run: func(cmd *cobra.Command, args []string) {
		if tableName == "" {
			fmt.Println("table name can not empty")
			os.Exit(0)
		}
		files, err := ioutil.ReadDir("db/migration")
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
		unix_name := helper.CreateName(len(files) / 2)
		file_name_down := unix_name + "_migration_" + tableName + ".down.sql"
		path_down := "db/migration/" + file_name_down
		file_down, err := os.Create(path_down)
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(0)
		}

		file_name_up := unix_name + "_migration_" + tableName + ".up.sql"
		path_up := "db/migration/" + file_name_up
		file_up, err := os.Create(path_up)
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(0)
		}

		var query_up string
		var query_down string
		switch os.Getenv("DB_DRIVER") {
		case "postgres":
			query_up += fmt.Sprintf("CREATE TABLE IF NOT EXISTS %s(\n", tableName)
			query_up += "\tid SERIAL,\n"
			query_up += "\tname VARCHAR (225) NOT NULL,\n"
			query_up += "\tcreated_at INTEGER NOT NULL,\n"
			query_up += "\tupdated_at INTEGER NULL,\n"
			query_up += fmt.Sprintf("\tCONSTRAINT %s_pkey PRIMARY KEY (id)\n", tableName)
			query_up += ");\n"
			query_down += fmt.Sprintf("DROP TABLE IF EXISTS %s;", tableName)
		case "mysql":
			query_up += fmt.Sprintf("CREATE TABLE IF NOT EXISTS %s(\n", tableName)
			query_up += "\tid INT NOT NULL AUTO_INCREMENT PRIMARY KEY,\n"
			query_up += "\tname VARCHAR (225) NOT NULL,\n"
			query_up += "\tcreated_at INTEGER NOT NULL,\n"
			query_up += "\tupdated_at INTEGER NULL\n"
			query_up += ");\n"
			query_down += fmt.Sprintf("DROP TABLE IF EXISTS %s;", tableName)
		}

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

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// migrationCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// migrationCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
