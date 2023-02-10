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

// getCmd represents the get command
var getCmd = &cobra.Command{
	Use:   "get",
	Short: "ftdata is a data retrieval tool for Focustronic product owners.",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
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

		deviceType, _ := cmd.Flags().GetString("type")
		output, _ := cmd.Flags().GetString("output")
		days, _ := cmd.Flags().GetInt("days")
		deviceid, _ := cmd.Flags().GetInt("id")
		file, _ := cmd.Flags().GetString("file")
		fmt.Printf("Device type: %s, Days: %d, Output: %s\n", deviceType, days, output)
		GetRecords(client, deviceType, deviceid, days, output, file)

	},
}

var DeviceType, Output, File string
var Days, ID int

func init() {
	rootCmd.AddCommand(getCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// getCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// getCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	getCmd.Flags().StringVarP(&DeviceType, "type", "t", "", "Device type to retreive data for. Valid types are alkatronic, dosetronic, or mastertronic")
	getCmd.Flags().IntVarP(&ID, "id", "i", 0, "ID for device. Can be retrieved from the 'list-devices' command.")
	getCmd.Flags().IntVarP(&Days, "days", "d", 7, "Days worth of data to retrieve. Valid options are 7, 30, or 90.")
	getCmd.Flags().StringVarP(&Output, "output", "o", "table", "Output format. Valid options are table or csv.")
	getCmd.Flags().StringVarP(&File, "file", "f", "", "Output file name and location.")

}

func GetRecords(c *api.FocustronicClient, deviceType string, deviceId int, days int, output, file string) {

	if deviceType == "alkatronic" {
		records, err := c.GetAlkatronicRecords(deviceId, days)
		if err != nil {
			log.Fatalf("error getting alkatronic records: %v", err)
		}

		headers := []string{"KhValue", "Solution Added(mL)", "Time"}
		var rows [][]string

		for _, v := range records.Data {
			row := []string{
				fmt.Sprintf("%.2f", v.KhValue),
				fmt.Sprintf("%.2f", v.SolutionAdded),
				//time.Unix(v.CreateTime, 0).String(),
				TimeFormatter(v.CreateTime),
			}
			rows = append(rows, row)
		}

		if output == "table" {
			tableWriter(headers, rows)
		}
		if output == "csv" {
			csvWriter(headers, rows, file)
		}
	}

	if deviceType == "mastertronic" {
		records, err := c.GetMastertronicRecords(days, deviceId, "")
		if err != nil {
			log.Fatalf("Error getting record data: %v", err)
		}

		headers := []string{"Parameter", "Value", "Time"}
		var rows [][]string

		for _, v := range records.Data {
			row := []string{
				v.Parameter,
				fmt.Sprintf("%.2f", v.Value),
				//time.Unix(int64(v.RecordTime), 0).String(),
				TimeFormatter(v.RecordTime),
			}
			rows = append(rows, row)
		}

		if output == "table" {
			tableWriter(headers, rows)
		}
		if output == "csv" {
			csvWriter(headers, rows, file)
		}
	}

	if deviceType == "dosetronic" {
		records, err := c.GetDosetronicRecords(deviceId, days)
		if err != nil {
			log.Fatalf("Error getting record data: %v", err)
		}

		headers := []string{"PumpId", "DoseVolume", "DoseMode", "Time"}
		var rows [][]string

		for _, v := range records.Data {
			for _, v := range v {
				row := []string{
					fmt.Sprintf("%d", v.PumpID),
					fmt.Sprintf("%.2f", v.DoseVolume/100),
					fmt.Sprintf("%d", v.DoseMode),
					TimeFormatter(v.RecordTime),
				}

				rows = append(rows, row)

			}
		}

		if output == "table" {
			tableWriter(headers, rows)
		}
		if output == "csv" {
			csvWriter(headers, rows, file)
		}
	}
}
