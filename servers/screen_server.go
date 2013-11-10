package pty_servers

import (
  "net"
)

func ScreenServer(channel chan []byte){
  const MAX_CLIENTS = 100
  key_server, err := net.Listen("tcp", ":2001")
  if err != nil { }

  connections := make([]net.Conn, 0, MAX_CLIENTS)
  go func() {
    for {
      conn, err := key_server.Accept()
      if err != nil { }
      connections = append(connections, conn)
    }
  }()

  for {
    bytes := <-channel
    for _, conn := range connections {
      conn.Write(bytes)
    }
  }

}
