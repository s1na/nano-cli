package cmd

import (
	"fmt"

	"github.com/s1na/nano-go/rpc"
	"github.com/spf13/cobra"
)

var walletCmd = &cobra.Command{
	Use:   "wallet",
	Short: "Manage your wallets",
	Long:  `Manage your wallets`,
}

var walletCreateCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a wallet",
	Long:  `Create a wallet`,
	RunE: func(cmd *cobra.Command, args []string) error {
		rpc.SetRPCServer(addr)
		id, err := rpc.CreateWallet()
		if err != nil {
			return err
		}

		fmt.Printf("New wallet id: %s\n", id)
		return nil
	},
}

func init() {
	rootCmd.AddCommand(walletCmd)
	walletCmd.AddCommand(walletCreateCmd)
}
