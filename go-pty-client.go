package main

import (
  "dapplebeforedawn/share-pty/options"
  "dapplebeforedawn/share-pty/clients"
  "dapplebeforedawn/share-pty/zeroconf"
  "code.google.com/p/goncurses"
  "os"
)

func main() {
  opts := options.Client{}
  opts.Parse()

  if (opts.ServerIP == "") {
    zc := zeroconf.NewClient()
    zc.Dial()
    opts.ServerIP = zc.Host
    opts.Port     = zc.Port
  }

  write_ip_file(opts.ServerIP)

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

func write_ip_file(server_ip string) {
  // for the tunnel to pickup and read for easy client > server redirection
  ip_filename := "/tmp/go-pty-tunnel~ip"
  ip_file, _  := os.Create(ip_filename)
  ip_file.WriteString(server_ip)
}

