/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/pterm/pterm"
	"github.com/spf13/cobra"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all saved bookmarks in a formatted table",
	Long: `List displays all bookmarks stored in the local database.
The output is formatted as a table with ID, Name, and URL columns for easy viewing.

Example:
  bookmark list
`,
	Run: func(cmd *cobra.Command, args []string) {

		var bookmarks []Bookmark
		// fmt.Println("list called", b)

		//Open DB
		db, err := gorm.Open(sqlite.Open("bookmark.db"), &gorm.Config{})
		if err != nil {
			panic("failed to connect database")
		}

		db.AutoMigrate(&Bookmark{})

		// Get all records
		result := db.Find(&bookmarks)
		// SELECT * FROM users;

		rs := result.RowsAffected // returns found records count, equals `len(users)`
		err = result.Error        // returns error

		if err != nil {
			fmt.Println("argh something wrong", rs, err)
		}

		// Creazione della tabella
		tableData := pterm.TableData{
			{"ID", "Name", "URL"}, //intestazione
		}

		for _, b := range bookmarks {
			row := []string{
				fmt.Sprintf("%d", b.ID),
				b.Name,
				b.Url,
			}
			tableData = append(tableData, row)
		}

		alternateStyle := pterm.NewStyle(pterm.BgDarkGray)
		// pterm.DefaultTable.WithHasHeader().WithData(tableData).Render()
		// pterm.DefaultTable.WithHasHeader().WithRowSeparator("-").WithHeaderRowSeparator("-").WithData(tableData).Render()
		pterm.DefaultTable.WithHasHeader().WithBoxed().WithData(tableData).WithAlternateRowStyle(alternateStyle).Render()

	},
}

func init() {
	rootCmd.AddCommand(listCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// listCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// listCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
