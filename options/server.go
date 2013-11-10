package options

import (
  "flag"
  "os"
  "code.google.com/p/go.crypto/ssh/terminal"
)

type Server struct {
  Rows          uint16
  Cols          uint16
  KeyPort      string
  ScreenPort   string
  App           string
}

type Client struct {
}

func (s *Server) Parse() {
  tty, _        := os.Open("/dev/tty")
  tty_fd        := int( tty.Fd() )

  rows, cols, _ := terminal.GetSize( tty_fd )
  key_port      := "2000"
  screen_port   := "2001"

  flag.IntVar(   &rows,         "rows",        rows,         "terminal rows (defaults to rows of current terminal)")
  flag.IntVar(   &cols,         "cols",        cols,         "terminal columns (defaults to columns of current terminal)")
  flag.StringVar(&s.KeyPort,    "key_port",    key_port,     "port to run the key server on")
  flag.StringVar(&s.ScreenPort, "screen_port", screen_port,  "port to run the screen server on")

  flag.Parse()

  s.Rows = uint16(rows)
  s.Cols = uint16(cols)
  s.App  = flag.Arg(0)
}

