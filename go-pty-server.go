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

  create_feed := make(chan pty_servers.PtyShare)
  delete_feed := make(chan string)
  view := pty_views.NewSupervisor()
  go view.CreateFeed(create_feed)
  go view.DeleteFeed(delete_feed)
  go view.WatchCommands(opts.Port)
  view.Refresh()

  supervisor := pty_servers.NewSupervisor(create_feed, delete_feed)
  supervisor.Listen(opts.Port)
}
