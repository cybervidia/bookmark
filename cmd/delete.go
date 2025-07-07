/*
Copyright © 2025 maKs <eliteknow@youknowwhere.to>
*/
package cmd

import (
	"fmt"
	"log"

	"github.com/spf13/cobra"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// deleteCmd represents the delete command
var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete a bookmark by its name",
	Long: `Delete removes the bookmark identified by the specified name from the local database.
You must provide the bookmark name to delete it.

Example:
  bookmark delete google
`,
	Run: func(cmd *cobra.Command, args []string) {

		if len(args) != 0 {
			fmt.Println("lenght args:", len(args))

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

			bookmark := Bookmark{}
			db.First(&bookmark, "name = ?", args[0]) // carica il record
			db.Unscoped().Delete(&Bookmark{}, bookmark.ID)

			fmt.Println("bookmark deleted:", bookmark.Name, bookmark.Url)

		} else {
			fmt.Println("you need to specify something to delete, \n for exemple:\nbookmark delete aRecord")
		}

	},
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
	rootCmd.AddCommand(deleteCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// deleteCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// deleteCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
