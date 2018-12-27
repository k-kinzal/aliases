package posix

import (
	"fmt"
	"io"
	"os"
	"os/exec"
	"os/signal"
	"syscall"

	"github.com/kr/pty"
	"golang.org/x/crypto/ssh/terminal"
)

func Run(cmd exec.Cmd) error {
	sh := exec.Command("sh", "-c", String(cmd))
	sh.Env = os.Environ()
	sh.Stdin = os.Stdin
	sh.Stdout = os.Stdout
	sh.Stderr = os.Stderr

	if oldState, err := terminal.MakeRaw(int(os.Stdin.Fd())); err == nil {
		ptmx, err := pty.Start(sh)
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

	return sh.Run()
}
