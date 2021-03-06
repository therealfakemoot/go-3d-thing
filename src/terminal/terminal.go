package terminal

import (
  "os"
  "fmt"
  "bufio"
  "syscall"
  "unsafe"
)


type WinSize struct {
  rows    uint16
  cols    uint16
  xpixels uint16
  ypixels uint16
}

type Terminal struct {
  width, height int
  buffer bufio.Writer
}

func NewTerminal() Terminal {
  term := Terminal {
    width: 16,
    height: 16,
    buffer: *bufio.NewWriterSize(os.Stdout, 4096),
  }
  term.UpdateSize()
  return term
}

// Update our size to match the real TTY session
func (t *Terminal) UpdateSize() {
  var winSize WinSize
  syscall.Syscall(
    syscall.SYS_IOCTL,
    os.Stdin.Fd(),
    syscall.TIOCGWINSZ,
    uintptr(unsafe.Pointer(&winSize)),
  )

  t.width = int(winSize.cols)
  t.height = int(winSize.rows)
}

func (t *Terminal) Size() (int, int) {
  return t.width, t.height
}

func (t *Terminal) Width() int {
  return t.width
}

func (t *Terminal) Height() int {
  return t.height
}

func (t *Terminal) Draw(callback func()) {
  callback()
  t.Flush()
}

func (t *Terminal) Flush() {
  t.buffer.Flush()
}

func (t *Terminal) Write(format string, a ...interface{}) {
  fmt.Fprintf(&t.buffer, format, a...)
}

func (t *Terminal) WriteRune(char rune) {
  fmt.Fprint(&t.buffer, string(char))
}
