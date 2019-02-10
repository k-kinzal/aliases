package posix

import (
	"fmt"
	"io"
	"os"
	"os/exec"
	"os/signal"
	"path"
	"strings"
	"syscall"

	"github.com/kr/pty"
	"golang.org/x/crypto/ssh/terminal"
)

// Cmd is the base of the POSIX commands.
type Cmd struct {
	*exec.Cmd
}

// Run command
// If there is a terminal connected to the file descriptor, run it via PTS.
func (cmd *Cmd) Run() error {
	oldState, err := terminal.MakeRaw(int(os.Stdin.Fd()))
	if err != nil {
		return cmd.Cmd.Run()
	}
	ptmx, err := pty.Start(cmd.Cmd)
	if err != nil {
		return err
	}
	defer func() { _ = ptmx.Close() }()

	ch := make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGWINCH)
	go func() {
		for range ch {
			if err := pty.InheritSize(os.Stdin, ptmx); err != nil {
				panic(fmt.Sprintf("error resizing pty: %s", err))
			}
		}
	}()
	ch <- syscall.SIGWINCH

	if err != nil {
		return err
	}
	defer func() {
		_ = terminal.Restore(int(os.Stdin.Fd()), oldState)
	}()

	go func() {
		_, _ = io.Copy(ptmx, os.Stdin)
	}()
	_, _ = io.Copy(os.Stdout, ptmx)

	return nil
}

// String returns command string.
func (cmd *Cmd) String() string {
	return fmt.Sprintf("%s %s", path.Base(cmd.Cmd.Args[0]), strings.Join(cmd.Cmd.Args[1:], " "))
}

// Command creates a new posix.Cmd.
func Command(name string, arg ...string) *Cmd {
	return &Cmd{exec.Command(name, arg...)}
}
