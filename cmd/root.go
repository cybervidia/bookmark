/*
Copyright © 2025 maKs <eliteknow@youknowwhere.to>
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
	Name string `gorm:"unique"`
	Url  string
}

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "bookmark",
	Short: " Bookmark manager CLI with add, list, get, and delete commands",
	Long: `Bookmark is a simple command-line application to manage your bookmarks locally.
You can add bookmarks by name and URL, list all saved bookmarks,
retrieve a bookmark URL to the clipboard, or delete bookmarks by name.

Usage examples:
  bookmark add <name> <url>           Add a bookmark with the given name and URL
  bookmark add -c <name>               Add a bookmark using the URL from the clipboard
  bookmark list                       Display all saved bookmarks
  bookmark get <name>                 Copy the URL of the named bookmark to the clipboard
  bookmark delete <name>              Remove the named bookmark
`,
	// Args: cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("dP                         dP                                    dP       ")
		fmt.Println("88                         88                                    88       ")
		fmt.Println("88d888b. .d8888b. .d8888b. 88  .dP  88d8b.d8b. .d8888b. 88d888b. 88  .dP  ")
		fmt.Println("88'  `88 88'  `88 88'  `88 88888\"   88'`88'`88 88'  `88 88'  `88 88888\"   ")
		fmt.Println("88.  .88 88.  .88 88.  .88 88  `8b. 88  88  88 88.  .88 88       88  `8b. ")
		fmt.Println("88Y8888' `88888P' `88888P' dP   `YP dP  dP  dP `88888P8 dP       dP   `YP ")
		fmt.Println("                                                                          ")
		fmt.Println("                  [ 記 ] ki                                                 ")
		fmt.Println("               What you mark, remembers you.                               ")

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
