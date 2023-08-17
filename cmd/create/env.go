/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>

*/
package create

import (
	"fmt"
	"github.com/spf13/cobra"
	backend "github.com/dream11/odin/internal/backend"
)

// envCmd represents the env command
var name string
var envCmd = &cobra.Command{
	Use:   "env",
	Short: "Creates an environment in which services will be deployed",
	Long: `Creates an environment in which services will be deployed`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(fmt.Sprintf("env name %s", name))
		backend.GrpcClient()
	},
}

func init() {
	envCmd.Flags().StringVarP(&name, "name", "n", "", "Name of the env")
	envCmd.MarkFlagRequired("name")

	createCmd.AddCommand(envCmd)

}
