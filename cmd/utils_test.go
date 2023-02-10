package cmd

import (
	"bytes"
	"encoding/csv"
	"io"
	"os"
	"reflect"
	"strings"
	"testing"
)

// TestSliceInterfaceConverter tests the sliceInterfaceConverter function
func TestSliceInterfaceConverter(t *testing.T) {
	t.Parallel()

	// Test data
	testData := []string{"one", "two", "three"}

	// Expected result
	expected := []interface{}{"one", "two", "three"}

	// Actual result
	actual := sliceInterfaceConverter(testData)

	// Check if the actual result is equal to the expected result
	if len(actual) != len(expected) {
		t.Errorf("sliceInterfaceConverter() = %v, want %v", actual, expected)
	}
}

// TestTimeFormatter tests the TimeFormatter function
func TestTimeFormatter(t *testing.T) {
	t.Parallel()

	// Test data
	testData := int64(1675434543)

	// Expected result
	expected := "2023-02-03T08:29:03-06:00"

	// Actual result
	actual := TimeFormatter(testData)

	// Check if the actual result is equal to the expected result
	if actual != expected {
		t.Errorf("TimeFormatter() = %v, want %v", actual, expected)
	}
}

// TestTableWriter tests the tableWriter function
func TestTableWriter(t *testing.T) {
	t.Parallel()

	// Test data
	headers := []string{"KhValue", "Solution Added", "Time"}
	rows := [][]string{
		{"7.00", "0.00", "2023-01-27T20:09:36-06:00"},
		{"7.01", "0.00", "2023-01-31T08:09:42-06:00"},
	}

	expected := `
KhValue  Solution Added  Time
7.00     0.00            2023-01-27T20:09:36-06:00
7.01     0.00            2023-01-31T08:09:42-06:00
`

	// Capture the output of tableWriter
	oldStdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	// Call the function
	tableWriter(headers, rows)

	// Read the output from the pipe
	outC := make(chan string)
	go func() {
		var buf bytes.Buffer
		io.Copy(&buf, r)
		outC <- buf.String()
	}()

	// Close the pipe
	w.Close()
	os.Stdout = oldStdout

	// Get the output and check it
	output := <-outC
	if output != strings.TrimLeft(expected, " ") {
		t.Fatalf("incorrect output: expected %v, got %v", strings.TrimLeft(expected, " "), output)
	}
}

// TestCsvWriter tests the csvWriter function
func TestCsvWriter(t *testing.T) {
	t.Parallel()

	// Test data
	headers := []string{"KhValue", "Solution Added", "Time"}
	rows := [][]string{
		{"7.00", "0.00", "2023-01-27T20:09:36-06:00"},
		{"7.01", "0.00", "2023-01-31T08:09:42-06:00"},
	}

	filename := "test.csv"
	defer func() {
		os.Remove(filename)
	}()

	// Call the function
	csvWriter(headers, rows, "test.csv")

	// Read the file
	file, err := os.Open(filename)
	if err != nil {
		t.Fatalf("failed to open file: %v", err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	readHeaders, err := reader.Read()
	if err != nil {
		t.Fatalf("failed to read headers: %v", err)
	}
	if !reflect.DeepEqual(readHeaders, headers) {
		t.Fatalf("incorrect headers: expected %v, got %v", headers, readHeaders)
	}

	var readRows [][]string
	for {
		row, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			t.Fatalf("failed to read row: %v", err)
		}
		readRows = append(readRows, row)
	}

	if !reflect.DeepEqual(readRows, rows) {
		t.Fatalf("incorrect rows: expected %v, got %v", rows, readRows)
	}

}
