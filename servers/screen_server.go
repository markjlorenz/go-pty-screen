package pty_servers

import (
  "net"
)

func ScreenServer(channel chan []byte){
  const MAX_CLIENTS = 100
  server, err := net.Listen("tcp", ":2001")
  if err != nil { }

  connections := make([]net.Conn, 0, MAX_CLIENTS)
  go accept_connections(server, &connections)

  for {
    bytes := <-channel
    for _, conn := range connections {
      conn.Write(bytes)
    }
  }

}

func accept_connections(server net.Listener, connections *[]net.Conn){
  for {
    conn, err := server.Accept()
    if err != nil { }
    *connections = append(*connections, conn)
  }
}
