package main

import (
  "fmt"
  "io"
  "os"
  "net"
  "dapplebeforedawn/share-pty/options"
  "github.com/dapplebeforedawn/go.crypto/ssh/terminal"
)

func main() {
  opts := options.Client{}
  opts.Parse()

  key_conn, err     := net.Dial("tcp", opts.ServerIP+":"+opts.KeyPort)
  if err != nil { panic(err) }

  screen_conn, err  := net.Dial("tcp", opts.ServerIP+":"+opts.ScreenPort)
  if err != nil { panic(err) }

  tty, _       := os.Open("/dev/tty")
  tty_fd       := int( tty.Fd() )
  tty_mode, _  := terminal.MakeRaw( tty_fd )

  go io.Copy(key_conn, os.Stdin)
  io.Copy(os.Stdout, screen_conn)

  fmt.Println("type 'reset' to restore your terminal")
  terminal.Restore( tty_fd, tty_mode )  // why u no work?
}
