// +build darwin

package tcell

// The Darwin system is *almost* a real BSD system, but it suffers from
// a brain damaged TTY driver.  This TTY driver does not actually
// wake up in poll() or similar calls, which means that we cannot reliably
// shut down the terminal without resorting to obscene custom C code
// and a dedicated poller thread.
//
// So instead, we do a best effort, and simply try to do the close in the
// background.  Probably this will cause a leak of two goroutines and
// maybe also the file descriptor, meaning that applications on Darwin
// can't reinitialize the screen, but that's probably a very rare behavior,
// and accepting that is the best of some very poor alternative options.
//
// Maybe someday Apple will fix there tty driver, but its been broken for
// a long time (probably forever) so holding one's breath is contraindicated.
//
// NOTE: In this fork of tcell, we fix this issue by using the library
// zyedidia/poller to properly interface with the tty such that when we
// close it, it actually closes

import (
	"io"

	"github.com/zyedidia/poller"
)

func openTty() (io.Reader, io.Writer, error) {
	var out io.Writer
	in, err := poller.Open("/dev/tty", poller.O_RO)
	if err == nil {
		out, err = poller.Open("/dev/tty", poller.O_WO)
	}

	return in, out, err
}

func closeTty(f interface{}) {
	f.(*poller.FD).Close()
}

func (t *tScreen) fd() int {
	return t.out.(*poller.FD).Sysfd()
}
