package options

import (
  "flag"
  "io/ioutil"
  "os"
)

type Tunnel struct {
  Port      string
  ServerIP  string
}

func (c *Tunnel) Parse() {
  port := "2001"

  ip_filename  := "/tmp/go-pty-tunnel~ip"
  ip_file, _   := os.Open(ip_filename)
  server_ip, _ := ioutil.ReadAll(ip_file)

  flag.StringVar(&c.Port, "port", port, "the tunnel port")
  flag.StringVar(&c.ServerIP, "ip", string(server_ip), "the ip you're connecting to")
  flag.Parse()
}

