package main

import (
	"fmt"
	"os"
	"text/tabwriter"
)

// Write : write provided input as tabular format
func Write(headers []string, data [][]interface{}) error {
	w := tabwriter.NewWriter(os.Stdout, 1, 1, 1, ' ', 0)

	var tabbedHeader string
	for _, val := range headers {
		tabbedHeader += fmt.Sprintf("%s\t|\t", val)
	}

	_, err := fmt.Fprintln(w, tabbedHeader)
	if err != nil {
		return err
	}

	for _, dataSet := range data {
		var tabbedData string
		for _, val := range dataSet {
			tabbedData += fmt.Sprintf("%v\t|\t", val)
		}

		_, err := fmt.Fprintln(w, tabbedData)
		if err != nil {
			return err
		}
	}

	return w.Flush()
}

/*
Usage -

func main() {
	headers := []string{
		"h1",
		"h2",
		"h3",
	}

	data := [][]interface{}{
		{"Row", "number", 3},
		{"Row", "number", 2},
	}

	Write(headers, data)
}
*/
