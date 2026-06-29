package cmd

import (
	"fmt"
	"runtime"

	"github.com/abhijeetrajhansgithub/tsyx/internal/uptime"
	"github.com/spf13/cobra"
)

var uptimeCmd = &cobra.Command{
	Use:   "uptime",
	Short: "Display system uptime",
	RunE: func(cmd *cobra.Command, args []string) error {
		// Currently, only Linux is supported.
		if runtime.GOOS != "linux" {
			return fmt.Errorf("uptime is not yet supported on %s", runtime.GOOS)
		}

		uptimeInfo, err := uptime.LinuxCollect()
		if err != nil {
			return err
		}

		fmt.Println(uptime.Format(uptimeInfo))

		return nil
	},
}

func init() {
	rootCmd.AddCommand(uptimeCmd)
}