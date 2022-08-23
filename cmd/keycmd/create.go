// Copyright (C) 2022, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.
package keycmd

import (
	"errors"
	"regexp"

	"github.com/ava-labs/avalanche-cli/pkg/key"
	"github.com/ava-labs/avalanche-cli/pkg/ux"
	"github.com/spf13/cobra"
)

const (
	forceFlag = "force"
)

var (
	forceCreate bool
	filename    string
)

func createKey(cmd *cobra.Command, args []string) error {
	keyName := args[0]

	if match, _ := regexp.MatchString("\\s", keyName); match {
		return errors.New("key name contains whitespace")
	}

	if app.KeyExists(keyName) && !forceCreate {
		return errors.New("key already exists. Use --" + forceFlag + " parameter to overwrite")
	}

	if filename == "" {
		// Create key from scratch
		ux.Logger.PrintToUser("Generating new key...")
		k, err := key.NewSoft(0)
		if err != nil {
			return err
		}
		keyPath := app.GetKeyPath(keyName)
		if err := k.Save(keyPath); err != nil {
			return err
		}
		ux.Logger.PrintToUser("Key created")
		if err := printAddresses([]string{keyPath}); err != nil {
			return err
		}
	} else {
		// Load key from file
		// TODO add validation that key is legal
		ux.Logger.PrintToUser("Loading user key...")
		if err := app.CopyKeyFile(filename, keyName); err != nil {
			return err
		}
		ux.Logger.PrintToUser("Key loaded")
	}

	return nil
}

func newCreateCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create [keyName]",
		Short: "Create a signing key",
		Long: `The key create command generates a new private key to use for creating and controlling
test subnets. Keys generated by this command are NOT cryptographically secure enough to
use in production environments. DO NOT use these keys on mainnet.

The command works by generating a secp256 key and storing it with the provided keyName. You can use this key
in other commands by providing this keyName.

If you'd like to import and existing key instead of generating one from scatch, provide the --file flag.`,
		Args:         cobra.ExactArgs(1),
		RunE:         createKey,
		SilenceUsage: true,
	}

	cmd.Flags().StringVar(
		&filename,
		"file",
		"",
		"import the key from an existing key file",
	)
	cmd.Flags().BoolVarP(
		&forceCreate,
		forceFlag,
		"f",
		false,
		"overwrite an existing key with the same name",
	)
	return cmd
}
