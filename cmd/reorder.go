/*
Copyright © 2022 DanangKonang danangkonang21@gmail.com
*/
package cmd

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"sort"
	"strings"

	"github.com/spf13/cobra"
)

var sequence int
var reorderCmd = &cobra.Command{
	Use:   "reorder",
	Short: "Reorder Migration Or Seeder",
	Long:  "Reorder Migration Or Seeder",
	Run: func(cmd *cobra.Command, args []string) {
		if dirName != "" {
			dirName = fmt.Sprintf("%s/", strings.TrimRight(dirName, "/"))
		} else {
			dirName = "migration/"
		}
		if tableName == "" {
			fmt.Println("table name can not empty")
			os.Exit(0)
		}
		files, err := os.ReadDir(dirName)
		if err != nil {
			panic(err)
		}

		// re := regexp.MustCompile(`^(\d{4})_(.+)\.sql$`)
		reFile := regexp.MustCompile(`^(\d{4})_(.+)\.sql$`)
		reMatch := regexp.MustCompile(fmt.Sprintf(`_(%s)\.`, regexp.QuoteMeta(tableName)))

		var migrationFiles []string
		for _, f := range files {
			if f.IsDir() || !strings.HasSuffix(f.Name(), ".sql") {
				continue
			}
			migrationFiles = append(migrationFiles, f.Name())
		}

		sort.Strings(migrationFiles)

		var moved, remaining []string
		for _, f := range migrationFiles {
			if reMatch.MatchString(f) {
				moved = append(moved, f)
			} else {
				remaining = append(remaining, f)
			}
		}
		if sequence > len(remaining) {
			sequence = len(remaining)
		}
		finalOrder := append([]string{}, remaining[:sequence]...)
		finalOrder = append(finalOrder, moved...)
		finalOrder = append(finalOrder, remaining[sequence:]...)

		for i, name := range finalOrder {
			match := reFile.FindStringSubmatch(name)
			if len(match) < 3 {
				continue
			}
			oldPath := filepath.Join(dirName, name)
			newName := fmt.Sprintf("%04d_%s.sql", i+1, match[2])
			newPath := filepath.Join(dirName, newName)

			if name != newName {
				fmt.Printf("Rename: %s -> %s\n", name, newName)
				if err := os.Rename(oldPath, newPath); err != nil {
					panic(err)
				}
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(reorderCmd)
	reorderCmd.PersistentFlags().StringVarP(&tableName, "table", "t", "", "Name of tables")
	reorderCmd.PersistentFlags().IntVarP(&sequence, "seq", "", 1, "sequence file")
}
