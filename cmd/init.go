/*
Copyright © 2022 DanangKonang danangkonang21@gmail.com

*/
package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

// var (
// 	driver   string
// 	host     string
// 	port     string
// 	database string
// 	user     string
// 	password string
// )

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Generate structure directory migration",
	Long:  "Generate structure directory migration",
	Run: func(cmd *cobra.Command, args []string) {
		if dirName != "" {
			dirName = fmt.Sprintf("%s/", strings.TrimRight(dirName, "/"))
		}
		if _, err := os.Stat(fmt.Sprintf("%sdb", dirName)); os.IsNotExist(err) {
			os.Mkdir(fmt.Sprintf("%sdb", dirName), 0700)
			os.Mkdir(fmt.Sprintf("%sdb/migration", dirName), 0700)
			os.Mkdir(fmt.Sprintf("%sdb/seeder", dirName), 0700)
			msg := fmt.Sprintf("create directory %sdb", dirName)
			fmt.Println(msg)
		}
		if _, err := os.Stat(fmt.Sprintf("%sdb/migration", dirName)); os.IsNotExist(err) {
			os.Mkdir(fmt.Sprintf("%sdb/migration", dirName), 0700)
			fmt.Println("create directory" + fmt.Sprintf("%sdb/migration", dirName))
		}
		if _, err := os.Stat(fmt.Sprintf("%sdb/seeder", dirName)); os.IsNotExist(err) {
			os.Mkdir(fmt.Sprintf("%sdb/seeder", dirName), 0700)
			fmt.Println("create directory" + fmt.Sprintf("%sdb/seeder", dirName))
		}
		// if _, err := os.Stat(fmt.Sprintf("%s.env", dirName)); os.IsNotExist(err) {
		// 	createEnvFile(fmt.Sprintf("%s.env", dirName))
		// 	fmt.Println("create file" + fmt.Sprintf("%s.env", dirName))
		// }
	},
}

func init() {
	rootCmd.AddCommand(initCmd)
	// initCmd.PersistentFlags().StringVarP(&driver, "driver", "", "", "databse driver")
	// initCmd.PersistentFlags().StringVarP(&host, "host", "", "", "databse host")
	// initCmd.PersistentFlags().StringVarP(&port, "port", "", "", "databse port")
	// initCmd.PersistentFlags().StringVarP(&database, "database", "", "", "databse database")
	// initCmd.PersistentFlags().StringVarP(&user, "user", "", "", "databse user")
	// initCmd.PersistentFlags().StringVarP(&password, "password", "", "", "databse password")
}

// func createEnvFile(name string) {
// 	var file, err = os.Create(name)
// 	if err != nil {
// 		fmt.Println(err.Error())
// 	}
// 	if driver == "" {
// 		driver = "mysql"
// 	}
// 	file_env := fmt.Sprintf(
// 		"DB_DRIVER=%s\nDB_HOST=%s\nDB_PORT=%s\nDB_NAME=%s\nDB_USER=%s\nDB_PASSWORD=%s\n",
// 		driver,
// 		host,
// 		port,
// 		database,
// 		user,
// 		password,
// 	)
// 	file.WriteString(file_env)
// 	defer file.Close()
// }
