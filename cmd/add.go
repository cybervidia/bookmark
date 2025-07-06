/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"log"
	"strings"

	"github.com/atotto/clipboard"
	"github.com/spf13/cobra"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a new bookmark by specifying a name and URL",
	Long: `Add saves a new bookmark with a unique name and its URL into the local database.
You can either specify the URL directly as an argument or use the '-c' flag
to take the URL from the system clipboard.

Examples:
  bookmark add google https://google.com
  bookmark add -c favoriteSite
`,
	Run: func(cmd *cobra.Command, args []string) {

		var name, url string

		clip, _ := cmd.Flags().GetBool("clip")
		if clip {

			var err error = nil
			url, err = clipboard.ReadAll()
			if err != nil {
				log.Fatal(err)
			}
			name = args[0]

		} else {
			if len(args) == 2 {
				name = args[0]
				url = args[1]
			} else {
				fmt.Println("4 - ToDo metti descrizione di come funziona il comando ", args)
			}
		}

		db, err := gorm.Open(sqlite.Open("bookmark.db"), &gorm.Config{
			Logger: logger.Default.LogMode(logger.Silent),
		})
		if err != nil {
			panic("failed to connect database")
		}

		db.AutoMigrate(&Bookmark{})

		bkmrk := Bookmark{Name: name, Url: url}

		//db.Create(&bkmrk)
		result := db.Create(&bkmrk)

		if result.Error != nil {
			if strings.Contains(result.Error.Error(), "UNIQUE constraint failed") {
				fmt.Println("Errore: nome duplicato")
			} else {
				fmt.Println("Errore generico:", result.Error)
			}
			return // o os.Exit(1)
		}

		fmt.Println("Bookmark <", bkmrk.Name, "> inserito con successo")
	},
}

func init() {
	rootCmd.AddCommand(addCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// addCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	addCmd.Flags().BoolP("clip", "c", false, "use the clipboard as <url> parameter")
}
