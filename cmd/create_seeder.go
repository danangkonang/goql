/*
Copyright © 2022 DanangKonang danangkonang21@gmail.com
*/
package cmd

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/brianvoe/gofakeit"
	"github.com/danangkonang/goql/helper"
	"github.com/google/uuid"
	"github.com/spf13/cobra"
)

// createSeederCmd represents the createSeeder command
var field string
var count int
var createSeederCmd = &cobra.Command{
	Use:  "seeder",
	Long: "Generate seeder file",
	Run: func(cmd *cobra.Command, args []string) {
		if dirName != "" {
			dirName = fmt.Sprintf("%s/", strings.TrimRight(dirName, "/"))
		} else {
			dirName = "seeder/"
		}

		if tableName == "" {
			fmt.Println("table name can not empty")
			os.Exit(0)
		}
		names := strings.Split(tableName, ",")
		for _, v := range names {

			files, err := os.ReadDir(dirName)
			if err != nil {
				os.Mkdir(dirName, 0700)
			}

			var isDown bool
			for _, file := range files {
				filename := file.Name()
				downSplit := strings.Split(filename, ".down.")

				if len(downSplit) > 1 {
					originalUpName := strings.Split(downSplit[0], "_seeder_")
					if v == originalUpName[1] {
						isDown = true
					}
				}

			}

			query_down := ""
			query_down += fmt.Sprintf("TRUNCATE %s;", v)

			var nextName int
			nextName = len(files)

			if !isDown {
				unix_down_name := helper.CreateName(len(files))
				fmt.Println(unix_down_name)
				nextName += 1
				file_name_seeder := unix_down_name + "_seeder_" + v + ".down.sql"
				path_down_seeder := dirName + file_name_seeder
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
				query += fmt.Sprintf("INSERT INTO %s\n", v)
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
							da += fmt.Sprintf("%s,", time.Now())
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

				file_name_seeder := unix_up_name + "_seeder_" + v + ".up.sql"
				path_seeder := dirName + file_name_seeder
				file_seeder, err := os.Create(path_seeder)
				if err != nil {
					fmt.Println(err.Error())
					os.Exit(0)
				}

				file_seeder.WriteString(query)
				defer file_seeder.Close()
				fmt.Println(string(helper.GREEN), "success", string(helper.WHITE), "created", path_seeder)
			}
		}
	},
}

func init() {
	createCmd.AddCommand(createSeederCmd)
	createSeederCmd.PersistentFlags().StringVarP(&tableName, "table", "t", "", "Name of tables to generate migrations for (comma-separated for multiple tables, e.g., users,products)")
	createSeederCmd.PersistentFlags().StringVarP(&field, "field", "", "", "data type seeder, format 'colum:data type'")
	createSeederCmd.PersistentFlags().IntVarP(&count, "count", "", 1, "many seeder data will be generate")
}
