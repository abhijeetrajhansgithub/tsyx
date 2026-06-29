package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/abhijeetrajhansgithub/tsyx/internal/cpu"
)

var cpuCmd = &cobra.Command{
	Use:   "cpu",
	Short: "Display CPU information",
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Println("CPU INITIALIZED AND INVOKED")

		cpuInfo := cpu.Collect()
		fmt.Println(cpu.Format(cpuInfo))

		return nil
	},
}

func init() {
	rootCmd.AddCommand(cpuCmd)
}