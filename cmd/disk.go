package cmd

import (
	"fmt"
	"runtime"

	"github.com/abhijeetrajhansgithub/tsyx/internal/disk"
	"github.com/spf13/cobra"
)

var unitDisk string

var diskCmd = &cobra.Command{
	Use:   "disk",
	Short: "Display disk information",
	RunE: func(cmd *cobra.Command, args []string) error {
		// Currently, only Linux is supported.
		if runtime.GOOS != "linux" {
			return fmt.Errorf("disk is not yet supported on %s", runtime.GOOS)
		}

		diskInfo, err := disk.LinuxCollect()
		if err != nil {
			return err
		}

		// Normalize unit
		switch unitDisk {
		case "k", "kb":
			unitDisk = "kb"
		case "m", "mb":
			unitDisk = "mb"
		case "g", "gb":
			unitDisk = "gb"
		case "t", "tb":
			unitDisk = "tb"
		}

		// Validate unit
		switch unitDisk {
		case "auto", "b", "kb", "mb", "gb", "tb":
		default:
			return fmt.Errorf("invalid unit: %s", unitDisk)
		}

		fmt.Println(disk.Format(diskInfo, unitDisk))

		return nil
	},
}

func init() {
	rootCmd.AddCommand(diskCmd)

	diskCmd.Flags().StringVarP(
		&unitDisk,
		"unit",
		"u",
		"auto",
		"Disk size unit: auto, b, kb, mb, gb, tb",
	)
}