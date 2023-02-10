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
	Use:   "ftdata",
	Short: "ftdata is a CLI data retrieval tool for Focustronic product owners.",
	Long: `ftdata is used to retreive data for Focustronic product owners. It
can list all the devices in your account, and retreive data such as Alkatronic
and Mastertronic test results. It can also retrieve Dosetronic dosing records
as well. Examples:

ftdata list-devices
ftdata get --type alkatronic --id 001
ftdata get --type mastertronic --id 001 --output csv`,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

var Verbose bool

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.ftdata.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	rootCmd.PersistentFlags().BoolVarP(&Verbose, "verbose", "v", false, "verbose level output")

}
