package pty_servers

import (
  "net"
  "os"
)

type RCLoader struct {
  rc_file *os.File
  feed    chan int
}

func NewRCLoader(feed chan int, rc_filename string) (rc *RCLoader) {
    rc_file, f_err := os.Open(rc_filename)
    if ( f_err != nil ) { return &RCLoader{ feed: feed, } }
    rc = &RCLoader{
      rc_file: rc_file,
      feed:    feed,
    }
    return
}

func(rc *RCLoader) OnReady() {
    port, _ := <-rc.feed
    visor, err := net.DialTCP("tcp", nil, &net.TCPAddr{Port: port})
    defer visor.Close()
    if ( err != nil ) { panic(err) }

    _, _ = visor.ReadFrom(rc.rc_file)
}
