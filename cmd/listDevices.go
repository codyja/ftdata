/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/codyja/focustronic/api"
	"github.com/spf13/cobra"
)

// listDevicesCmd represents the listDevices command
var listDevicesCmd = &cobra.Command{
	Use:   "list-devices",
	Short: "List all available devices in your account.",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("listDevices called")

		username, ok := os.LookupEnv("FOCUSTRONIC_USERNAME")
		if !ok {
			log.Fatalf("FOCUSTRONIC_USERNAME not set")
		}
		password, ok := os.LookupEnv("FOCUSTRONIC_PASSWORD")
		if !ok {
			log.Fatalf("FOCUSTRONIC_PASSWORD not set")
		}
		client, err := api.NewFocustronicClient()
		if err != nil {
			fmt.Errorf("error initializing new Focustronic Client")
		}
		client.Authenticate(username, password)

		listDevices(client)
	},
}

func init() {
	rootCmd.AddCommand(listDevicesCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// listDevicesCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// listDevicesCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func listDevices(client *api.FocustronicClient) {
	allDevices, err := client.GetDevices()
	if err != nil {
		fmt.Println(err)
		log.Fatalf("Error getting devices: %v", err)
	}

	deviceList := client.ListDevices(allDevices)

	for k, v := range deviceList {
		fmt.Printf("Device Name: %s, Device ID: %v\n", v, k)
	}
}
