package client_views

import (
  "dapplebeforedawn/share-pty/servers"
  "code.google.com/p/goncurses"
  "strings"
  "strconv"
  "fmt"
)

type List struct {
  *goncurses.Window
  header_color int16
}

func NewList() (list *List){
  list = new(List)
  win_height, win_width := goncurses.StdScr().Maxyx()
  window, err           := goncurses.NewWindow(win_height, win_width, 0, 0)
  if err != nil { panic(err) }

  list.header_color = 20
  err = goncurses.InitPair(list.header_color, goncurses.C_MAGENTA, goncurses.C_BLACK)
  if err != nil { panic(err) }

  list.Window = &window
  list.draw_initial()
  return
}

func (list *List) draw_initial() {
  list.Move(1, 2)
  list.Border()
}

func (list *List) AddItem(item pty_servers.PtyShare) (){
  lasty, _    := list.Getyx()
  list.MovePrintln(lasty, 2, list.build_row(item.Alias, item.Command))
  list.Border()
}

func (list *List) build_row(alias, command string) (string){
  _, row_length := list.Maxyx()
  field_count   := 3 // 5 fields in PtyShare
  segment_size  := (row_length / field_count) - 1
  format_string := strings.Repeat("%-"+strconv.Itoa(segment_size)+"s", field_count)

  return fmt.Sprintf(format_string, alias, command, "")
}

func (list *List) Border() {
  lasty, lastx := list.Getyx()
  list.Box('|', '_')
  list.MovePrint(0, 2, "[ Available PTYs ]")
  list.Move(lasty, lastx)
}
