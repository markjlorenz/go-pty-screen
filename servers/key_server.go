package pty_servers

import (
  "net"
)

const READSIZE = 1024

type KeyServer struct {
}

func NewKeyServer() (ks *KeyServer) {
  return new(KeyServer)
}

func (ks *KeyServer) Listen(port string, channel chan []byte) {
  server, err := net.Listen("tcp", ":"+port)
  if err != nil { }

  for {
    conn, err := server.Accept()
    if err != nil { }
    go ks.connection_to_channel(conn, channel)
  }
}

func (ks *KeyServer) connection_to_channel(conn net.Conn, channel chan []byte) {
  for {
    bytes     := make([]byte, READSIZE)
    read, _   := conn.Read(bytes)
    bytes = bytes[:read]
    channel <- bytes
  }
}


