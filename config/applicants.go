package config

func WrapBuilderOption(fn ...CommandOption) Applicant {
	return func(builder *Command) {
		builder.opts = append(builder.opts, fn...)
	}
}

func SubCommand(name string, applicants ...Applicant) Applicant {
	return func(builder *Command) {
		subCommand := &Command{
			name: name,
		}
		subCommand.apply(applicants...)
		builder.subCommands = append(builder.subCommands, subCommand)
	}
}
