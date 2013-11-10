package pty_servers

import (
  "net"
)

func KeyServer(channel chan []byte) {
  key_server, err := net.Listen("tcp", ":2000")
  if err != nil { }

  for {
    conn, err := key_server.Accept()
    if err != nil { }

    go func(){
      for {
        bytes := make([]byte, 4096)
        conn.Read(bytes)
        channel <- bytes
      }
    }()
  }
}


