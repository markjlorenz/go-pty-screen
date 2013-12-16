package pty_servers

import (
  "net"
  "bytes"
  "strconv"
)

type ScreenServer struct {
  log_file   bytes.Buffer
  server     net.Listener
  Port       int
}

func NewScreenServer() (ss *ScreenServer) {
  ss = new(ScreenServer)

  var log_buffer bytes.Buffer
  ss.log_file = log_buffer

  return
}

func (ss *ScreenServer) Listen (port int, channel chan []byte){
  const MAX_CLIENTS = 100
  var err error
  port_string := strconv.Itoa(port)
  ss.server, err = net.Listen("tcp", ":"+port_string)
  if err != nil { panic(err) }

  ss.Port = ss.server.Addr().(*net.TCPAddr).Port

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
