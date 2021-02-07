# Bootstrapping for Golang CLI Applications

`soil` is using [Cobra](https://github.com/spf13/cobra), [Viper](https://github.com/spf13/viper) and [Zerolog]
(https://github.com/rs/zerolog) for doing the 
actual heavy lifting and adds a convenient interface on top of them.

## Getting started

Include soil in your project using modules:

```bash
go get github.com/mtrense/soil
```

and import it in your code:

```go
package main

import "github.com/mtrense/soil"
```

A minimal skeleton for your CLI library might look like:

```go
package main

import (
    . "github.com/mtrense/soil/config"
    l "github.com/mtrense/soil/logging"
)

var (
    version = "none"  // The current version of the program 
                      // (Set with -X main.version=${VERSION} on go build)
    commit  = "none"  // The current version of the program 
                      // (Set with -X main.commit=$(git rev-parse --short HEAD 2>/dev/null || echo \"none\") on 
                      // go build)
    app = NewCommandline("my_app",      // Creates a new root command object
          Short("New app using soil"),  // Short description
          FlagLogFile(),                // Adds a flag to specify the logfile location (--logfile FILE)
          FlagLogLevel("warn"),         // Adds a flag to specify the log level (--loglevel warn)
          Run(execute),                 // Defines function to run on handling this command
          Version(version, commit),     // Adds a subcommand for showing the version and commit info 
                                        // of the current build
          Completion(),                 // Add a subcommand for completion on bash, fish and zsh
    ).GenerateCobra()
)

func init() {
    EnvironmentConfig("MY_APP")
    l.ConfigureDefaultLogging()
}

func main() {
    if err := app.Execute(); err != nil {
        panic(err)
    }
}

func execute(cmd *cobra.Command, args []string) {
    l.L().Info().Msg("My App starting")
    // Your code
}
```