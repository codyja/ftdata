package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"github.com/codyja/focustronic/api"
	"io/ioutil"

	//"io/ioutil"
	"log"
	"os"
	"time"

	//"time"

	"github.com/fatih/color"
	"github.com/rodaine/table"
)

const (
	usage = `usage: %s

ftdata - Focustronic Data
This tool is used to login to your Focustronic account and extra your data that has been uploaded by your devices.

Environment Variables:
FOCUSTRONIC_USERNAME="user_here"
FOCUSTRONIC_PASSWORD="pass_here"

Options:
`
)

func tableWriter(headers []string, rows [][]string) {
	headerFmt := color.New(color.FgGreen, color.Underline).SprintfFunc()
	columnFmt := color.New(color.FgYellow).SprintfFunc()

	log.Printf("Creating table")

	h := sliceInterfaceConverter(headers)

	tbl := table.New(h...)
	tbl.WithHeaderFormatter(headerFmt).WithFirstColumnFormatter(columnFmt)

	for _, row := range rows {
		r := sliceInterfaceConverter(row)
		tbl.AddRow(r...)
	}

	tbl.Print()
}

func csvWriter(headers []string, rows [][]string, file string) {
	output := os.Stdout
	if file != "" {
		csvFile, err := os.Create(file)
		if err != nil {
			log.Fatalf("failed creating file: %s", err)
		}
		defer csvFile.Close()

		output = csvFile

		log.Printf("Saving CSV to %s", file)

	} else {
		log.Printf("Displaying CSV on stdout")
	}

	w := csv.NewWriter(output)
	defer w.Flush()


	if err := w.Write(headers); err != nil {
		log.Fatal("Error writing csv output: ", err)
	}

	for _, row := range rows {
		if err := w.Write(row); err != nil {
			log.Fatal("Error writing csv output: ", err)
		}
	}
}

func sliceInterfaceConverter(s []string) []interface{} {
	i := make([]interface{}, len(s))
	for k, v := range s {
		i[k] = v
	}
	return i
}

func TimeFormatter(timeStamp int64) string {
	t := time.Unix(timeStamp, 0)
	f := t.Format(time.RFC3339)
	return f
}

func GetRecords(c *api.FocustronicClient, deviceType string, deviceId int, days , output , file string) {

	if  deviceType == "alkatronic" {
		records, err := c.GetRecords(deviceType, deviceId, days)
		if err != nil {
			log.Fatalf("Error getting record data: #{err}", err)
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

	if  deviceType == "mastertronic" {
		records, err := c.GetRecords(deviceType, deviceId, days)
		if err != nil {
			log.Fatalf("Error getting record data: #{err}", err)
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
			log.Fatalf("Error getting record data: #{err}", err)
		}

		headers := []string{"PumpId", "DoseVolume", "DoseMode", "Time"}
		var rows [][]string

		for _, v := range records.Data {
			for _, v := range v {
				row := []string{
					fmt.Sprintf("%d", v.PumpId),
					fmt.Sprintf( "%.2f", v.DoseVolume / 100),
					fmt.Sprintf( "%d", v.DoseMode),
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

func listDevices(client *api.FocustronicClient) {
	allDevices, err := client.GetAllDevices()
	if err != nil {
		log.Fatalf("Error getting devices: %s", err, err)
	}

	deviceList := client.ListDevices(allDevices)

	for k, v := range deviceList {
		fmt.Printf("Device Name: %s, Device ID: %v\n", v, k)
	}

}

func main() {

	username, ok := os.LookupEnv("FOCUSTRONIC_USERNAME")
	if !ok {
		log.Fatalf("FOCUSTRONIC_USERNAME not set")
	}
	password, ok := os.LookupEnv("FOCUSTRONIC_PASSWORD")
	if !ok {
		log.Fatalf("FOCUSTRONIC_PASSWORD not set")
	}

	log.SetOutput(ioutil.Discard)

	// Initialize new Focustronic Client
	client, err := api.NewFocustronicClient()
	if err != nil {
		fmt.Errorf("error initializing new Focustronic Client")
	}

	// read flags
	flagDebug := flag.Bool("debug", false, "Runs in debug mode with extra logging output")
	flagList := flag.Bool("list", false, "List devices registered to your account")
*.csv
	flagDeviceType := flag.String("type", "", "Specify the device `model` to retrieve data for. Valid selections are 'alkatronic', 'dosetronic', or 'mastertronic'.")
	flagDeviceId := flag.Int("device-id", 0, "Specify the device id for the specific device you want to retrieve records for")
	flagDays := flag.String("days", "", "Specify the number of days worth of records to retrieve. Valid values are '7', '30', '90'.")
	flagFormat := flag.String("format", "table", "Output format. Possible values are 'table' or 'csv'.")
	flagFile := flag.String("file", "", "File to save results. Overwrites file if existing. If not specified, prints data directly to terminal stdout.")
    //flagTimeFormat := flag.String("time-format", "friendly", "By default shows time in '2006-01-02 15:04:05.999999999 -0700 MST' format, but can also specify 'epoch'.")
	flag.Usage = func() {
		fmt.Fprintf(flag.CommandLine.Output(), usage, os.Args[0])
		flag.PrintDefaults()
	}
	flag.Parse()

	if *flagDebug {
		fmt.Println("running in debug")
		if *flagDebug {
			log.SetOutput(os.Stdout)
		}
	}

	// Call auth methods
	client.Authenticate(username, password)

	if *flagList {
		listDevices(client)
		os.Exit(0)
	}

	if *flagDeviceType != "" && *flagDeviceId != 0 && *flagDays != ""  {
		log.Printf("Records requested for type=%s id=%d days=%s", *flagDeviceType, *flagDeviceId, *flagDays  )
		GetRecords(client, *flagDeviceType, *flagDeviceId, *flagDays, *flagFormat, *flagFile)
	} else {
		log.Fatalf("Error, must specify flags....")
	}


}
