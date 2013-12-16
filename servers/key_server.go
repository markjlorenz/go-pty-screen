package pty_servers

import (
  "net"
  "strconv"
)

const READSIZE = 1024

type KeyServer struct {
  Port int
}

func NewKeyServer() (ks *KeyServer) {
  return new(KeyServer)
}

func (ks *KeyServer) Listen(port int, channel chan []byte) {
  port_string := strconv.Itoa(port)
  server, err := net.Listen("tcp", ":"+port_string)
  if err != nil { panic(err) }

  ks.Port = server.Addr().(*net.TCPAddr).Port

  for {
    conn, err := server.Accept()
    if err != nil { panic(err) }
    go ks.connection_to_channel(conn, channel)
  }
}

func (ks *KeyServer) connection_to_channel(conn net.Conn, channel chan []byte) {
  for {
    bytes     := make([]byte, READSIZE)
    read, err := conn.Read(bytes)
    if err != nil { return }
    bytes = bytes[:read]
    channel <- bytes
  }
}


