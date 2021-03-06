package nonce

import (
	"context"

	"github.com/ethereum/go-ethereum/common"
	"github.com/spf13/cobra"

	"github.com/tzero-dev/go-t0ken/cli"
)

var nonce uint64

var Command = &cobra.Command{
	Use:   "nonce",
	Short: "Nonce utilities",
}

var NextCommand = &cobra.Command{
	Use:    "next <address>",
	Short:  "gets the next nonce for an address",
	Args:   cobra.ExactArgs(1),
	PreRun: cli.Connect,
	Run: func(cmd *cobra.Command, args []string) {
		var addr common.Address
		var err error

		if common.IsHexAddress(args[0]) {
			addr, err = cli.GetArgAddress(0, args)
		} else {
			addr, _, _, err = cli.AddressForKeystoreAlias(args[0])
		}
		cli.CheckErr(cmd, err)

		n, err := cli.Conn.PendingNonceAt(context.Background(), addr)
		cli.CheckErr(cmd, err)
		cmd.Println(n)
	},
}

// SetNonce sets the transactors nonce to the given value, or to the next pending nonce when 0.
func SetNonce(nonce uint64) error {
	if nonce == 0 {
		return cli.Conn.SetNextNonce()
	} else {
		cli.Conn.SetNonce(nonce)
	}
	return nil
}

// Get returns either the nonce flag value or the next pending nonce.
func Get() (uint64, error) {
	// If the nonce flag is set, use its value
	if nonce > 0 {
		return nonce, nil
	}
	//return cli.Conn.PendingNonceAt(context.Background(), common.HexToAddress(viper.GetString("keystoreAddress")))
	n, err := cli.Conn.PendingNonceAt(context.Background(), cli.Conn.Opts.From)
	return n, err
}

// Flag adds the 'nonce' flag to the given command.
func Flag(cmd *cobra.Command) {
	cmd.Flags().Uint64Var(&nonce, "nonce", 0, "manually set the nonce for the transaction")
}
