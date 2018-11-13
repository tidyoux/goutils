package cmd

import (
	"fmt"

	"github.com/howeyc/gopass"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

type Command struct {
	cmd *cobra.Command
}

func New(use, short, example string, run func(*Command) error) *Command {
	var c *Command
	c = &Command{
		cmd: &cobra.Command{
			Use:     use,
			Short:   short,
			Example: example,
			RunE: func(cmd *cobra.Command, args []string) error {
				return run(c)
			},
		},
	}
	return c
}

func (c *Command) CobraCmd() *cobra.Command {
	return c.cmd
}

func (c *Command) Flags() *pflag.FlagSet {
	return c.cmd.PersistentFlags()
}

func (c *Command) Password(name string) error {
	password, err := c.Flags().GetString(name)
	if err != nil {
		return err
	}

	if len(password) == 0 {
		fmt.Printf("password:")
		ps, err := gopass.GetPasswdMasked()
		if err != nil {
			return err
		}

		c.Flags().Set(name, string(ps))
	}
	return nil
}

func (c *Command) Execute() error {
	return c.cmd.Execute()
}
