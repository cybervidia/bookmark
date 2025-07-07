/*
Copyright © 2025 maKs <eliteknow@youknowwhere.to>
*/
package cmd

import (
	"fmt"
	"log"

	"github.com/atotto/clipboard"
	"github.com/spf13/cobra"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// getCmd represents the get command
var getCmd = &cobra.Command{
	Use:   "get",
	Short: "Copy the URL of a bookmark to the clipboard by name",
	Long: `Get retrieves the URL associated with the given bookmark name and copies it
to the system clipboard for easy pasting elsewhere.

Example:
  bookmark get google
`,
	Run: func(cmd *cobra.Command, args []string) {

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

		err = clipboard.WriteAll(bookmark.Url)
		if err != nil {
			fmt.Println("Errore nel copiare nella clipboard:", err)
			return
		}
		fmt.Println("Testo copiato nella clipboard:", bookmark.Url)
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
	rootCmd.AddCommand(getCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// getCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	getCmd.Flags().BoolP("id", "i", false, "Help message for toggle")
}
