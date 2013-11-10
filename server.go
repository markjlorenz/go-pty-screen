package main

import (
  // "fmt"
  // "github.com/kr/pty"
  // "code.google.com/p/go.crypto/ssh/terminal"
  // "io"
  // "os"
  // "os/exec"
  "dapplebeforedawn/share-pty/servers"
)


func main() {
  channel := make(chan []byte)
  go pty_servers.KeyServer(channel)
  pty_servers.ScreenServer(channel)
}

