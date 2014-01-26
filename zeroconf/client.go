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

    c.Host = resolveReply.HostTarget
    c.Port = int(resolveReply.Port)
    break
  }
}
