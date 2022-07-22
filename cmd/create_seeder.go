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

	"github.com/brianvoe/gofakeit/v6"
	"github.com/danangkonang/goql/helper"
	"github.com/google/uuid"
	"github.com/spf13/cobra"
)

// createSeederCmd represents the createSeeder command
var field string
var count int
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
		// fmt.Println(files)
		// fmt.Println(strings.Split(field, ","))

		query := ""
		column := ""
		datatype := []string{}
		// fmt.Println(strings.Split(field, ","))
		// fmt.Println(query)
		// fmt.Println(count)

		// if field != "" {
		for _, v := range strings.Split(field, ",") {
			fl := strings.Split(v, ":")
			column += fmt.Sprintf("%s,", fl[0])
			datatype = append(datatype, fl[1])
		}
		// }
		// fmt.Println(datatype)
		// os.Exit(0)
		query += fmt.Sprintf("INSERT INTO %s\n", tableName)
		query += fmt.Sprintf("\t(%s)\n", strings.TrimSuffix(column, ","))
		query += "VALUES\n"

		for i := 0; i < count; i++ {
			var da string
			for _, tp := range datatype {
				switch tp {
				case "uuid":
					id := uuid.New().String()
					da += fmt.Sprintf("'%s',", id)
				case "email":
					da += fmt.Sprintf("'%s',", gofakeit.Email())
				case "name":
					da += fmt.Sprintf("'%s',", gofakeit.Name())
				case "phone":
					da += fmt.Sprintf("'%s',", gofakeit.Phone())
				case "color":
					da += fmt.Sprintf("'%s',", gofakeit.Color())
				case "gender":
					da += fmt.Sprintf("'%s',", gofakeit.Gender())
				case "hobby":
					da += fmt.Sprintf("'%s',", gofakeit.Hobby())
				case "street":
					da += fmt.Sprintf("'%s',", gofakeit.Street())
				case "city":
					da += fmt.Sprintf("'%s',", gofakeit.City())
				case "country":
					da += fmt.Sprintf("'%s',", gofakeit.Country())
				case "lat":
					da += fmt.Sprintf("%f,", gofakeit.Latitude())
				case "lng":
					da += fmt.Sprintf("%f,", gofakeit.Longitude())
				case "time":
					da += fmt.Sprintf("%d,", time.Now())
				case "unixtime":
					da += fmt.Sprintf("%d,", time.Now().Unix())
				}
			}
			if i == count-1 {
				query += fmt.Sprintf("\t(%s);\n", strings.TrimSuffix(da, ","))
			} else {
				query += fmt.Sprintf("\t(%s),\n", strings.TrimSuffix(da, ","))
			}
		}

		query_down := ""
		switch os.Getenv("DB_DRIVER") {
		case "postgres":
			query_down += fmt.Sprintf("TRUNCATE %s;", tableName)
		case "mysql":
			query_down += fmt.Sprintf("TRUNCATE %s;", tableName)
		}

		unix_name := helper.CreateName(len(files) / 2)

		file_name_seeder_down := unix_name + "_seeder_" + tableName + ".down.sql"
		path_seeder_down := "db/seeder/" + file_name_seeder_down
		file_seeder_down, err := os.Create(path_seeder_down)
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(0)
		}

		file_name_seeder := unix_name + "_seeder_" + tableName + ".up.sql"
		path_seeder := "db/seeder/" + file_name_seeder
		file_seeder, err := os.Create(path_seeder)
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(0)
		}

		// var query string
		// switch os.Getenv("DB_DRIVER") {
		// case "postgres":
		// 	fmt.Println(field)
		// 	// query += fmt.Sprintf("INSERT INTO %s\n", tableName)
		// 	// query += "\t(id, name, created_at, updated_at)\n"
		// 	// query += "VALUES\n"
		// 	// query += fmt.Sprintf("\t(1, 'bob', %d, %d);\n", time.Now().Unix(), 0)
		// case "mysql":
		// 	query += ");\n"
		// }

		file_seeder_down.WriteString(query_down)
		defer file_seeder_down.Close()

		file_seeder.WriteString(query)
		defer file_seeder.Close()
		fmt.Println(string(helper.GREEN), "success", string(helper.WHITE), "created", path_seeder_down)
		fmt.Println(string(helper.GREEN), "success", string(helper.WHITE), "created", path_seeder)
	},
}

func init() {
	createCmd.AddCommand(createSeederCmd)
	createSeederCmd.PersistentFlags().StringVarP(&tableName, "table", "t", "", "Name of table")
	createSeederCmd.PersistentFlags().StringVarP(&field, "field", "", "", "Data tipe seeder, format 'colum:data type'")
	createSeederCmd.PersistentFlags().IntVarP(&count, "count", "", 1, "Many seeder data will be generate")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// createSeederCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// createSeederCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
