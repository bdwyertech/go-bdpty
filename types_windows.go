//go:build windows
// +build windows

package bdpty

import (
	"os/exec"

	log "github.com/sirupsen/logrus"

	"github.com/UserExistsError/conpty"
)

type tty struct {
	*conpty.ConPty
	cmd *exec.Cmd
}

func NewTTY(cmd *exec.Cmd) (tty, error) {
	cpty, err := conpty.Start(cmd.String())
	return tty{cpty, cmd}, err
}

func (t *tty) Setsize(size TTYSize) error {
	return t.Resize(int(size.Cols), int(size.Rows))
}

func (t *tty) Cleanup() {
	if err := t.Close(); err != nil {
		log.Warnf("failed to close spawned tty gracefully: %s", err)
	}
}
