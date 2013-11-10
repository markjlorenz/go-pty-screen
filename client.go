package main

import (
  "fmt"
  "io"
  "os"
  "net"
  "code.google.com/p/go.crypto/ssh/terminal"
)

func main() {
  key_conn, err     := net.Dial("tcp", "localhost:2000")
  if err != nil { }

  screen_conn, err  := net.Dial("tcp", "localhost:2001")
  if err != nil { }

  tty, _       := os.Open("/dev/tty")
  tty_fd       := int( tty.Fd() )
  tty_mode, _  := terminal.MakeRaw( tty_fd )

  go io.Copy(key_conn, os.Stdin)
  io.Copy(os.Stdout, screen_conn)

  fmt.Println("type 'reset' to restore your terminal")
  terminal.Restore( tty_fd, tty_mode )  // why u no work?
}
