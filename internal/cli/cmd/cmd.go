package cmd

import (
	"github.com/AnnexK/eask-golang-cli/internal/cli/cmd/info"
	"github.com/urfave/cli/v2"
)


func NewCommands() []*cli.Command {
	return []*cli.Command{
		info.NewInfoCommand(),
	}
}
