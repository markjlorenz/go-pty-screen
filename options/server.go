package options

import (
  "flag"
  "os"
  "code.google.com/p/go.crypto/ssh/terminal"
)

type Server struct {
  Rows uint16
  Cols uint16
  Port int
  App  string
}

func (s *Server) Parse() {
  tty, _        := os.Open("/dev/tty")
  tty_fd        := int( tty.Fd() )

  cols, rows, _ := terminal.GetSize( tty_fd )
  port          := 2000

  flag.IntVar(&rows, "rows", rows, "terminal rows (defaults to rows of current terminal)")
  flag.IntVar(&cols, "cols", cols, "terminal columns (defaults to columns of current terminal)")
  flag.IntVar(&port, "port", port, "port to run the server on")

  flag.Parse()

  s.Rows = uint16(rows)
  s.Cols = uint16(cols)
  s.Port = port
  s.App  = flag.Arg(0)
}
