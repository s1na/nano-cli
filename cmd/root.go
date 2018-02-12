package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var (
	addr   string
	client *nano.Client
)

var rootCmd = &cobra.Command{
	Use:   "nanoc",
	Short: "CLI for Nano node",
	Long:  `A cli wallet for the Nano cryptocurrency.`,
	//	Run: func(cmd *cobra.Command, args []string) { },
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().StringVar(&addr, "addr", "http://127.0.0.1:7076", "Address of node's RPC interface")
}
