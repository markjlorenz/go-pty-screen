package main

import (
  "dapplebeforedawn/share-pty/servers"
  "dapplebeforedawn/share-pty/options"

  "fmt"
)

func main() {
  fmt.Println("Started")
  opts := options.Server{}
  opts.Parse()

  supervisor := pty_servers.NewSupervisor()
  supervisor.Listen(opts.KeyPort)
}
