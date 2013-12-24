package options

import (
  "flag"
  "os"
)

type Server struct {
  Port       int
  RCFilename string
}

func (s *Server) Parse() {
  rc_filename := os.Getenv("HOME")+"/.go-pty-rc"
  port        := 2000

  flag.IntVar(&port, "port", port, "port to run the server on")

  flag.StringVar(&rc_filename, "config-file", rc_filename,
    "a config file to run on server startup.")

  flag.Parse()

  s.Port       = port
  s.RCFilename = rc_filename
}
