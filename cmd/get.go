/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

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

		//Open DB
		db, err := gorm.Open(sqlite.Open("bookmark.db"), &gorm.Config{})
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
