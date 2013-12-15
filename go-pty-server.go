package main

import (
  "dapplebeforedawn/share-pty/servers"
  // "dapplebeforedawn/share-pty/pty_interface"
  "dapplebeforedawn/share-pty/options"

  "fmt"
)

func main() {
  opts := options.Server{}
  opts.Parse()

  // key_channel     := make(chan []byte)
  // screen_channel  := make(chan []byte)

  // screen_server := pty_servers.NewScreenServer()
  // key_server    := pty_servers.NewKeyServer()

  // go key_server.Listen(opts.KeyPort, key_channel)
  // go screen_server.Listen(opts.ScreenPort, screen_channel)

  supervisor := pty_servers.NewSupervisor()
  supervisor.Listen(opts.KeyPort)

  fmt.Println("Server started running: " + opts.App)
  fmt.Println("Clients can connect with the command: go-pty-client --key_port="+string(opts.KeyPort)+" --screen_port="+string(opts.ScreenPort)+" --host=<ip address>")
  // pty_interface.Pty(opts.App, opts.Rows, opts.Cols, key_channel, screen_channel)
}
