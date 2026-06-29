package cmd

import (
	"fmt"
	"runtime"

	"github.com/spf13/cobra"
	"github.com/abhijeetrajhansgithub/tsyx/internal/cpu"
)

var cpuCmd = &cobra.Command{
	Use:   "cpu",
	Short: "Display CPU information",
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Println("CPU INITIALIZED AND INVOKED")

		_os := runtime.GOOS
		fmt.Printf("OS: %s\n", _os)

		if _os != "linux" {
			return fmt.Errorf("unsupported OS: %s", _os)
		}
		cpuInfo := cpu.LinuxCollect()
		fmt.Println(cpu.Format(cpuInfo))

		return nil
	},
}

func init() {
	rootCmd.AddCommand(cpuCmd)
}