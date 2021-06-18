package config

import (
	"time"

	"github.com/iancoleman/strcase"

	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

type FlagInfo struct {
	name         string
	persistent   bool
	mandatory    bool
	dirname      bool
	filename     []string
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
	if s.dirname {
		cmd.MarkFlagDirname(s.name)
	}
	if s.filename != nil {
		cmd.MarkFlagFilename(s.name, s.filename...)
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

// Env option allows this Flag to be set from the environment. Please note that the name of the Flag is
// converted to snake_case. For more control over the name use EnvName.
func Env() FlagOption {
	return func(fi *FlagInfo) {
		fi.env = strcase.ToSnake(fi.name)
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

func Filename(extensions ...string) FlagOption {
	return func(fi *FlagInfo) {
		fi.filename = extensions
	}
}

func Dirname() FlagOption {
	return func(fi *FlagInfo) {
		fi.dirname = true
	}
}

// Defines a string flag with a default value.
func Str(def string) FlagType {
	return func(fi *FlagInfo, fs *pflag.FlagSet) {
		if fi.abbreviation == "" {
			fs.String(fi.name, def, fi.description)
		} else {
			fs.StringP(fi.name, fi.abbreviation, def, fi.description)
		}
	}
}

// Defines a boolean flag. Boolean flags are false by default.
func Bool() FlagType {
	return func(fi *FlagInfo, fs *pflag.FlagSet) {
		if fi.abbreviation == "" {
			fs.Bool(fi.name, false, fi.description)
		} else {
			fs.BoolP(fi.name, fi.abbreviation, false, fi.description)
		}
	}
}

// Defines an integer flag with a default value.
func Int(def int) FlagType {
	return func(fi *FlagInfo, fs *pflag.FlagSet) {
		if fi.abbreviation == "" {
			fs.Int(fi.name, def, fi.description)
		} else {
			fs.IntP(fi.name, fi.abbreviation, def, fi.description)
		}
	}
}

// Defines a decimal flag with a default value.
func Float64(def float64) FlagType {
	return func(fi *FlagInfo, fs *pflag.FlagSet) {
		if fi.abbreviation == "" {
			fs.Float64(fi.name, def, fi.description)
		} else {
			fs.Float64P(fi.name, fi.abbreviation, def, fi.description)
		}
	}
}

// Defines a duration flag with a default value.
func Duration(def time.Duration) FlagType {
	return func(fi *FlagInfo, fs *pflag.FlagSet) {
		if fi.abbreviation == "" {
			fs.Duration(fi.name, def, fi.description)
		} else {
			fs.DurationP(fi.name, fi.abbreviation, def, fi.description)
		}
	}
}
