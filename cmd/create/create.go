/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>

*/
package create

import (

	"github.com/spf13/cobra"
	"github.com/dream11/odin/cmd"
)

// createCmd represents the create command
var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Creates resources",
	Long: `Creates resources`,
}

func init() {
    cmd.RootCmd.AddCommand(createCmd)
}
