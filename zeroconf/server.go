package zeroconf

import (
  "github.com/dapplebeforedawn/go-dnssd"
)

func StartAnnounce(port int) {
  rc := make(chan *dnssd.RegisterReply)
  _, err := dnssd.ServiceRegister(
    dnssd.DNSServiceFlagsSuppressUnusable,
    0,
    "GoPtyScreen",
    "_goptyscreen._tcp.",
    "",
    "",
    (uint16)(port),
    "markstxtrecord=1",
    rc,
  )
  if err != nil { panic(err); return }
}
