package options

import (
  "flag"
)

type Client struct {
  Port      int
  ServerIP  string
}

func (c *Client) Parse() {
  port       := 2000
  c.ServerIP = "localhost"

  flag.IntVar(&c.Port, "port", port, "port the supervisor is on")
  flag.Parse()
  c.ServerIP = flag.Arg(0)
}

