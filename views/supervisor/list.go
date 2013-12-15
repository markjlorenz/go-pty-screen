package supervisor_views

import (
  "code.google.com/p/goncurses"
  "dapplebeforedawn/share-pty/servers"
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
  list.Move(1, 1)
  list.Border()
  return
}

func (list *List) AddItem(item pty_servers.PtyShare) (){
  lasty, _ := list.Getyx()
  list.MovePrintln(lasty, 2, item.Alias)
  list.Border()
}

func (list *List) Border() {
  lasty, lastx := list.Getyx()
  list.Box('|', '_')
  list.MovePrint(0, 2, "[ Available PTYs ]")
  list.Move(lasty, lastx)
}
