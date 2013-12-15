package pty_servers

import (
  "net"
  "net/http"
  "io"
  "bufio"
  "strconv"

  "fmt"
)

type PtyShare struct {
  key_server    *KeyServer
  screen_server *ScreenServer
}

type Supervisor struct {
  pty_shares map[string]*PtyShare
}

func NewSupervisor() (*Supervisor){
  return new(Supervisor)
}

func (visor *Supervisor) Listen(port uint16) {
  port_string := strconv.Itoa(int(port))
  server, err := net.Listen("tcp", ":"+port_string)
  if err != nil { panic(err) }

  for {
    conn, err := server.Accept()
    if err != nil { panic(err) }
    go visor.process_request(conn)
  }
}

func (visor *Supervisor) process_request(conn net.Conn) {
  reader   := bufio.NewReader(conn)
  req, err := http.ReadRequest(reader)
  if err == io.EOF { return }
  if err != nil    { panic(err) }

  responder := visor.route(req)
  responder(conn)
}

func (visor *Supervisor) route(req *http.Request) (func(conn net.Conn)) {
  fmt.Println(req.URL)
  switch route_string := (req.Method + req.URL.Path)
  route_string {
  case "GET/servers":
    return visor.serve_list
  case "POST/servers":
    return visor.serve_create
  default:
    return visor.four_oh_four
  }
}

func (visor *Supervisor) serve_list(conn net.Conn) {
  conn.Write([]byte(visor.http_response(200, "OK", "HI")))
  conn.Close()
}

func (visor *Supervisor) serve_create(conn net.Conn) {
}

func (visor *Supervisor) four_oh_four(conn net.Conn) {
  conn.Write([]byte(visor.http_response(404, "Not Found", "I'm sorry, I can't do that.")))
  conn.Close()
}

func (visor *Supervisor) http_response(status uint, message string, body string) (resp string) {
  resp = "HTTP/1.1 "+strconv.Itoa(int(status))+" "+message+"\r\n\r\n"
  resp += body
  resp += "\r\n"
  return
}
