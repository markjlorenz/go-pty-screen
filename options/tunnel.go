package options

import (
  "flag"
)

type Tunnel struct {
  Port      string
  ServerIP  string
  PublicId  string
}

func (c *Tunnel) Parse() {
  port := "2001"
  host := ""
  id   := "nooneelseistunnelingonthisnetwork"

  flag.StringVar(&c.Port, "port", port, "the tunnel port")
  flag.StringVar(&c.ServerIP, "ip", host, "the ip you're connecting to")
  flag.StringVar(&c.PublicId, "id", id, "a unique id to you and your peer.  So you bonjour to the right peer.")
  flag.Parse()
}

