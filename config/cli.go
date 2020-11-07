package config

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"os"
)

func DefaultCLI(app *cobra.Command, version, commit, envPrefix string) {
	app.AddCommand(CompletionCommand(), VersionCommand(version, commit))
	ApplyLogFlags(app)
	viper.SetEnvPrefix(envPrefix)
	viper.AutomaticEnv()
}

func ApplyLogFlags(app *cobra.Command) {
	app.PersistentFlags().StringP("loglevel", "l", "warn", "Minimum level for logmessages (one of 'debug', 'info', 'warn', 'error', 'fatal', 'panic'")
	viper.BindPFlag("loglevel", app.PersistentFlags().Lookup("loglevel"))
	app.PersistentFlags().String("logfile", "-", "Write logfiles to the given file (- for stderr)")
	viper.BindPFlag("logfile", app.PersistentFlags().Lookup("logfile"))
	app.PersistentFlags().String("logformat", "json", "Which format to use for writing the logs (json or text)")
	viper.BindPFlag("logformat", app.PersistentFlags().Lookup("logformat"))
}

func VersionCommand(version, commit string) *cobra.Command {
	return &cobra.Command{
		Use:   "version",
		Short: "Show engage version",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Printf("Version: %s (ref: %s)\n", version, commit)
		},
	}
}

func CompletionCommand() *cobra.Command {
	return &cobra.Command{
		Use:                   "completion [bash|zsh|fish|powershell]",
		Short:                 "Generate completion script",
		DisableFlagsInUseLine: true,
		ValidArgs:             []string{"bash", "zsh", "fish", "powershell"},
		Args:                  cobra.ExactValidArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
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
		},
	}
}
