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

## Commands

The general Syntax for commands is `*Command(name, ...options)`. All options are implemented as functions following the 
same interface and will be applied to the command at the end of the surrounding function call. Options are 
stateless and safe to apply multiple times, in case you have an option or flag common to multiple commands/subcommands.

### Subcommands

You can add subcommands to any command by simply adding the `SubCommand(name, ...options)` option to the parent command 
like this:

```go
NewCommandline("my_app",
    SubCommand("subcommand1",
        SubCommand("nested-subcommand1",
        	options...),
        options...),
    SubCommand("subcommand2",
        options...),
)
```

### Command Options

<dl>
<dt><code>Alias(...aliases)</code></dt><dd>Adds the given strings as aliases to this command.</dd>
<dt><code>Short(description)</code></dt><dd>Sets the short description of this command.</dd>
<dt><code>Long(description)</code></dt><dd>Sets the long description of this command.</dd>
<dt><code>ValidArgs(...args)</code></dt><dd>Adds the given string arguments as valid one's.</dd>
<dt><code>Hidden()</code></dt><dd>Marks the command as hidden.</dd>
<dt><code>Deprecated(message)</code></dt><dd>Marks the command as deprecated.</dd>
<dt><code>Args(positionalArgs)</code></dt><dd>Set cobra positional args on the command.</dd>
<dt><code>Run(function)</code></dt><dd>Set the function to run.</dd> 
</dl>


### Common Commands

## Flags

Adding flags to a command works the same way as options and sub commands. Use the `Flag(name, type, ...options)` option

### Flag Options

### Common Flags

