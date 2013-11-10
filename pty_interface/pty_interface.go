package pty_interface

import (
  "fmt"
  // "github.com/kr/pty"
  "github.com/dapplebeforedawn/pty"
  "os/exec"
)

const READSIZE = 1024

func Pty(command string, in_chan chan []byte, out_chan chan []byte) {

  c := exec.Command(command)
  f, err := pty.Start(c)
  if err != nil { panic(err) }

  pty.Setsize( f, 20, 40 )

  go func(){
    for bytes := range in_chan {
      fmt.Print( string(bytes) )
      f.Write(bytes)
    }
  }()

  go func(){
    for {
      bytes   := make([]byte, READSIZE)
      read, _ := f.Read(bytes)
      out_chan <- bytes[:read]
    }
  }()

  c.Wait()
}

