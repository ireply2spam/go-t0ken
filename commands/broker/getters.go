package broker

import (
	"github.com/spf13/cobra"

	"github.com/tzero-dev/go-t0ken/cli"
	"github.com/tzero-dev/go-t0ken/commands/lockable"
	"github.com/tzero-dev/go-t0ken/commands/ownable"
	"github.com/tzero-dev/go-t0ken/contracts/registry"
)

var GetterCommands = []*cobra.Command{
	&cobra.Command{
		Use:     "abi",
		Short:   "Outputs the Storage ABI",
		Example: "t0ken investor abi",
		Args:    cobra.NoArgs,
		Run:     func(cmd *cobra.Command, args []string) { cmd.Println(registry.StorageABI) },
	},
	&cobra.Command{
		Use:     "bin",
		Short:   "Outputs the Storage Binary",
		Example: "t0ken investor bin",
		Args:    cobra.NoArgs,
		Run:     func(cmd *cobra.Command, args []string) { cmd.Println(registry.StorageBin) },
	},
	&cobra.Command{
		Use:     "storage",
		Short:   "Gets the Storage contract address",
		Example: "t0ken investor storage",
		Args:    cobra.NoArgs,
		PreRun:  connectCaller,
		Run: func(cmd *cobra.Command, args []string) {
			cli.CheckAddressGetter(cmd)(callSession.Store())
		},
	},
}

func init() {
	// Add the Ownable, Lockable contract getter commands
	GetterCommands = append(GetterCommands, ownable.NewGetterCommands(contractKey)...)
	GetterCommands = append(GetterCommands, lockable.NewGetterCommands(contractKey)...)

	for i, cmd := range GetterCommands {
		// Skip ABI/Bin
		if i < 2 {
			continue
		}

		// Allow providing contract 'address' flag
		cmd.Flags().String("address", "", `address of the BrokerDealer registry contract (default "[`+contractKey+`] value from config")`)
	}
}
