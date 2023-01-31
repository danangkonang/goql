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

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Generate Structure Directory Migration",
	Long:  "Generate Structure Directory Migration",
	Run: func(cmd *cobra.Command, args []string) {
		if dirName != "" {
			dirName = fmt.Sprintf("%s/", strings.TrimRight(dirName, "/"))
		}
		if _, err := os.Stat(fmt.Sprintf("%smigration", dirName)); os.IsNotExist(err) {
			os.Mkdir(fmt.Sprintf("%smigration", dirName), 0700)
			fmt.Println("create directory " + fmt.Sprintf("%smigration", dirName))
		}
		if _, err := os.Stat(fmt.Sprintf("%sseeder", dirName)); os.IsNotExist(err) {
			os.Mkdir(fmt.Sprintf("%sseeder", dirName), 0700)
			fmt.Println("create directory " + fmt.Sprintf("%sseeder", dirName))
		}
	},
}

func init() {
	rootCmd.AddCommand(initCmd)
}
