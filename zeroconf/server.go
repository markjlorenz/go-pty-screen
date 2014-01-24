package zeroconf

import (
  "os/exec"
  "strconv"
)

type Server struct {
  zeroconf *exec.Cmd
}

func StartAnnounce(port int) {
  // Start mDNS/zeroconf/Bonjour/Avahi, whatever it's called
  zeroconf := exec.Command("dns-sd", "-R", "GoPtyScreen", "_http._tcp", ".", strconv.Itoa(port))
  defer zeroconf.Process.Kill()

  zeroconf.Start()
  zeroconf.Wait()
}

