package table

import (
	"fmt"
	"os"

	"github.com/olekukonko/tablewriter"
)

func setDefaultConfig(table *tablewriter.Table) {
	table.SetRowLine(false)
	table.SetColumnSeparator("|")
	table.SetAlignment(tablewriter.ALIGN_LEFT)
	table.SetHeaderAlignment(tablewriter.ALIGN_LEFT)
	table.SetBorders(tablewriter.Border{Left: false, Top: false, Right: false, Bottom: false})
	table.SetAutoWrapText(false)
}

// Write : write provided input as tabular format
func Write(headers []string, data [][]interface{}) {

	table := tablewriter.NewWriter(os.Stdout)

	// table properties
	setDefaultConfig(table)
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
