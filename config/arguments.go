package config

import (
	"fmt"

	"github.com/spf13/cobra"
)

func And(fns ...cobra.PositionalArgs) cobra.PositionalArgs {
	return func(cmd *cobra.Command, args []string) error {
		var err error
		for _, f := range fns {
			err = f(cmd, args)
			if err != nil {
				return err
			}
		}
		return nil
	}
}

func Range(min, max int) cobra.PositionalArgs {
	return func(cmd *cobra.Command, args []string) error {
		actual := len(args)
		if actual < min {
			return fmt.Errorf("too few arguments (expected at least %d, got %d)", min, actual)
		}
		if actual > max {
			return fmt.Errorf("too few arguments (expected at most %d, got %d)", max, actual)
		}
		return nil
	}
}

func One() cobra.PositionalArgs {
	return Range(1, 1)
}
