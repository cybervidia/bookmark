/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
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
  bookmark add <name> <url>
  bookmark add -c <name> (takes URL from clipboard)
  bookmark list (list all the bookmark)
  bookmark get <name> (put the url in the clipboard)
  bookmark delete <name> (delete a bookmark)`,
	// Args: cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Bookmark app")
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
	//rootCmd.Flags().BoolP("clip", "c", false, "take the clipboard as bookmarkurl")
}
