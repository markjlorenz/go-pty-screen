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
    port,
    rc,
  )
  if err != nil { println(err); return }

  go dnssd.Process(ctx)
}
