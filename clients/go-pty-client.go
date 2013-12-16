package pty_client

import (
  "fmt"
  "io"
  "os"
  "net"
  "strconv"
  "github.com/dapplebeforedawn/go.crypto/ssh/terminal"
)

func Connect(host string, key_port, screen_port int) {
  key_conn, err     := net.Dial("tcp", host+":"+ strconv.Itoa(key_port))
  if err != nil { panic(err) }

  screen_conn, err  := net.Dial("tcp", host+":"+strconv.Itoa(screen_port))
  if err != nil { panic(err) }

  tty, _       := os.Open("/dev/tty")
  tty_fd       := int( tty.Fd() )
  tty_mode, _  := terminal.MakeRaw( tty_fd )

  go io.Copy(key_conn, os.Stdin)
  io.Copy(os.Stdout, screen_conn)

  fmt.Println("type 'reset' to restore your terminal")
  terminal.Restore( tty_fd, tty_mode )  // why u no work?
}
