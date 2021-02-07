package config

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

func Version(version, commit string) Applicant {
	return func(b *Command) {
		b.Sub("version",
			Short("Show the current version"),
			Run(func(cmd *cobra.Command, args []string) {
				fmt.Printf("%s (%s)\n", version, commit)
			}),
		)
	}
}

func Completion() Applicant {
	return func(b *Command) {
		b.Sub("completion [bash|zsh|fish|powershell]",
			Short("Generate completion script"),
			ValidArgs("bash", "zsh", "fish", "powershell"),
			Args(And(One(), cobra.OnlyValidArgs)),
			Run(func(cmd *cobra.Command, args []string) {
				switch args[0] {
				case "bash":
					cmd.Root().GenBashCompletion(os.Stdout)
				case "zsh":
					cmd.Root().GenZshCompletion(os.Stdout)
				case "fish":
					cmd.Root().GenFishCompletion(os.Stdout, true)
				case "powershell":
					cmd.Root().GenPowerShellCompletion(os.Stdout)
				}
			}),
		)
	}
	//DisableFlagsInUseLine: true,
}

func FlagLogLevel(level string) Applicant {
	if level == "" {
		level = "warn"
	}
	return Flag("loglevel", Str(level),
		Abbr("l"),
		Description("Minimum level for log messages (one of 'debug', 'info', 'warn', 'error', 'fatal', 'panic')"),
		Persistent(),
		Env())
}

func FlagLogFile() Applicant {
	return Flag("logfile", Str("-"),
		Description("Write logfiles to the given file ('-' for stderr)"),
		Persistent(),
		Env())
}

func FlagLogFormat() Applicant {
	return Flag("logformat", Str("json"),
		Description("Which format to use for writing the logs ('json' or 'text')"),
		Persistent(),
		Env())
}
