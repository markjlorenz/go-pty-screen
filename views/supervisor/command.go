package supervisor_views

import (
  "code.google.com/p/goncurses"
)

type Command struct {
  *goncurses.Window
  prompt_color int16
  error_color int16
}

func NewCommand() (command *Command){
  command = new(Command)
  win_height, win_width := goncurses.StdScr().Maxyx()
  window, err           := goncurses.NewWindow(3, win_width, win_height-3, 0)
  if err != nil { panic(err) }

  command.init_colors()

  command.Window = &window
  command.clear()
  return
}

func (command *Command) init_colors() {
  command.prompt_color = 10
  err := goncurses.InitPair(command.prompt_color, goncurses.C_MAGENTA, goncurses.C_BLACK)
  if err != nil { panic(err) }

  command.error_color = 11
  err = goncurses.InitPair(command.error_color, goncurses.C_RED, goncurses.C_BLACK)
  if err != nil { panic(err) }
}

func (command *Command) Border() {
  lasty, lastx := command.Getyx()
  command.Box('|', '_')
  command.MovePrint(0, 2, "[ Enter a command: ]")
  command.Move(lasty, lastx)
}

func (command *Command) clear() {
  command.Clear()
  command.ColorOn(command.prompt_color)
  command.MovePrint(1, 1, "> ")
  command.ColorOff(command.prompt_color)
  command.Border()
  command.Refresh()
}

func (command *Command) GetInput() (input string){
  input_limit := 1024
  input, _ = command.GetString(input_limit)
  command.clear()
  command.Touch() // so we get the cursor back
  return
}

func (command *Command) FlashError(message string) (){
  command.clear()
  command.ColorOn(command.error_color)
  command.Print(message)
  command.ColorOn(command.error_color)
  command.Refresh()
  goncurses.NapMilliseconds(1000)
  command.clear()
}
