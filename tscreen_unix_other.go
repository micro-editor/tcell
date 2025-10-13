// +build aix dragonfly freebsd linux netbsd openbsd solaris zos

package tcell

import (
	"io"
	"os"
)

func openTty() (io.Reader, io.Writer, error) {
	var out io.Writer
	in, err := os.OpenFile("/dev/tty", os.O_RDONLY, 0)
	if err == nil {
		out, err = os.OpenFile("/dev/tty", os.O_WRONLY, 0)
	}

	return in, out, err
}

func closeTty(f interface{}) {
	f.(*os.File).Close()
}

func (t *tScreen) fd() int {
	return int(t.out.(*os.File).Fd())
}
