/*
Copyright © 2022 DanangKonang danangkonang21@gmail.com

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

		var isDown bool

		for _, file := range files {
			filename := file.Name()
			downSplit := strings.Split(filename, ".down.")

			if len(downSplit) > 1 {
				originalUpName := strings.Split(downSplit[0], "_seeder_")
				if tableName == originalUpName[1] {
					isDown = true
				}
			}

		}
		query_down := ""
		// switch os.Getenv("DB_DRIVER") {
		// case "postgres":
		query_down += fmt.Sprintf("TRUNCATE %s;", tableName)
		// case "mysql":
		// 	query_down += fmt.Sprintf("TRUNCATE %s;", tableName)
		// }

		var nextName int
		nextName = len(files)

		if !isDown {
			nextName += 1
			unix_down_name := helper.CreateName(len(files))
			file_name_seeder := unix_down_name + "_seeder_" + tableName + ".down.sql"
			path_down_seeder := "db/seeder/" + file_name_seeder
			file_down_seeder, err := os.Create(path_down_seeder)
			if err != nil {
				fmt.Println(err.Error())
				os.Exit(0)
			}
			file_down_seeder.WriteString(query_down)
			defer file_down_seeder.Close()
			fmt.Println(string(helper.GREEN), "success", string(helper.WHITE), "created", path_down_seeder)
		}

		column := ""
		datatype := []string{}

		if field != "" {
			for _, v := range strings.Split(field, ",") {
				fl := strings.Split(v, ":")
				column += fmt.Sprintf("%s,", fl[0])
				datatype = append(datatype, fl[1])
			}
		}

		var loopCount []int
		if count > 1000 {
			for i := 0; i < (count / 1000); i++ {
				loopCount = append(loopCount, 1000)
			}
			sisabagi := count % 1000
			if sisabagi > 0 {
				loopCount = append(loopCount, sisabagi)
			}
		} else {
			loopCount = append(loopCount, count)
		}
		for j, many := range loopCount {
			query := ""
			query += fmt.Sprintf("INSERT INTO %s\n", tableName)
			query += fmt.Sprintf("\t(%s)\n", strings.TrimSuffix(column, ","))
			query += "VALUES\n"
			for i := 0; i < many; i++ {
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

			unix_up_name := helper.CreateName(nextName + j)

			file_name_seeder := unix_up_name + "_seeder_" + tableName + ".up.sql"
			path_seeder := "db/seeder/" + file_name_seeder
			file_seeder, err := os.Create(path_seeder)
			if err != nil {
				fmt.Println(err.Error())
				os.Exit(0)
			}

			file_seeder.WriteString(query)
			defer file_seeder.Close()
			fmt.Println(string(helper.GREEN), "success", string(helper.WHITE), "created", path_seeder)

		}

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
