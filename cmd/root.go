/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/atotto/clipboard"
	"github.com/spf13/cobra"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Bookmark struct {
	gorm.Model
	Name string
	Url  string
}

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "bookmark",
	Short: "A brief description of your application",
	Long: `Usage:
  bookmark <name> <url>
  bookmark -c <name> (takes URL from clipboard)
  bookmark list
  bookmark get <name>
  bookmark delete <name>`,
	Args: cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("1 - mks: root ")

		var name, url string

		// fmt.Println("mks: args è lungo:", len(args))
		// fmt.Println("mks: args è :", args)
		clip, _ := cmd.Flags().GetBool("clip")
		if clip {
			fmt.Println("mks: root -c clip")

			var err error = nil
			url, err = clipboard.ReadAll()
			if err != nil {
				log.Fatal(err)
			}
			name = args[0]

		} else {
			fmt.Println("3 - mks: root noclip")
			if len(args) == 2 {
				// fmt.Println("mks: root len2")
				name = args[0]
				url = args[1]
			} else {
				// fmt.Println("mks: root len no2")
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
		fmt.Println("5 - mks: last", name, url, args)
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.bookmark.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("clip", "c", false, "take the clipboard as bookmarkurl")
}
