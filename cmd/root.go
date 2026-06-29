// root.go
package cmd

import "github.com/spf13/cobra"

var rootCmd = &cobra.Command{
	Use:   "tsyx",
	Short: "A CLI for TSYX",
}

func Execute() error {
	return rootCmd.Execute()
}