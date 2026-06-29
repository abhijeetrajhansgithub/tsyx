package cmd

import (
	"fmt"
	"runtime"

	"github.com/spf13/cobra"
	"github.com/abhijeetrajhansgithub/tsyx/internal/memory"
)

var (
	view string
	unit string
)

var memCmd = &cobra.Command{
	Use:   "mem",
	Short: "Display memory information",
	RunE: func(cmd *cobra.Command, args []string) error {
		// Currently, only Linux is supported.
		_os := runtime.GOOS

		if _os != "linux" {
			return fmt.Errorf("memory is not yet supported on %s", runtime.GOOS)
		}

		memInfo, err := memory.LinuxCollect()
		if err != nil {
			return err
		}

		// validate view
		switch view {
			case "summary":
			case "detailed":
			case "kernel":
			default:
				return fmt.Errorf("invalid view: %s", view)
		}

		// validate unit
		switch unit {
			case "kb", "k", "mb", "m", "gb", "g", "tb", "t":
			default:
				return fmt.Errorf("invalid unit: %s", unit)
		}

		switch view {
			case "summary":
				fmt.Println(memory.FormatSummary(memInfo, unit))
			case "detailed":
				fmt.Println(memory.FormatDetailed(memInfo, unit))
			case "kernel":
				fmt.Println(memory.FormatKernel(memInfo, unit))
		}

		return nil
	},
}

func init() {
	rootCmd.AddCommand(memCmd)

	memCmd.Flags().StringVarP(
		&view,
		"view",
		"v",
		"summary",
		"View: summary, detailed, kernel",
	)

	memCmd.Flags().StringVarP(
		&unit,
		"unit",
		"u",
		"kb",
		"Memory unit: kb, mb, gb, tb",
	)
}