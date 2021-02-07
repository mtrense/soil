package config

import (
	"io"

	"github.com/spf13/cobra"
)

type Command struct {
	parent      *Command
	name        string
	opts        []CommandOption
	subCommands []*Command
	flags       map[string]*FlagInfo
}

type CommandOption func(cmd *cobra.Command)
type Applicant func(builder *Command)

func NewCommandline(name string, applicants ...Applicant) *Command {
	command := &Command{
		name: name,
	}
	command.apply(applicants...)
	return command
}

func (s *Command) apply(applicants ...Applicant) {
	for _, applicant := range applicants {
		applicant(s)
	}
}

func (s *Command) Sub(name string, applicants ...Applicant) *Command {
	subCommand := &Command{
		name: name,
	}
	subCommand.apply(applicants...)
	s.subCommands = append(s.subCommands, subCommand)
	subCommand.parent = s
	return subCommand
}

func (s *Command) S(name string, opts ...Applicant) *Command {
	return s.Sub(name, opts...)
}

func (s *Command) Apply(applicants ...Applicant) *Command {
	for _, applicant := range applicants {
		applicant(s)
	}
	return s
}

func (s *Command) A(fn func(c *Command)) *Command {
	return s.Apply(fn)
}

func (s *Command) Debug(w io.StringWriter) {
	//w.
}

func (s *Command) GenerateCobra() *cobra.Command {
	cmd := &cobra.Command{
		Use: s.name,
	}
	for _, opt := range s.opts {
		opt(cmd)
	}
	for _, subCommand := range s.subCommands {
		subCmd := subCommand.GenerateCobra()
		cmd.AddCommand(subCmd)
	}
	return cmd
}
