package soil

import (
	"github.com/mtrense/soil/config"
	"github.com/mtrense/soil/logging"
	"github.com/spf13/cobra"
)

func DefaultCLI(app *cobra.Command, version, commit, envPrefix string) {
	config.DefaultCLI(app, version, commit, envPrefix)
}

func ConfigureDefaultLogging() {
	logging.ConfigureDefaultLogging()
}
