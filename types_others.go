//go:build !windows
// +build !windows

package bdpty

import (
	"os"
	"os/exec"

	log "github.com/sirupsen/logrus"

	"github.com/creack/pty"
)

type tty struct {
	*os.File
	f   *os.File
	cmd *exec.Cmd
}

func NewTTY(cmd *exec.Cmd) (tty, error) {
	_tty, err := pty.Start(cmd)
	return tty{_tty, _tty, cmd}, err
}

func (t *tty) Setsize(size TTYSize) error {
	return pty.Setsize(t.f, &pty.Winsize{
		Rows: size.Rows,
		Cols: size.Cols,
	})
}

func (t *tty) Cleanup() {
	if err := t.cmd.Process.Kill(); err != nil {
		log.Warnf("failed to kill process: %s", err)
	}
	if _, err := t.cmd.Process.Wait(); err != nil {
		log.Warnf("failed to wait for process to exit: %s", err)
	}
	if err := t.Close(); err != nil {
		log.Warnf("failed to close spawned tty gracefully: %s", err)
	}
}
