package env

import (
	"log/slog"
	"os"
	"path/filepath"

	"github.com/hairyhenderson/go-which"
)

type Information struct {
	EmacsPath string
	EmacsPresent bool
	ExecutablePath string
	ElispPath      string
}

var (
	Info Information
)

func init() {
	Info = NewDefaultInformation()
}

func NewDefaultInformation() Information {
	emacsPath, emacsPresent := getEmacsPath()
	execPath := getExecutablePath()
	elispPath := getElispPath(execPath)

	return Information{
		EmacsPath: emacsPath,
		EmacsPresent: emacsPresent,
		ExecutablePath: execPath,
		ElispPath: elispPath,
	}
}

func getEmacsPath() (string, bool) {
	path := os.Getenv("EASK_EMACS")
	if path == "" {
		path = os.Getenv("EMACS")
	}
	if path == "" {
		path = "emacs"
	}

	fullPath := which.Which(path)
	if fullPath != "" {
		path = fullPath
	}
	return path, fullPath != ""
}

func getExecutablePath() string {
	path, err := os.Executable()
	if err != nil {
		slog.Error("Failed to get executable path", slog.String("error", err.Error()))
		return ""
	}
	return path
}

func getElispPath(execPath string) string {
	if execPath == "" {
		slog.Error("Unable to get Elisp source path; exec path is empty")
		return ""
	}

	execDir := filepath.Dir(execPath)
	return filepath.Join(execDir, "lisp")
}
