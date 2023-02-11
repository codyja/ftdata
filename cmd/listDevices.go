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
	Short: "List all the available devices in your account.",
	Long: `Lists all the available devices in you account
and prints the device ID.`,
	Run: func(cmd *cobra.Command, args []string) {

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
