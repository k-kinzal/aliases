package posix

import (
	"fmt"
	"io"
	"os"
	"os/exec"
	"os/signal"
	"strings"
	"syscall"

	"github.com/kr/pty"
	"golang.org/x/crypto/ssh/terminal"
)

type Cmd struct {
	*exec.Cmd
}

func (sh *Cmd) Run() error {
	sh.Cmd.Env = os.Environ()
	sh.Cmd.Stdin = os.Stdin
	sh.Cmd.Stdout = os.Stdout
	sh.Cmd.Stderr = os.Stderr

	if oldState, err := terminal.MakeRaw(int(os.Stdin.Fd())); err == nil {
		ptmx, err := pty.Start(sh.Cmd)
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

	return sh.Cmd.Run()
}

func (sh *Cmd) String() string {
	return strings.Join(sh.Cmd.Args, " ")
}

func Command(name string, arg ...string) *Cmd {
	cmd := new(Cmd)
	cmd.Cmd = exec.Command(name, arg...)

	return cmd
}
