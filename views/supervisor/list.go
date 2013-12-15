package supervisor_views

import (
  "code.google.com/p/goncurses"
)

type List struct {
  *goncurses.Window
}

func NewList() (list *List){
  list = new(List)
  win_height, win_width := goncurses.StdScr().Maxyx()
  window, err := goncurses.NewWindow(win_height-5, win_width, 0, 0)
  if err != nil { panic(err) }

  list.Window = &window
  list.Box('|', '_')
  return
}

