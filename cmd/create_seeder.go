/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"time"

	"github.com/danangkonang/goql/helper"
	"github.com/spf13/cobra"
)

// createSeederCmd represents the createSeeder command
var createSeederCmd = &cobra.Command{
	Use:   "seeder",
	Short: "Generate seeder file",
	Long:  "Generate seeder file",
	Run: func(cmd *cobra.Command, args []string) {
		// fmt.Println("tableName", tableName)
		if tableName == "" {
			fmt.Println("table name can not empty")
			os.Exit(0)
		}
		files, err := ioutil.ReadDir("db/seeder")
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(0)
		}
		for _, file := range files {
			filename := file.Name()
			rmExtension := strings.Split(filename, ".")
			rmMigration := strings.Split(rmExtension[0], "_seeder_")
			originalname := rmMigration[1]
			if tableName == originalname {
				fmt.Println("seeder table name already exists")
				os.Exit(0)
			}
		}
		unix_name := helper.CreateName(len(files))
		file_name_seeder := unix_name + "_seeder_" + tableName + ".sql"
		path_seeder := "db/seeder/" + file_name_seeder
		file_seeder, err := os.Create(path_seeder)
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(0)
		}

		var query string
		switch os.Getenv("DB_DRIVER") {
		case "postgres":
			query += fmt.Sprintf("INSERT INTO %s\n", tableName)
			query += "\t(id, name, created_at, updated_at)\n"
			query += "VALUES\n"
			query += fmt.Sprintf("\t(1, 'bob', %d, %d);\n", time.Now().Unix(), 0)
		case "mysql":
			query += ");\n"
		}

		file_seeder.WriteString(query)
		defer file_seeder.Close()
		fmt.Println(string(helper.GREEN), "success", string(helper.WHITE), "created", path_seeder)
	},
}

func init() {
	createCmd.AddCommand(createSeederCmd)
	createSeederCmd.PersistentFlags().StringVarP(&tableName, "table", "t", "", "For table name")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// createSeederCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// createSeederCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
