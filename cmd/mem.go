package cmd

import (
	"fmt"
	"runtime"

	"github.com/abhijeetrajhansgithub/tsyx/internal/memory"
	"github.com/spf13/cobra"
)

var (
	viewMem string
	unitMem string
)

var memCmd = &cobra.Command{
	Use:   "mem",
	Short: "Display memory information",
	RunE: func(cmd *cobra.Command, args []string) error {
		// Currently, only Linux is supported.
		if runtime.GOOS != "linux" {
			return fmt.Errorf("memory is not yet supported on %s", runtime.GOOS)
		}

		memInfo, err := memory.LinuxCollect()
		if err != nil {
			return err
		}

		// Normalize view
		switch viewMem {
		case "s", "summary":
			viewMem = "summary"
		case "d", "detailed":
			viewMem = "detailed"
		case "k", "kernel":
			viewMem = "kernel"
		}

		// Validate view
		switch viewMem {
		case "summary", "detailed", "kernel":
		default:
			return fmt.Errorf("invalid view: %s", viewMem)
		}

		// Normalize unit
		switch unitMem {
		case "k", "kb":
			unitMem = "kb"
		case "m", "mb":
			unitMem = "mb"
		case "g", "gb":
			unitMem = "gb"
		case "t", "tb":
			unitMem = "tb"
		}

		// Validate unit
		switch unitMem {
		case "auto", "b", "kb", "mb", "gb", "tb":
		default:
			return fmt.Errorf("invalid unit: %s", unitMem)
		}

		switch viewMem {
		case "summary":
			fmt.Println(memory.FormatSummary(memInfo, unitMem))
		case "detailed":
			fmt.Println(memory.FormatDetailed(memInfo, unitMem))
		case "kernel":
			fmt.Println(memory.FormatKernel(memInfo, unitMem))
		}

		return nil
	},
}

func init() {
	rootCmd.AddCommand(memCmd)

	memCmd.Flags().StringVarP(
		&viewMem,
		"view",
		"v",
		"summary",
		"View: summary (s), detailed (d), kernel (k)",
	)

	memCmd.Flags().StringVarP(
		&unitMem,
		"unit",
		"u",
		"auto",
		"Memory unit: auto, b, kb, mb, gb, tb",
	)
}