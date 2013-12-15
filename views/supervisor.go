package pty_views

import (
  "code.google.com/p/goncurses"
  "dapplebeforedawn/share-pty/servers"
  "dapplebeforedawn/share-pty/views/supervisor"
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

  goncurses.Cursor(0) // no cursor please

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

func (supervisor *Supervisor) WatchFeed(feed chan pty_servers.PtyShare) {
  for update := range feed {
    supervisor.list_window.AddItem(update)
    supervisor.list_window.Refresh()
  }
}
