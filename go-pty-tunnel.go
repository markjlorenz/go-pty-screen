package main

import (
  "net"
  "os"
  "strconv"
  "dapplebeforedawn/share-pty/options"
  "dapplebeforedawn/share-pty/zeroconf"
  "github.com/nu7hatch/gouuid"
)

var go_pty_service = "_goptytunnel._tcp."
var public_id_key  = "public_id"
var private_id_key = "private_id"
var private_id = func() string {
  u4, _ := uuid.NewV4()
  return u4.String()
}()

func main() {
  opts := options.Tunnel{}
  opts.Parse()

  stdin_chan := StdinChannel()
  stdout_chan := StdoutChannel(opts.Port, opts.PublicId)

  for {
    select {
    case data, ok := <-stdin_chan:
      sender := NewSender(opts.ServerIP, opts.Port, opts.PublicId)
      sender.SendData(data, !ok)
      if (!ok) { return }
    case data, ok := <-stdout_chan:
      ReceiveData(data)
      if (!ok) { return }
    }
  }

}

const READSIZE = 1024

type Sender struct {
  conn net.Conn
  host string
  port string
}

func NewSender(host string, port string, public_id string) *Sender {
  // give priority to host/port style
  if host != "" {
    return &Sender{
      host: host,
      port: port,
    }
  }

  zc := zeroconf.NewClient(go_pty_service)

  matcher := func(txtRecords map[string]string) bool {
    return txtRecords[public_id_key]  == public_id &&
           txtRecords[private_id_key] != private_id  // don't connect to yourself
  }

  for zc.Host == "" { // loop until a connection is available
    zc.DialWhenMatch(matcher)
  }

  zcStringPort := strconv.Itoa(zc.Port)
  return &Sender{
    host: zc.Host,
    port: zcStringPort,
  }

}

func (sender *Sender) SendData(data []byte, closed bool) {
  if sender.conn == nil {
    var cantConnect error
    sender.conn, cantConnect = net.Dial("tcp", sender.host+":"+sender.port)
    if cantConnect !=nil { panic(cantConnect) }
  }
  sender.conn.Write(data)

  if(closed) { sender.conn.Close() }
}

func ReceiveData(data []byte) {
  os.Stdout.Write(data)
}

func StdinChannel() (channel chan []byte){
  channel = make(chan []byte)
  go func(){
    for {
      bytes     := make([]byte, READSIZE)
      read, err := os.Stdin.Read(bytes)
      if err != nil { close(channel); break }
      channel <- bytes[:read]
    }
  }()
  return
}

func StdoutChannel(port string, public_id string) (channel chan []byte){
  channel = make(chan []byte)

  // start listening
  listener, unix_err := net.Listen("tcp", ":"+port)
  if (unix_err != nil) { panic(unix_err) }

  server := zeroconf.NewServer(go_pty_service)
  server.TxtRecords = map[string]string{
    public_id_key:  public_id,
    private_id_key: private_id,
  }

  string_port, _ := strconv.Atoi(port)
  bonjour, err := server.StartAnnounce(string_port)
  if err != nil { panic(err) }

  go func(){
    for {
      conn, _ := listener.Accept()
      go func(){
        for {
          bytes     := make([]byte, READSIZE)
          read, err := conn.Read(bytes)
          if err != nil { close(channel); break }
          channel <- bytes[:read]
        }
      }()
    }
    defer bonjour.Release()
  }()

  return
}
