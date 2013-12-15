package pty_views

import (
  "code.google.com/p/goncurses"
  "dapplebeforedawn/share-pty/views/supervisor"
  "time"
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

  supervisor.list_window    = supervisor_views.NewList()
  supervisor.command_window = supervisor_views.NewCommand()

  supervisor.list_window.Refresh()
  supervisor.command_window.Refresh()

  return
}

