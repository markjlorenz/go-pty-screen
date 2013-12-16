package main

import (
  "dapplebeforedawn/share-pty/options"
  "dapplebeforedawn/share-pty/clients"
  "code.google.com/p/goncurses"
)

func main() {
  opts := options.Client{}
  opts.Parse()

  // connect the list and it's backend
  //  -> list draws
  //  -> list selection made
  //  -> backend sarts client

  defer goncurses.End()
  _, err := goncurses.Init()
  if err != nil { panic(err) }

  goncurses.Cursor(0) // high visibilty cursor
  goncurses.StartColor()

  list := pty_client.NewList(opts.ServerIP, opts.Port)
  list.Fetch()
  list.GetSelection()
}
