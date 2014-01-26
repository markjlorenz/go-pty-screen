package zeroconf

import (
  "github.com/dapplebeforedawn/go-dnssd"
)

func StartAnnounce(port int) {
  rc := make(chan *dnssd.RegisterReply)
  ctx, err := dnssd.ServiceRegister(
    0,
    0,
    "GoPtyScreen",
    "_goptyscreen._tcp.",
    "",
    "",
    2000,
    rc,
  )
  if err != nil { println(err); return }

  go dnssd.Process(ctx)
}

// package zeroconf
//
// import (
//   "os/exec"
//   "strconv"
// )
//
// type Server struct {
//   zeroconf *exec.Cmd
// }
//
// func StartAnnounce(port int) {
//   // Start mDNS/zeroconf/Bonjour/Avahi, whatever it's called
//   zeroconf := exec.Command("dns-sd", "-R", "GoPtyScreen", "_goptyscreen._tcp", ".", strconv.Itoa(port))
//   defer zeroconf.Process.Kill()
//
//   zeroconf.Start()
//   zeroconf.Wait()
// }
