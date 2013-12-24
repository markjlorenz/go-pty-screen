package pty_servers

import (
  "net"
  "net/http"
  "io"
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
  create_chan chan PtyShare
  delete_chan chan string
}

type instruction struct {
  alias     string
  command   string
  rows      int
  cols      int
}

func NewSupervisor(creates chan PtyShare, deletes chan string) (visor *Supervisor){
  visor = new(Supervisor)
  visor.pty_shares  = make(map[string]*PtyShare)
  visor.create_chan = creates
  visor.delete_chan = deletes
  return
}

func (visor *Supervisor) Listen(port int, ready chan int) {
  port_string := strconv.Itoa(port)
  server, err := net.Listen("tcp", ":"+port_string)
  if err != nil { panic(err) }

  ready <- port
  close(ready)
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
  switch route_string := (req.Method + req.URL.Path)
  route_string {
  case "GET/servers":
    return visor.serve_list()
  case "POST/servers":
    instructions := visor.parse_instructions(req.Body)
    for _, instr := range instructions {
      go visor.new_server(instr.alias, instr.command, instr.rows, instr.cols)
    }
    return visor.serve_create()
  default:
    return visor.four_oh_four()
  }
}

func (visor *Supervisor) new_server(alias string, command string, rows int, cols int){
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

  share := PtyShare{
    KeyServer:     key_server,
    ScreenServer:  screen_server,
    Command:       command,
    Alias:         alias,
  }

  visor.pty_shares[alias] = &share
  visor.create_chan <- share

  pty.Start()

  // if you get here this server is dead, you can remove it from the list
  delete(visor.pty_shares, alias)
  visor.delete_chan <- share.Alias
}

func (visor *Supervisor) parse_instructions(instructions io.Reader) (instr_set []instruction){
  scanner := bufio.NewScanner(instructions)
  for scanner.Scan() {
    fields  := strings.Fields(scanner.Text())
    rows, _ := strconv.Atoi(fields[2])
    cols, _ := strconv.Atoi(fields[3])
    instr_set = append(instr_set, instruction{
      alias:    fields[0],
      command:  fields[1],
      cols:     cols,
      rows:     rows,
    })
  }
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
