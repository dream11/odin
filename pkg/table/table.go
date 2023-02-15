package table

import (
	"fmt"
	"os"

	"github.com/olekukonko/tablewriter"
)

// Write : write provided input as tabular format
func Write(headers []string, data [][]interface{}) {

	table := tablewriter.NewWriter(os.Stdout)

	// table properties
	table.SetRowLine(false)
	table.SetColumnSeparator("|")
	table.SetAlignment(tablewriter.ALIGN_LEFT)
	table.SetHeaderAlignment(tablewriter.ALIGN_LEFT)
	table.SetBorders(tablewriter.Border{Left: false, Top: false, Right: false, Bottom: false})
	table.SetAutoWrapText(false)
	var allHeaderColors []tablewriter.Colors
	for i := 0; i < len(headers); i++ {
		allHeaderColors = append(allHeaderColors, tablewriter.Colors{tablewriter.Bold, tablewriter.FgHiBlueColor})
	}

	//table data
	table.SetHeader(headers)
	table.SetHeaderColor(allHeaderColors...)
	for _, row := range data {
		s := make([]string, len(row))
		for i, v := range row {
			s[i] = fmt.Sprint(v)
		}
		table.Append(s)
	}
	table.Render()
}

func AppendRow(row []interface{}, colWidth []int) {
	table := tablewriter.NewWriter(os.Stdout)

	table.SetRowLine(false)
	table.SetColumnSeparator("|")
	table.SetAlignment(tablewriter.ALIGN_LEFT)
	table.SetBorders(tablewriter.Border{Left: false, Top: false, Right: false, Bottom: false})
	table.SetAutoWrapText(false)

	for index, width := range colWidth {
		table.SetColMinWidth(index, width)
	}

	s := make([]string, len(row))
	for i, v := range row {
		s[i] = fmt.Sprint(v)
	}
	table.Append(s)
	table.Render()
}

func PrintHeader(headers []string, colWidth []int) {
	table := tablewriter.NewWriter(os.Stdout)

	table.SetRowLine(false)
	table.SetColumnSeparator("|")
	table.SetAlignment(tablewriter.ALIGN_LEFT)
	table.SetHeaderAlignment(tablewriter.ALIGN_LEFT)
	table.SetBorders(tablewriter.Border{Left: false, Top: false, Right: false, Bottom: false})
	table.SetAutoWrapText(false)
	var allHeaderColors []tablewriter.Colors
	for i := 0; i < len(headers); i++ {
		allHeaderColors = append(allHeaderColors, tablewriter.Colors{tablewriter.Bold, tablewriter.FgHiBlueColor})
	}
	for index, width := range colWidth {
		table.SetColMinWidth(index, width)
	}

	table.SetHeader(headers)
	table.SetHeaderColor(allHeaderColors...)
	table.Render()
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
