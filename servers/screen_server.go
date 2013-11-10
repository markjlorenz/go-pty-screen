package pty_servers

import (
  "net"
  "bytes"
)

type ScreenServer struct {
  log_file     bytes.Buffer
  server       net.Listener
}

func NewScreenServer() (ss *ScreenServer) {
  ss = new(ScreenServer)

  var log_buffer bytes.Buffer
  ss.log_file     = log_buffer

  return
}

func (ss *ScreenServer) Listen (port string, channel chan []byte){
  const MAX_CLIENTS = 100
  var err error
  ss.server, err = net.Listen("tcp", ":"+port)
  if err != nil { panic(err) }

  connections := make([]net.Conn, 0, MAX_CLIENTS)
  go ss.accept_connections(&connections)

  for {
    screen_bytes := <-channel
    ss.log_file.Write(screen_bytes)
    for _, conn := range connections {
      conn.Write(screen_bytes)
    }
  }
}

func (ss *ScreenServer) accept_connections(connections *[]net.Conn){
  for {
    conn, err := ss.server.Accept()
    if err != nil { panic(err) }
    conn.Write( ss.log_file.Bytes() )
    *connections = append(*connections, conn)
  }
}
