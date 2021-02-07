package config

import "github.com/spf13/cobra"

func Alias(aliases ...string) Applicant {
	return WrapBuilderOption(func(cmd *cobra.Command) {
		cmd.Aliases = aliases
	})
}

func Short(description string) Applicant {
	return WrapBuilderOption(func(cmd *cobra.Command) {
		cmd.Short = description
	})
}

func Long(description string) Applicant {
	return WrapBuilderOption(func(cmd *cobra.Command) {
		cmd.Long = description
	})
}

func ValidArgs(args ...string) Applicant {
	return WrapBuilderOption(func(cmd *cobra.Command) {
		cmd.ValidArgs = args
	})
}

func Hidden() Applicant {
	return WrapBuilderOption(func(cmd *cobra.Command) {
		cmd.Hidden = true
	})
}

func Deprecated(message string) Applicant {
	return WrapBuilderOption(func(cmd *cobra.Command) {
		cmd.Deprecated = message
	})
}

func Args(fn cobra.PositionalArgs) Applicant {
	return WrapBuilderOption(func(cmd *cobra.Command) {
		cmd.Args = fn
	})
}

func Run(fn func(cmd *cobra.Command, args []string)) Applicant {
	return WrapBuilderOption(func(cmd *cobra.Command) {
		cmd.Run = fn
	})
}
