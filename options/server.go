package options

import (
  "flag"
  "os"
  "code.google.com/p/go.crypto/ssh/terminal"
)

type Server struct {
  Rows         uint16
  Cols         uint16
  KeyPort      int
  ScreenPort   int
  App          string
}

func (s *Server) Parse() {
  tty, _        := os.Open("/dev/tty")
  tty_fd        := int( tty.Fd() )

  cols, rows, _ := terminal.GetSize( tty_fd )
  key_port      := 2000
  screen_port   := 2001

  flag.IntVar(&rows,        "rows",        rows,         "terminal rows (defaults to rows of current terminal)")
  flag.IntVar(&cols,        "cols",        cols,         "terminal columns (defaults to columns of current terminal)")
  flag.IntVar(&key_port,    "key_port",    key_port,     "port to run the key server on")
  flag.IntVar(&screen_port, "screen_port", screen_port,  "port to run the screen server on")

  flag.Parse()

  s.Rows       = uint16(rows)
  s.Cols       = uint16(cols)
  s.KeyPort    = key_port
  s.ScreenPort = screen_port
  s.App        = flag.Arg(0)
}

