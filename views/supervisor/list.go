package supervisor_views

import (
  "code.google.com/p/goncurses"
  "dapplebeforedawn/share-pty/servers"
  "strings"
  "strconv"
  "fmt"
)

type List struct {
  *goncurses.Window
}

func NewList() (list *List){
  list = new(List)
  win_height, win_width := goncurses.StdScr().Maxyx()
  window, err := goncurses.NewWindow(win_height-3, win_width, 0, 0)
  if err != nil { panic(err) }

  list.Window = &window
  list.MovePrintln(1, 2, list.build_row("ALIAS", "COMMAND", "KEY_PORT", "SCREEN_PORT"))
  list.MovePrintln(2, 2, "No servers running.  Type: `new <alias> <command> <rows> <cols>` into the command window to start one.")
  list.Move(2, 2)
  list.Border()
  return
}

func (list *List) AddItem(item pty_servers.PtyShare) (){
  lasty, _    := list.Getyx()

  key_port    := strconv.Itoa(item.KeyServer.Port)
  screen_port := strconv.Itoa(item.ScreenServer.Port)
  list.MovePrintln(lasty, 2, list.build_row(item.Alias, item.Command, key_port, screen_port))

  list.Border()
}

func (list *List) build_row(alias, command, key_port, screen_port string) (string){
  _, row_length := list.Maxyx()
  field_count   := 5 // 5 fields in PtyShare
  segment_size  := row_length / field_count
  format_string := strings.Repeat("%-"+strconv.Itoa(segment_size)+"s", field_count)

  return fmt.Sprintf(format_string, alias, command, key_port, screen_port, "")
}

func (list *List) Border() {
  lasty, lastx := list.Getyx()
  list.Box('|', '_')
  list.MovePrint(0, 2, "[ Available PTYs ]")
  list.Move(lasty, lastx)
}
