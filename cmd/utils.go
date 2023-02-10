package cmd

import (
	"encoding/csv"
	"log"
	"os"
	"time"

	"github.com/fatih/color"
	"github.com/rodaine/table"
)

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

// type TableStringer struct {
// 	table.Table
// }

// func (t TableStringer) String() string {
// 	return strings.TrimSpace(t.Table.String())
// }

func tableWriter(headers []string, rows [][]string) {
	headerFmt := color.New(color.FgGreen, color.Underline).SprintfFunc()
	columnFmt := color.New(color.FgYellow).SprintfFunc()

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
