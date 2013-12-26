package main

import (
  "net"
  "os"
  "dapplebeforedawn/share-pty/options"
)

func main() {
  opts := options.Tunnel{}
  opts.Parse()

  sender := &Sender{
    host: opts.ServerIP,
    port: opts.Port,
  }

  stdin_chan := StdinChannel()
  stdout_chan := StdoutChannel(opts.Port)

  for {
    select {
    case data, ok := <-stdin_chan:
      sender.SendData(data, !ok)
      if (!ok) { return }
    case data, ok := <-stdout_chan:
      ReceiveData(data)
      if (!ok) { return }
    }
  }

}

const READSIZE = 1024

type Sender struct {
  conn net.Conn
  host string
  port string
}

func (sender *Sender) SendData(data []byte, closed bool) {
  for sender.conn == nil {
    sender.conn, _ = net.Dial("tcp", sender.host+":"+sender.port)
  }
  sender.conn.Write(data)

  if(closed) { sender.conn.Close() }
}

func ReceiveData(data []byte) {
  os.Stdout.Write(data)
}

func StdinChannel() (channel chan []byte){
  channel = make(chan []byte)
  go func(){
    for {
      bytes     := make([]byte, READSIZE)
      read, err := os.Stdin.Read(bytes)
      if err != nil { close(channel); break }
      channel <- bytes[:read]
    }
  }()
  return
}

func StdoutChannel(port string) (channel chan []byte){
  channel = make(chan []byte)

  listener, unix_err := net.Listen("tcp", ":"+port)
  if (unix_err != nil) { panic(unix_err) }
  go func(){
    for {
      conn, _ := listener.Accept()
      go func(){
        for {
          bytes     := make([]byte, READSIZE)
          read, err := conn.Read(bytes)
          if err != nil { close(channel); break }
          channel <- bytes[:read]
        }
      }()
    }
  }()

  return
}
