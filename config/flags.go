package config

import (
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

type FlagInfo struct {
	name         string
	persistent   bool
	mandatory    bool
	abbreviation string
	description  string
	env          string
	flagType     FlagType
}

func (s *FlagInfo) apply(cmd *cobra.Command) {
	var flagSet *pflag.FlagSet
	if s.persistent {
		flagSet = cmd.PersistentFlags()
	} else {
		flagSet = cmd.Flags()
	}
	s.flagType(s, flagSet)
	flag := flagSet.Lookup(s.name)
	if s.env != "" {
		viper.BindPFlag(s.env, flag)
	}
	if s.mandatory {
		cmd.MarkFlagRequired(s.name)
	}
}

type FlagOption func(fi *FlagInfo)
type FlagType func(fi *FlagInfo, fs *pflag.FlagSet)

func Flag(name string, flagType FlagType, options ...FlagOption) Applicant {
	fi := &FlagInfo{
		name:     name,
		flagType: flagType,
	}
	for _, opt := range options {
		opt(fi)
	}
	return WrapBuilderOption(func(cmd *cobra.Command) {
		fi.apply(cmd)
	})
}

func Persistent() FlagOption {
	return func(fi *FlagInfo) {
		fi.persistent = true
	}
}

func Description(desc string) FlagOption {
	return func(fi *FlagInfo) {
		fi.description = desc
	}
}

func Abbr(char string) FlagOption {
	return func(fi *FlagInfo) {
		fi.abbreviation = char
	}
}

func Env() FlagOption {
	return func(fi *FlagInfo) {
		fi.env = fi.name
	}
}

func EnvName(name string) FlagOption {
	return func(fi *FlagInfo) {
		fi.env = name
	}
}

func Mandatory() FlagOption {
	return func(fi *FlagInfo) {
		fi.mandatory = true
	}
}

func Str(def string) FlagType {
	return func(fi *FlagInfo, fs *pflag.FlagSet) {
		if fi.abbreviation == "" {
			fs.String(fi.name, def, fi.description)
		} else {
			fs.StringP(fi.name, fi.abbreviation, def, fi.description)
		}
	}
}
