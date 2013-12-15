package main

import (
  "dapplebeforedawn/share-pty/servers"
  "dapplebeforedawn/share-pty/options"
  "dapplebeforedawn/share-pty/views"

  "fmt"
)

func main() {
  fmt.Println("Started")
  opts := options.Server{}
  opts.Parse()

  view_feed := make(chan pty_servers.PtyShare)
  view := pty_views.NewSupervisor()
  go view.WatchFeed(view_feed)
  go view.WatchCommands(opts.Port)
  view.Refresh()

  supervisor := pty_servers.NewSupervisor(view_feed)
  supervisor.Listen(opts.Port)
}
