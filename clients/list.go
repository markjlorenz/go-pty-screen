package pty_client

import (
  "strings"
  "strconv"
  "bufio"
  "net/http"
  "dapplebeforedawn/share-pty/servers"
  "dapplebeforedawn/share-pty/views/client"
)

type List struct {
  supervisor_host string
  supervisor_port int
  view            *client_views.List
}

func NewList(supervisor_host string, supervisor_port int) (list *List) {
  list = new(List)
  list.supervisor_host = supervisor_host
  list.supervisor_port = supervisor_port
  list.view            = client_views.NewList()
  return
}

func (list *List) Fetch() {
  url := "http://"+list.supervisor_host+":"+strconv.Itoa(list.supervisor_port)+"/servers"
  resp, err := http.Get(url)
  if err != nil { panic(err) }

  scanner := bufio.NewScanner(resp.Body)
  scanner.Split(bufio.ScanLines)

  for scanner.Scan() {
    text := strings.TrimSpace(scanner.Text())
    if(text == ""){ continue }
    fields             := strings.Fields(text)

    key_port, _        := strconv.Atoi(fields[2])
    screen_port, _     := strconv.Atoi(fields[3])
    key_server_info    := &pty_servers.KeyServer{Port: key_port}
    screen_server_info := &pty_servers.ScreenServer{Port: screen_port}
    pty_share          := pty_servers.PtyShare{
                            Alias:        fields[0],
                            Command:      fields[1],
                            KeyServer:    key_server_info,
                            ScreenServer: screen_server_info,
                          }
    list.view.AddItem(pty_share)
  }
}

func (list *List) GetSelection() (key_port, screen_port int){
  pty_share := list.view.SelectRow()
  return pty_share.KeyServer.Port, pty_share.ScreenServer.Port
}
