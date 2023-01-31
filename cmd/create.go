/*
Copyright © 2022 DanangKonang danangkonang21@gmail.com

*/
package cmd

import (
	"fmt"
	"os"

	"github.com/danangkonang/goql/helper"
	"github.com/spf13/cobra"
)

var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create Migration Or Seeder",
	Long:  "Create Migration Or Seeder",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 {
			cmd.Help()
			os.Exit(0)
		}
		if args[0] != "migration" && args[0] != "seeder" {
			msg := fmt.Sprintf("unknow %s%s%s command", string(helper.READ), args[0], string(helper.WHITE))
			fmt.Println(msg)
			fmt.Println(`Use "goql create [command] --help" for more information about a command.`)
			os.Exit(0)
		}
	},
}

func init() {
	rootCmd.AddCommand(createCmd)
}
