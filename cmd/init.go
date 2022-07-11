/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var (
	driver   string
	host     string
	port     string
	database string
	user     string
	password string
)

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Generate structure directory migration",
	Long:  "Generate structure directory migration",
	Run: func(cmd *cobra.Command, args []string) {
		if _, err := os.Stat("db"); os.IsNotExist(err) {
			os.Mkdir("db", 0700)
			os.Mkdir("db/migration", 0700)
			os.Mkdir("db/seeder", 0700)
			fmt.Println("create directory" + "db")
		}
		if _, err := os.Stat("db/migration"); os.IsNotExist(err) {
			os.Mkdir("db/migration", 0700)
			fmt.Println("create directory" + "db/migration")
		}
		if _, err := os.Stat("db/seeder"); os.IsNotExist(err) {
			os.Mkdir("db/seeder", 0700)
			fmt.Println("create directory" + "db/seeder")
		}
		if _, err := os.Stat(".env"); os.IsNotExist(err) {
			createEnvFile(".env")
			fmt.Println("create file" + ".env")
		}
	},
}

func init() {
	rootCmd.AddCommand(initCmd)
	initCmd.PersistentFlags().StringVarP(&driver, "driver", "", "", "databse driver")
	initCmd.PersistentFlags().StringVarP(&host, "host", "", "", "databse host")
	initCmd.PersistentFlags().StringVarP(&port, "port", "", "", "databse port")
	initCmd.PersistentFlags().StringVarP(&database, "database", "", "", "databse database")
	initCmd.PersistentFlags().StringVarP(&user, "user", "", "", "databse user")
	initCmd.PersistentFlags().StringVarP(&password, "password", "", "", "databse password")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// initCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// initCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func createEnvFile(name string) {

	var file, err = os.Create(name)
	if err != nil {
		fmt.Println(err.Error())
	}
	file_env := fmt.Sprintf(
		"DB_DRIVER=%s\nDB_HOST=%s\nDB_PORT=%s\nDB_NAME=%s\nDB_USER=%s\nDB_PASSWORD=%s\n",
		driver,
		host,
		port,
		database,
		user,
		password,
	)
	file.WriteString(file_env)
	defer file.Close()

}
