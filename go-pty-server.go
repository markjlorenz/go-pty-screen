package main

import (
  "dapplebeforedawn/share-pty/servers"
  "dapplebeforedawn/share-pty/options"
  "dapplebeforedawn/share-pty/views"
  "os"
)

func main() {
  opts := options.Server{}
  opts.Parse()

  create_feed := make(chan pty_servers.PtyShare)
  delete_feed := make(chan string)
  ready_feed  := make(chan int)
  view := pty_views.NewSupervisor()
  go view.CreateFeed(create_feed)
  go view.DeleteFeed(delete_feed)
  go view.WatchCommands(opts.Port)
  view.Refresh()

  rc_filename := os.Getenv("HOME")+"/.go-pty-rc"
  rc_loader   := pty_servers.NewRCLoader(ready_feed, rc_filename)
  go rc_loader.OnReady()

  supervisor := pty_servers.NewSupervisor(create_feed, delete_feed)
  supervisor.Listen(opts.Port, ready_feed)
}
