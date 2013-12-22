package pty_views

import (
  "code.google.com/p/goncurses"
  "dapplebeforedawn/share-pty/servers"
  "dapplebeforedawn/share-pty/views/supervisor"
  "strings"
  "strconv"
  "bytes"
  "net/http"
)

type Supervisor struct {
  list_window    *supervisor_views.List
  command_window *supervisor_views.Command
}

func NewSupervisor() (supervisor *Supervisor){
  supervisor = new(Supervisor)

  defer goncurses.End()
  _, err := goncurses.Init()
  if err != nil { panic(err) }

  goncurses.Cursor(2) // high visibilty cursor
  goncurses.StartColor()

  supervisor.list_window    = supervisor_views.NewList()
  supervisor.command_window = supervisor_views.NewCommand()

  supervisor.list_window.Refresh()
  supervisor.command_window.Refresh()

  return
}

func (supervisor *Supervisor) Refresh() {
  supervisor.list_window.Refresh()
  supervisor.command_window.Refresh()
}

func (supervisor *Supervisor) CreateFeed(feed chan pty_servers.PtyShare) {
  for update := range feed {
    supervisor.list_window.AddItem(update)
  }
}

func (supervisor *Supervisor) DeleteFeed(feed chan string) {
  for update := range feed {
    supervisor.list_window.RemoveItem(update)
  }
}

func (supervisor *Supervisor) WatchCommands(port int){
  for {
    input := supervisor.command_window.GetInput()

    fields := strings.Fields(input)
    if (len(fields) != 5) {
      supervisor.command_window.FlashError("5 fields required")
      continue
    }

    body := bytes.NewBufferString(strings.Join(fields[1:], " "))
    resp, err := http.Post("http://localhost:"+strconv.Itoa(port)+"/servers", "text/plain", body)
    defer resp.Body.Close()
    if err != nil { panic(err) }
  }
}
