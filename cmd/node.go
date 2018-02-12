package cmd

import (
	"fmt"

	"github.com/s1na/nano-go"
	"github.com/spf13/cobra"
)

var nodeCmd = &cobra.Command{
	Use:   "node",
	Short: "Node-related actions",
	Long:  `Get info regarding a node instance, or manipulate it`,
}

var nodeVersionCmd = &cobra.Command{
	Use:   "version",
	Short: "Version info of a node instance",
	Long:  `Version info of a node instance`,
	RunE: func(cmd *cobra.Command, args []string) error {
		client := nano.NewClient(addr)
		v, err := client.Version()
		if err != nil {
			return err
		}

		fmt.Printf("%v\n", v)
		return nil
	},
}

var nodeStopCmd = &cobra.Command{
	Use:   "stop",
	Short: "Stop a node instance",
	Long:  `Stop a node instance`,
	RunE: func(cmd *cobra.Command, args []string) error {
		client := nano.NewClient(addr)
		v, err := client.Stop()
		if err != nil {
			return err
		}

		return nil
	},
}

func init() {
	rootCmd.AddCommand(nodeCmd)
	nodeCmd.AddCommand(nodeVersionCmd)
	nodeCmd.AddCommand(nodeStopCmd)
}
