package pty_servers

import (
  "net"
)

const READSIZE = 1024

func KeyServer(channel chan []byte) {
  server, err := net.Listen("tcp", ":2000")
  if err != nil { }

  for {
    conn, err := server.Accept()
    if err != nil { }
    go connection_to_channel(conn, channel)
  }
}

func connection_to_channel(conn net.Conn, channel chan []byte) {
  for {
    bytes     := make([]byte, READSIZE)
    read, _   := conn.Read(bytes)
    bytes = bytes[:read]
    channel <- bytes
  }
}


