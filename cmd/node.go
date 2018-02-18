package cmd

import (
	"errors"
	"fmt"

	"github.com/s1na/nano-go/rpc"
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
		rpc.SetRPCServer(addr)
		v, err := rpc.Version()
		if err != nil {
			return err
		}

		fmt.Printf("Node vendor: %s\nStore version: %s\nRPC Version: %s\n", v["node_vendor"], v["store_version"], v["rpc_version"])
		return nil
	},
}

var nodeStopCmd = &cobra.Command{
	Use:   "stop",
	Short: "Stop a node instance",
	Long:  `Stop a node instance`,
	RunE: func(cmd *cobra.Command, args []string) error {
		rpc.SetRPCServer(addr)
		r, err := rpc.Stop()
		if err != nil {
			return err
		}

		if r == false {
			return errors.New("Stopping the node failed")
		}

		return nil
	},
}

func init() {
	rootCmd.AddCommand(nodeCmd)
	nodeCmd.AddCommand(nodeVersionCmd)
	nodeCmd.AddCommand(nodeStopCmd)
}
