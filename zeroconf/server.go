package zeroconf

import (
  "github.com/dapplebeforedawn/go-dnssd"
)

type Server struct {
  ServiceType string
  TxtRecords  map[string]string
}

func NewServer(serviceType string) *Server {
  return &Server{
    ServiceType: serviceType,
  }
}

//"_goptyscreen._tcp."
func StartAnnounce(port int) (*dnssd.Context, error){
  rc := make(chan *dnssd.RegisterReply)
  ctx, err := dnssd.ServiceRegister(
    dnssd.DNSServiceFlagsSuppressUnusable,
    0,
    "GoPtyScreen",
    s.ServiceType,
    "",
    "",
    (uint16)(port),
    c.TxtRecords,
    rc,
  )

  if err != nil { panic(err); }

  go dnssd.Process(ctx)

  _, _ = <-rc // wait for the register reply

  return ctx, nil
}
