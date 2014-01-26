package zeroconf

import (
  "github.com/dapplebeforedawn/go-dnssd"
)

type Client struct {
  Host string
  Port int
}

func NewClient() (*Client) {
  return &Client{}
}

func (c *Client) Dial() {
  println("Waiting for available server...")

  bc := make(chan *dnssd.BrowseReply)
  ctx, err := dnssd.Browse(dnssd.DNSServiceInterfaceIndexAny, "_goptyscreen._tcp", bc)
  if err != nil { println(err); return }

  go dnssd.Process(ctx)

  for {
    browseReply, ok := <-bc
    if !ok { println("No suitable server found."); break }

    rc := make(chan *dnssd.ResolveReply)
    rctx, err := dnssd.Resolve(
      dnssd.DNSServiceFlagsForceMulticast,
      browseReply.InterfaceIndex,
      browseReply.ServiceName,
      browseReply.RegType,
      browseReply.ReplyDomain,
      rc,
    )
    if err != nil { println(err); return }

    go dnssd.Process(rctx)
    resolveReply, _ := <-rc

    // qc := make(chan *dnssd.QueryRecordReply)
    // qctx, err := dnssd.QueryRecord(
    //   dnssd.DNSServiceFlagsForceMulticast,
    //   resolveReply.InterfaceIndex,
    //   resolveReply.FullName,
    //   dnssd.DNSServiceType_SRV,
    //   dnssd.DNSServiceClass_IN,
    //   qc,
    // )
    // if err != nil { println(err); return }

    // go dnssd.Process(qctx)
    // queryRecordReply, _ := <-qc

    // gc := make(chan *dnssd.GetAddrInfoReply)
    // gctx, err := dnssd.GetAddrInfo(
    //   dnssd.DNSServiceFlagsForceMulticast,
    //   0,
    //   dnssd.DNSServiceProtocol_IPv4,
    //   resolveReply.HostTarget,
    //   gc,
    // )
    // if err != nil { println(err); return }

    // go dnssd.Process(gctx)
    // getAddrInfoReply, _ := <-gc

    // c.Host = getAddrInfoReply.Ip
    c.Host = resolveReply.HostTarget
    c.Port = int(resolveReply.Port)
    break
  }

}
