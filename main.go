package main

import (
	"os"

	"github.com/AnnexK/eask-golang-cli/internal/cli"
	"github.com/AnnexK/eask-golang-cli/internal/cli/cmd"
)


func main() {
	cmds := cmd.NewCommands()
	app := cli.NewApp(cmds)
	_ = app.Run(os.Args)
}
