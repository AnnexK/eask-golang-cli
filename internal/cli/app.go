package cli

import "github.com/urfave/cli/v2"

const usageText = "CLI tool that helps you build, lint and test Emacs Lisp packages"

func NewApp(commands []*cli.Command) *cli.App {
	return &cli.App{
		Name: "eask",
		Usage: usageText,
		Args: true,
		Commands: commands,
	}
}
