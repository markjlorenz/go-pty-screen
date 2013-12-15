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
  window, err := goncurses.NewWindow(3, win_width, win_height-3, 0)
  if err != nil { panic(err) }

  command.Window = &window
  command.Move(1, 1)
  command.Border()
  return
}

func (command *Command) Border() {
  lasty, lastx := command.Getyx()
  command.Box('|', '_')
  command.MovePrint(0, 2, "[ Enter a command: ]")
  command.Move(lasty, lastx)
}
