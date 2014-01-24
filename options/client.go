package options

import (
  "flag"
  "os"
)

type Client struct {
  Port      int
  ServerIP  string
}

func (c *Client) Parse() {
  port       := 2000

  flag.IntVar(&c.Port, "port", port, "port the supervisor is on")
  flag.Parse()
  c.ServerIP = flag.Arg(0)

  // for the tunnel to pickup and read for easy client > server redirection
  ip_filename := "/tmp/go-pty-tunnel~ip"
  ip_file, _  := os.Create(ip_filename)
  ip_file.WriteString(c.ServerIP)
}

