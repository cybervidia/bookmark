/*
Copyright © 2025 maKs <eliteknow@youknowwhere.to>
*/
package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/pterm/pterm"
	"github.com/spf13/cobra"
	"golang.org/x/term"
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

		dbPath, err := getDatabasePath()
		if err != nil {
			log.Fatalf("Failed to get database path: %v", err)
		}

		//Open DB
		db, err := gorm.Open(sqlite.Open(dbPath), &gorm.Config{})
		if err != nil {
			panic("failed to connect database")
		}

		db.AutoMigrate(&Bookmark{})

		// Get all records
		result := db.Find(&bookmarks)
		// SELECT * FROM users;

		// rs := result.RowsAffected // returns found records count, equals `len(users)`
		err = result.Error // returns error

		if err != nil {
			fmt.Println("argh something wrong", err)
		}

		// Creazione della tabella
		tableData := pterm.TableData{
			{
				// "ID",
				"Name",
				"URL",
			}, //intestazione
		}

		for _, b := range bookmarks {
			row := []string{
				// fmt.Sprintf("%d", b.ID),
				b.Name,
				b.Url,
			}
			tableData = append(tableData, row)
		}

		alternateStyle := pterm.NewStyle(pterm.BgDarkGray)
		pterm.DefaultTable.WithHasHeader().WithBoxed().WithData(tableData).WithAlternateRowStyle(alternateStyle).Render()

		//solo temporaneo print per debug
		getTermWidth()

	},
}

func getTermWidth() int {
	width, _, err := term.GetSize(int(os.Stdout.Fd()))
	if err != nil {
		fmt.Println("Error:", err)
		return 0
	}
	fmt.Printf("Terminal size: width=%d, height=%d\n", width)
	return width
}

// func getDatabasePath() (string, error) {
// 	exePath, err := os.Executable()
// 	if err != nil {
// 		return "", err
// 	}
// 	exeDir := filepath.Dir(exePath)
// 	dbPath := filepath.Join(exeDir, "bookmark.db")
// 	return dbPath, nil
// }

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
