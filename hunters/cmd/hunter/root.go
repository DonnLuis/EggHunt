/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "hunter",
	Short: "A fun way to harden your network",
	Long: `A fun way to harden your network using an interactive cli. For example:

			The hunter of ports is a fun way to scan for opened ports. Everytime a port comes back opened,
			an eggs pops up. In this instance, Eggs are the ports!`,
}

// version Command
var versionCmd = &cobra.Command{
	Use: "version",
	Short: "Display the version of hunter",
	Long: "Display the current version of the hunter cli tool",
	Run: func(cmd *cobra.Command, args []string){
		fmt.Println("Hunter v1.0")
	},

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// commands
	rootCmd.AddCommand(versionCmd)

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.hunter.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	rootCmd.Flags().BoolP("hunter", "h", false, "hunter help")
}


