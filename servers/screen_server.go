package pty_servers

import (
  "net"
  // "os"
  // "io"
  "bytes"
)

type ScreenServer struct {
  log_file bytes.Buffer
}

func NewScreenServer() (ss *ScreenServer) {
  ss = new(ScreenServer)

  var log_buffer bytes.Buffer
  ss.log_file = log_buffer

  return
}

func (ss *ScreenServer) Listen (channel chan []byte){
  const MAX_CLIENTS = 100
  server, err := net.Listen("tcp", ":2001")
  if err != nil { panic(err) }

  connections := make([]net.Conn, 0, MAX_CLIENTS)
  go ss.accept_connections(server, &connections)

  for {
    screen_bytes := <-channel
    ss.log_file.Write(screen_bytes)
    for _, conn := range connections {
      conn.Write(screen_bytes)
    }
  }

}

func (ss *ScreenServer) accept_connections(server net.Listener, connections *[]net.Conn){
  for {
    conn, err := server.Accept()
    if err != nil { panic(err) }
    conn.Write( ss.log_file.Bytes() )
    *connections = append(*connections, conn)
  }
}
