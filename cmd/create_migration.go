/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"errors"
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
	Args: func(cmd *cobra.Command, args []string) error {
		if tableName == "" && len(args) < 1 {
			fmt.Println(args)
			return errors.New("accepts 1 arg(s)")
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		if tableName == "" {
			fmt.Println("err.Error()")
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
				fmt.Println("table already exists")
				os.Exit(0)
			}
		}
		file_name := helper.CreateName(len(files)) + "_migration_" + tableName + ".sql"
		path := "db/migration/" + file_name
		file, err := os.Create(path)
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(0)
		}
		var query string
		switch os.Getenv("DB_DRIVER") {
		case "postgres":
			query += fmt.Sprintf("CREATE TABLE IF NOT EXISTS %s(\n", tableName)
			query += "\tid SERIAL,\n"
			query += "\tname VARCHAR (225) NOT NULL,\n"
			query += "\tcreated_at INTEGER NOT NULL,\n"
			query += "\tupdated_at INTEGER NULL,\n"
			query += fmt.Sprintf("\tCONSTRAINT %s_pkey PRIMARY KEY (id)\n", tableName)
			query += ");\n"
		case "mysql":
			query += fmt.Sprintf("CREATE TABLE IF NOT EXISTS %s(\n", tableName)
			query += "\tid INT NOT NULL AUTO_INCREMENT PRIMARY KEY,\n"
			query += "\tname VARCHAR (225) NOT NULL,\n"
			query += "\tcreated_at INTEGER NOT NULL,\n"
			query += "\tupdated_at INTEGER NULL\n"
			query += ");\n"
		}

		file.WriteString(query)
		defer file.Close()
		fmt.Println(string(helper.GREEN), "success", string(helper.WHITE), "created", path)
	},
}

func init() {
	createCmd.AddCommand(migrationCmd)
	createCmd.PersistentFlags().StringVarP(&tableName, "table", "t", "", "A File name to unzip and open in IDE")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// migrationCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// migrationCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
