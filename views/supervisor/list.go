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
  items map[string] *pty_servers.PtyShare
  header_color int16
}

func NewList() (list *List){
  list = new(List)
  win_height, win_width := goncurses.StdScr().Maxyx()
  window, err           := goncurses.NewWindow(win_height-3, win_width, 0, 0)
  if err != nil { panic(err) }


  list.Window = &window
  list.items  = make(map[string] *pty_servers.PtyShare)
  list.init_colors()
  list.draw_initial()
  return
}

func (list *List) init_colors() {
  list.header_color = 20
  err := goncurses.InitPair(list.header_color, goncurses.C_MAGENTA, goncurses.C_BLACK)
  if err != nil { panic(err) }
}

func (list *List) draw_initial(){
  list.draw_list()
  list.MovePrintln(2, 2, "No servers running.  Type: `new <alias> <command> <rows> <cols>` into the command window to start one.")
  list.Border()
}

func (list *List) AddItem(item pty_servers.PtyShare) (){
  list.items[item.Alias] = &item
  list.draw_list()
  list.Refresh()
}

func (list *List) RemoveItem(alias string) (){
  delete(list.items, alias)
  list.draw_list()
  list.Refresh()
}

func (list *List) draw_list(){
  list.Clear()

  list.ColorOn(list.header_color)
  list.MovePrintln(1, 2, list.build_row("ALIAS", "COMMAND", "KEY_PORT", "SCREEN_PORT"))
  list.ColorOff(list.header_color)
  list.Move(2, 2)
  for _, item := range list.items {
    cury, _     := list.Getyx()
    key_port    := strconv.Itoa(item.KeyServer.Port)
    screen_port := strconv.Itoa(item.ScreenServer.Port)
    list.MovePrintln(cury, 2, list.build_row(item.Alias, item.Command, key_port, screen_port))
  }
  list.Border()
}

func (list *List) build_row(alias, command, key_port, screen_port string) (string){
  _, row_length := list.Maxyx()
  field_count   := 5 // 5 fields in PtyShare
  segment_size  := (row_length / field_count) - 1
  format_string := strings.Repeat("%-"+strconv.Itoa(segment_size)+"s", field_count)

  return fmt.Sprintf(format_string, alias, command, key_port, screen_port, "")
}

func (list *List) Border() {
  lasty, lastx := list.Getyx()
  list.Box('|', '-')
  list.MovePrint(0, 2, "[ Available PTYs ]")
  list.Move(lasty, lastx)
}
