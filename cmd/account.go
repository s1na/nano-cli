package cmd

import (
	"fmt"

	"github.com/s1na/nano-go/rpc"
	"github.com/spf13/cobra"
)

var accountCmd = &cobra.Command{
	Use:   "account",
	Short: "Manage your accounts",
	Long:  `Manage your accounts`,
}

var accountCreateCmd = &cobra.Command{
	Use:   "create WALLET_ID",
	Short: "Create an account",
	Long:  `Create an account in the provided wallet`,
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		rpc.SetRPCServer(addr)
		id, err := rpc.CreateAccount(args[0], true)
		if err != nil {
			return err
		}

		fmt.Printf("New account id: %s\n", id)
		return nil
	},
}

var accountListCmd = &cobra.Command{
	Use:   "list WALLET_ID",
	Short: "List accounts in a wallet",
	Long:  `List accounts in a wallet`,
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		rpc.SetRPCServer(addr)
		accounts, err := rpc.AccountList(args[0])
		if err != nil {
			return err
		}

		fmt.Printf("%v\n", accounts)
		return nil
	},
}

func init() {
	rootCmd.AddCommand(accountCmd)
	accountCmd.AddCommand(accountCreateCmd)
	accountCmd.AddCommand(accountListCmd)
}
