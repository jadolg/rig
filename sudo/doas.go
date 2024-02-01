package sudo

import (
	"github.com/alessio/shellescape"
	"github.com/k0sproject/rig/exec"
)

// Sudo is a DecorateFunc that will wrap the given command in a sudo call.
func Sudo(cmd string) string {
	return `sudo -n -- "${SHELL-sh}" -c ` + shellescape.Quote(cmd)
}

// RegisterSudo registers a sudo DecorateFunc with the given repository.
func RegisterSudo(repository *Repository) {
	repository.Register(func(c exec.SimpleRunner) exec.DecorateFunc {
		if c.IsWindows() {
			return nil
		}
		if c.Exec(Sudo("true")) != nil {
			return nil
		}
		return Sudo
	})
}
