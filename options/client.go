package options

import (
  "flag"
)

type Client struct {
  KeyPort     string
  ScreenPort  string
  ServerIP    string
}

func (c *Client) Parse() {
  key_port      := "2000"
  screen_port   := "2001"
  c.ServerIP     = "localhost"

  flag.StringVar(&c.KeyPort,    "key_port",    key_port,     "port the key server is on")
  flag.StringVar(&c.ScreenPort, "screen_port", screen_port,  "port the screen server is on")

  flag.Parse()

  c.ServerIP = flag.Arg(0)
}

