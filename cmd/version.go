package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Version number of nanoc",
	Long:  `Prints the version number of nanoc`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Nanoc v0.1.0")
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
