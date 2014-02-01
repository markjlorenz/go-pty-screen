package main

import (
  "dapplebeforedawn/share-pty/options"
  "dapplebeforedawn/share-pty/clients"
  "dapplebeforedawn/share-pty/zeroconf"
  "code.google.com/p/goncurses"
)

func main() {
  opts := options.Client{}
  opts.Parse()

  if (opts.ServerIP == "") {
    zc := zeroconf.NewClient("_goptyscreen._tcp.")
    zc.Dial()
    opts.ServerIP = zc.Host
    opts.Port     = zc.Port
  }

  defer goncurses.End()
  _, err := goncurses.Init()
  if err != nil { panic(err) }

  goncurses.Cursor(0) // no cursor please
  goncurses.StartColor()

  list := pty_client.NewList(opts.ServerIP, opts.Port)
  list.Fetch()
  key_port, screen_port := list.GetSelection()
  goncurses.End()

  pty_client.Connect(opts.ServerIP, key_port, screen_port)
}
