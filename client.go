package main

import (
  // "fmt"
  // "github.com/kr/pty"
  // "code.google.com/p/go.crypto/ssh/terminal"
  "io"
  "os"
  // "os/exec"
  "net"
)

// read from the screen socket and copy to stdout

// read from stdin and copy to the key socket
func main() {
  key_conn, err := net.Dial("tcp", "localhost:2000")
  if err != nil { }

  screen_conn, err := net.Dial("tcp", "localhost:2001")
  if err != nil { }

  go io.Copy(key_conn, os.Stdin)
  io.Copy(os.Stdout, screen_conn)
}
