package main

import (
  "dapplebeforedawn/share-pty/servers"
  "dapplebeforedawn/share-pty/pty_interface"
)


func main() {
  key_channel     := make(chan []byte)
  screen_channel  := make(chan []byte)

  go pty_servers.KeyServer(key_channel)
  go pty_servers.ScreenServer(screen_channel)

  pty_interface.Pty("vim", key_channel, screen_channel)
}

