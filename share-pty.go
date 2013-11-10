package main

import (
  "fmt"
  "github.com/kr/pty"
  "code.google.com/p/go.crypto/ssh/terminal"
  "io"
  "os"
  "os/exec"
)

func main() {
  tty, _       := os.Open("/dev/tty")
  tty_fd       := int( tty.Fd() )
  tty_mode, _  := terminal.MakeRaw( tty_fd )

  exit := func () {
    fmt.Println("type 'reset' to restore your terminal")
    // why u no work?
    terminal.Restore( tty_fd, tty_mode )
  }

  c := exec.Command("vim")
  f, err := pty.Start(c)
  if err != nil { panic(err) }

  go io.Copy(f, os.Stdin)
  go io.Copy(os.Stdout, f)

  c.Wait()
  exit()
}

