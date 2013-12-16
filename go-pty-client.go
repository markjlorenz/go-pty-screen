package main

import (
  "dapplebeforedawn/share-pty/options"
  "dapplebeforedawn/share-pty/clients"
  "code.google.com/p/goncurses"
)

func main() {
  opts := options.Client{}
  opts.Parse()

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
