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
  header_color     int16
  error_color      int16
  selection_color  int16
  current_item     int
  items            []pty_servers.PtyShare
}

func NewList() (list *List){
  list = new(List)
  win_height, win_width := goncurses.StdScr().Maxyx()
  window, err           := goncurses.NewWindow(win_height, win_width, 0, 0)
  if err != nil { panic(err) }

  list.Window = &window
  list.init_colors()
  list.draw_initial()

  go list.SelectRow()
  return
}

func (list *List) init_colors() {
  list.header_color = 20
  err := goncurses.InitPair(list.header_color, goncurses.C_MAGENTA, goncurses.C_BLACK)
  if err != nil { panic(err) }

  list.error_color = 21
  err = goncurses.InitPair(list.error_color, goncurses.C_RED, goncurses.C_BLACK)
  if err != nil { panic(err) }

  list.selection_color = 22
  err = goncurses.InitPair(list.selection_color, goncurses.C_BLACK, goncurses.C_WHITE)
  if err != nil { panic(err) }
}

func (list *List) draw_initial() {
  list.Move(1, 2)
  list.ColorOn(list.header_color)
  list.Println(list.build_row("ALIAS", "COMMAND"))
  list.ColorOff(list.header_color)
  list.Border()
}

func (list *List) refresh(){
  list.Clear()
  list.draw_initial()
  for _, item := range list.items {
    list.print_row(item)
  }
  list.Refresh()
}

func (list *List) AddItem(item pty_servers.PtyShare) (){
  list.items = append(list.items, item)
  list.refresh()
}

func (list *List) print_row(item pty_servers.PtyShare) (){
  lasty, _    := list.Getyx()
  if (item == list.items[list.current_item]){
    list.ColorOn(list.selection_color)
    list.MovePrintln(lasty, 2, list.build_row(item.Alias, item.Command))
    list.ColorOff(list.selection_color)
  } else {
    list.MovePrintln(lasty, 2, list.build_row(item.Alias, item.Command))
  }
  list.Border()
}

func (list *List) build_row(alias, command string) (string){
  _, row_length := list.Maxyx()
  field_count   := 3
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

func (list *List) SelectRow() (pty_servers.PtyShare){
  for {
    char := list.GetChar()
    switch char {
    case 'k':
      list.selection_up()
    case 'j':
      list.selection_down()
    case 10:
      return list.items[list.current_item]
    default:
      list.FlashError("Such fail. WoW.  Try 'j', 'k', or <enter>.")
    }
  }
}

func (list *List) selection_up() {
  if (list.current_item <= 0) {
    list.current_item = 0
  } else { list.current_item -= 1 }
  list.refresh()
}

func (list *List) selection_down(){
  if (list.current_item >= len(list.items) - 1) {
    list.current_item = len(list.items) - 1
  } else { list.current_item += 1 }
  list.refresh()
}

func (list *List) FlashError(message string) () {
  bottom, _ := list.Maxyx()
  list.refresh()
  list.ColorOn(list.error_color)
  list.MovePrint(bottom-2, 2, message)
  list.ColorOff(list.error_color)
  list.Refresh()
  goncurses.NapMilliseconds(1000)
  list.refresh()
}
