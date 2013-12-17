package pty_interface

import (
  "fmt"
  "github.com/dapplebeforedawn/pty"
  "os/exec"
  "os"
  "io"
)

const READSIZE = 1024

type PtyInterface struct {
  pty         *os.File
  command     *exec.Cmd
  LogWriter   io.Writer
  command_name string
  rows         uint16
  cols         uint16
  in_chan      chan []byte
  out_chan     chan []byte
}

func NewPty(command string, rows uint16, cols uint16, in_chan, out_chan chan []byte) (*PtyInterface){
  c := exec.Command(command)
  f, err := pty.Start(c)
  if err != nil { panic(err) }

  pty.Setsize( f, rows, cols )

  return &PtyInterface{
    pty:           f,
    command:       c,
    command_name:  command,
    LogWriter:     os.Stderr,
    rows:          rows,
    cols:          cols,
    in_chan:       in_chan,
    out_chan:      out_chan,
  }
}

func (pty_interface *PtyInterface) Start() {
  go func(){
    for bytes := range pty_interface.in_chan {
      fmt.Fprint(pty_interface.LogWriter, bytes)
      pty_interface.pty.Write(bytes)
    }
  }()

  go func(){
    for {
      bytes     := make([]byte, READSIZE)
      read, err := pty_interface.pty.Read(bytes)
      if (err != nil) {
        close(pty_interface.out_chan)
        return
      }
      pty_interface.out_chan <- bytes[:read]
    }
  }()

  pty_interface.command.Wait()
}

