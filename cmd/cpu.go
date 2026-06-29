package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var cpuCmd = &cobra.Command{
	Use:   "cpu",
	Short: "Display CPU information",
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Println("CPU INITIALIZED AND INVOKED")

		return nil
	},
}

func init() {
	rootCmd.AddCommand(cpuCmd)
}