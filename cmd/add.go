/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
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

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
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

		db, err := gorm.Open(sqlite.Open("bookmark.db"), &gorm.Config{})
		if err != nil {
			panic("failed to connect database")
		}

		db.AutoMigrate(&Bookmark{})

		bkmrk := Bookmark{Name: name, Url: url}

		db.Create(&bkmrk)
		fmt.Println("Bookmark ", name, url)
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
