package supervisor_views

import (
  "code.google.com/p/goncurses"
)

type Command struct {
  *goncurses.Window
}

func NewCommand() (command *Command){
  command = new(Command)
  win_height, win_width := goncurses.StdScr().Maxyx()
  window, err := goncurses.NewWindow(5, win_width, win_height-5, 0)
  if err != nil { panic(err) }

  command.Window = &window
  command.Box('|', '_')
  return
}

