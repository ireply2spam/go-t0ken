package token

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/spf13/cobra"

	"github.com/tzero-dev/go-t0ken/cli"
	"github.com/tzero-dev/go-t0ken/commands/lockable"
	"github.com/tzero-dev/go-t0ken/commands/ownable"
	"github.com/tzero-dev/go-t0ken/contracts/token/erc20"
)

var GetterCommands = []*cobra.Command{
	&cobra.Command{
		Use:     "abi",
		Short:   "Outputs the T0ken ABI",
		Example: "t0ken investor abi",
		Args:    cobra.NoArgs,
		Run:     func(cmd *cobra.Command, args []string) { cmd.Println(erc20.T0kenABI) },
	},
	&cobra.Command{
		Use:     "bin",
		Short:   "Outputs the T0ken Binary",
		Example: "t0ken investor bin",
		Args:    cobra.NoArgs,
		Run:     func(cmd *cobra.Command, args []string) { cmd.Println(erc20.T0kenBin) },
	},
	&cobra.Command{
		Use:     "allowance <owner> <spender>",
		Short:   "Gets the amount of tokens the <owner> has approved the <spender> to transfer",
		Example: "t0ken token allowance 0xf01ff29dcbee147e9ca151a281bfdf136f66a45b 0xa01a0a93716633058d69a28fbd472fd40e7c6b79",
		Args:    cli.ChainArgs(cobra.ExactArgs(2), cli.AddressArgFunc("owner", 0), cli.AddressArgFunc("spender", 1)),
		PreRun:  connectCaller,
		Run: func(cmd *cobra.Command, args []string) {
			owner := common.HexToAddress(args[0])
			spender := common.HexToAddress(args[1])
			cli.CheckGetter(cmd)(callSession.Allowance(owner, spender))
		},
	},
	&cobra.Command{
		Use:     "balanceOf <address>",
		Short:   "Gets the balance of the given <address>",
		Example: "t0ken token balanceOf 0xf01ff29dcbee147e9ca151a281bfdf136f66a45b",
		Args:    cli.ChainArgs(cobra.MaximumNArgs(1), cli.AddressArgFunc("address", 0)),
		PreRun:  connectCaller,
		Run: func(cmd *cobra.Command, args []string) {
			addr := common.HexToAddress(args[0])
			cli.CheckGetter(cmd)(callSession.BalanceOf(addr))
		},
	},
	&cobra.Command{
		Use:     "cancellations <address>",
		Short:   "Gets the replacement address of the given <address>, or a zero-address when it has not been cancelled",
		Example: "t0ken token cancellations 0xf01ff29dcbee147e9ca151a281bfdf136f66a45b",
		Args:    cli.ChainArgs(cobra.MaximumNArgs(1), cli.AddressArgFunc("address", 0)),
		PreRun:  connectCaller,
		Run: func(cmd *cobra.Command, args []string) {
			addr := common.HexToAddress(args[0])
			cli.CheckAddressGetter(cmd)(callSession.Cancellations(addr))
		},
	},
	&cobra.Command{
		Use:     "compliance",
		Short:   "Gets the compliance contract address for the t0ken",
		Example: "t0ken token compliance",
		Args:    cobra.NoArgs,
		PreRun:  connectCaller,
		Run: func(cmd *cobra.Command, args []string) {
			cli.CheckAddressGetter(cmd)(callSession.Compliance())
		},
	},
	&cobra.Command{
		Use:     "decimals",
		Short:   "Gets the number of decimals the t0ken is set to",
		Example: "t0ken token decimals",
		Args:    cobra.NoArgs,
		PreRun:  connectCaller,
		Run: func(cmd *cobra.Command, args []string) {
			cli.CheckGetter(cmd)(callSession.Decimals())
		},
	},
	&cobra.Command{
		Use:     "getSuperseded <address>",
		Short:   "Gets the superseded address of the given <address>",
		Example: "t0ken token getSuperseded 0xf01ff29dcbee147e9ca151a281bfdf136f66a45b",
		Args:    cli.ChainArgs(cobra.MaximumNArgs(1), cli.AddressArgFunc("address", 0)),
		PreRun:  connectCaller,
		Run: func(cmd *cobra.Command, args []string) {
			addr := common.HexToAddress(args[0])
			cli.CheckAddressGetter(cmd)(callSession.GetSuperseded(addr))
		},
	},
	&cobra.Command{
		Use:     "holderAt <index>",
		Short:   "Gets the holder address at the given <index>",
		Example: "t0ken token holderAt 5",
		Args:    cli.ChainArgs(cobra.MaximumNArgs(1), cli.IntArgFunc("index", 0)),
		PreRun:  connectCaller,
		Run: func(cmd *cobra.Command, args []string) {
			index, _ := new(big.Int).SetString(args[0], 10)
			cli.CheckAddressGetter(cmd)(callSession.HolderAt(index))
		},
	},
	&cobra.Command{
		Use:     "isHolder <address>",
		Short:   "Checks if the given <address> is a current holder",
		Example: "t0ken token isHolder 0xf01ff29dcbee147e9ca151a281bfdf136f66a45b",
		Args:    cli.ChainArgs(cobra.MaximumNArgs(1), cli.AddressArgFunc("address", 0)),
		PreRun:  connectCaller,
		Run: func(cmd *cobra.Command, args []string) {
			addr := common.HexToAddress(args[0])
			cli.CheckGetter(cmd)(callSession.IsHolder(addr))
		},
	},
	&cobra.Command{
		Use:     "isSuperseded <address>",
		Short:   "Checks if the <address> is superseded by another",
		Example: "t0ken token isSuperseded 0xf01ff29dcbee147e9ca151a281bfdf136f66a45b",
		Args:    cli.ChainArgs(cobra.MaximumNArgs(1), cli.AddressArgFunc("address", 0)),
		PreRun:  connectCaller,
		Run: func(cmd *cobra.Command, args []string) {
			addr := common.HexToAddress(args[0])
			cli.CheckGetter(cmd)(callSession.IsSuperseded(addr))
		},
	},
	&cobra.Command{
		Use:     "issuer",
		Short:   "Gets the issuer of the t0ken",
		Example: "t0ken token issuer",
		Args:    cobra.NoArgs,
		PreRun:  connectCaller,
		Run: func(cmd *cobra.Command, args []string) {
			cli.CheckAddressGetter(cmd)(callSession.Issuer())
		},
	},
	&cobra.Command{
		Use:     "issuingFinished",
		Short:   "Returns if issuing has been finished",
		Example: "t0ken token issuingFinished",
		Args:    cobra.NoArgs,
		PreRun:  connectCaller,
		Run: func(cmd *cobra.Command, args []string) {
			cli.CheckGetter(cmd)(callSession.IssuingFinished())
		},
	},
	&cobra.Command{
		Use:     "name",
		Short:   "Gets the name of the t0ken",
		Example: "t0ken token name",
		Args:    cobra.NoArgs,
		PreRun:  connectCaller,
		Run: func(cmd *cobra.Command, args []string) {
			cli.CheckGetter(cmd)(callSession.Name())
		},
	},
	&cobra.Command{
		Use:     "shareholders",
		Short:   "Gets the total number of shareholders",
		Example: "t0ken token shareholders",
		Args:    cobra.NoArgs,
		PreRun:  connectCaller,
		Run: func(cmd *cobra.Command, args []string) {
			cli.CheckGetter(cmd)(callSession.Shareholders())
		},
	},
	&cobra.Command{
		Use:     "symbol",
		Short:   "Gets the symbol of the t0ken",
		Example: "t0ken token symbol",
		Args:    cobra.NoArgs,
		PreRun:  connectCaller,
		Run: func(cmd *cobra.Command, args []string) {
			cli.CheckGetter(cmd)(callSession.Symbol())
		},
	},
	&cobra.Command{
		Use:     "totalSupply",
		Short:   "Gets the total supply",
		Example: "t0ken token totalSupply",
		Args:    cobra.NoArgs,
		PreRun:  connectCaller,
		Run: func(cmd *cobra.Command, args []string) {
			cli.CheckGetter(cmd)(callSession.TotalSupply())
		},
	},
}

func init() {
	// Add the Lockable, Ownable contract getter commands
	GetterCommands = append(GetterCommands, lockable.NewGetterCommands(contractKey)...)
	GetterCommands = append(GetterCommands, ownable.NewGetterCommands(contractKey)...)

	for i, cmd := range GetterCommands {
		// Skip ABI/Bin
		if i < 2 {
			continue
		}

		// Allow providing contract 'address' flag
		cmd.Flags().String("address", "", `address of the BrokerDealer registry contract (default "[`+contractKey+`] value from config")`)
	}
}
