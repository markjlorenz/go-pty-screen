package pty_servers

import (
  "net"
  "net/http"
  "io"
  "io/ioutil"
  "os"
  "bufio"
  "strconv"
  "strings"
  "time"
  "dapplebeforedawn/share-pty/pty_interface"
)

type PtyShare struct {
  KeyServer    *KeyServer
  ScreenServer *ScreenServer
  Command      string
  Alias        string
}

type Supervisor struct {
  pty_shares  map[string]*PtyShare
  update_chan chan PtyShare
}

func NewSupervisor(updates chan PtyShare) (visor *Supervisor){
  visor = new(Supervisor)
  visor.pty_shares  = make(map[string]*PtyShare)
  visor.update_chan = updates
  return
}

func (visor *Supervisor) Listen(port int) {
  port_string := strconv.Itoa(port)
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

  response := visor.route(req)
  conn.Write([]byte(response))
  conn.Close()
}

func (visor *Supervisor) route(req *http.Request) (string){
  // fmt.Println(req.URL)
  switch route_string := (req.Method + req.URL.Path)
  route_string {
  case "GET/servers":
    return visor.serve_list()
  case "POST/servers":
    alias, command, cols, rows := visor.parse_instructions(req.Body)
    // fmt.Println("Spinning Up: ", alias, command)
    visor.new_server(alias, command, cols, rows)
    return visor.serve_create()
  default:
    return visor.four_oh_four()
  }
}

func (visor *Supervisor) new_server(alias string, command string, cols int, rows int){
  // find 2 free ports
  // start key server, screen servers and pty
  // stash the new PtyShare in pty_shares under the alias name
  key_channel     := make(chan []byte)
  screen_channel  := make(chan []byte)

  screen_server := NewScreenServer()
  key_server    := NewKeyServer()

  // let the OS assign a port
  go key_server.Listen(0, key_channel)
  go screen_server.Listen(0, screen_channel)
  pty := pty_interface.NewPty(command, uint16(rows), uint16(cols), key_channel, screen_channel)

  timestamp      := time.Now().Unix()
  temp_file, err := os.Create("/tmp/"+strconv.Itoa(int(timestamp))+"~go-pty-screen~"+alias)
  if (err != nil) { panic(err) }
  pty.LogWriter   = temp_file
  go pty.Start()

  share := PtyShare{}
  share.KeyServer    = key_server
  share.ScreenServer = screen_server
  share.Command      = command
  share.Alias        = alias

  visor.pty_shares[alias] = &share
  visor.update_chan <- share
}

func (visor *Supervisor) parse_instructions(instructions io.Reader) (alias string, command string, cols int, rows int){

  raw_cmd, err := ioutil.ReadAll(instructions)
  if err != nil { panic(err) }
  fields := strings.Fields(string(raw_cmd))
  alias   = fields[0]
  command = fields[1]
  cols, _ = strconv.Atoi(fields[2])
  rows, _ = strconv.Atoi(fields[3])
  return
}

func (visor *Supervisor) serve_list() (string){
  response := ""
  for alias, pty_share := range visor.pty_shares {
    key_port    := pty_share.KeyServer.Port
    screen_port := pty_share.ScreenServer.Port

    response += alias+" "+pty_share.Command+" "+strconv.Itoa(key_port)+" "+strconv.Itoa(screen_port)+"\r\n"
  }
  return visor.http_response(200, "OK", response)
}

func (visor *Supervisor) serve_create() (string){
  return visor.http_response(201, "CREATED", "")
}

func (visor *Supervisor) four_oh_four() (string){
  return visor.http_response(404, "NOT FOUND", "I'm sorry, I can't do that.")
}

func (visor *Supervisor) http_response(status uint, message string, body string) (resp string){
  resp = "HTTP/1.1 "+strconv.Itoa(int(status))+" "+message+"\r\n\r\n"
  resp += body
  resp += "\r\n"
  return
}
