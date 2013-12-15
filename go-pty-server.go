package main

import (
  "dapplebeforedawn/share-pty/servers"
  "dapplebeforedawn/share-pty/options"
  "dapplebeforedawn/share-pty/views"

  "fmt"
)

func main() {
  fmt.Println("Started")
  opts := options.Server{}
  opts.Parse()

  supervisor := pty_servers.NewSupervisor()
  go supervisor.Listen(opts.Port)

  view := pty_views.NewSupervisor()
  print(view)
}
