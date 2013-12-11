package main

import (
  "fmt"
  "io"
  "os"
  "net"
  "dapplebeforedawn/share-pty/options"
)

func main() {
  opts := options.Client{}
  opts.Parse()

  key_conn, err     := net.Dial("tcp", opts.ServerIP+":"+opts.KeyPort)
  if err != nil { panic(err) }

  screen_conn, err  := net.Dial("tcp", opts.ServerIP+":"+opts.ScreenPort)
  if err != nil { panic(err) }

  go io.Copy(key_conn, os.Stdin)
  io.Copy(os.Stdout, screen_conn)

  fmt.Println("type 'reset' to restore your terminal")
}
