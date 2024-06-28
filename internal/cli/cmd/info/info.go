package info

import (
	"fmt"

	"github.com/AnnexK/eask-golang-cli/internal/env"
	"github.com/urfave/cli/v2"
)

const description = "print eask environment info such as emacs path, executable path and elisp source directory"

func NewInfoCommand() *cli.Command {
	return &cli.Command{
		Name:        "info",
		Usage:       "print eask environment info",
		UsageText:   "eask info",
		Description: description,
		Category:    "info",
		Args:        false,
		Action: func(ctx *cli.Context) error {
			emacsPresent := "present"
			if !env.Info.EmacsPresent {
				emacsPresent = "not present"
			}
			fmt.Printf("Emacs executable path: %s (%s)\n", env.Info.EmacsPath, emacsPresent)

			execPath := env.Info.ExecutablePath
			if execPath == "" {
				execPath = "<unknown>"
			}
			fmt.Printf("Eask executable path: %s\n", execPath)

			elispPath := env.Info.ElispPath
			if elispPath == "" {
				elispPath = "<unknown>"
			}
			fmt.Printf("Eask Elisp source directory: %s\n", env.Info.ElispPath)

			return nil
		},
		HideHelp: true,
	}
}
