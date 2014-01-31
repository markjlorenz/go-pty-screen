package zeroconf

import (
  "github.com/dapplebeforedawn/go-dnssd"
)

func StartAnnounce(port int) (*dnssd.Context, error){
  rc := make(chan *dnssd.RegisterReply)
  ctx, err := dnssd.ServiceRegister(
    dnssd.DNSServiceFlagsSuppressUnusable,
    0,
    "GoPtyScreen",
    "_goptyscreen._tcp.",
    "",
    "",
    (uint16)(port),
    nil,
    rc,
  )

  if err != nil { panic(err); }

  go dnssd.Process(ctx)

  _, _ = <-rc // wait for the register reply

  return ctx, nil
}
