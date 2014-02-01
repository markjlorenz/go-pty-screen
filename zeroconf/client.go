package zeroconf

import (
  "github.com/dapplebeforedawn/go-dnssd"
)

type Client struct {
  Host        string
  Port        int
  ServiceType string
}

func NewClient(serviceType string) (*Client) {
  return &Client{
    ServiceType: serviceType,
  }
}

func (c *Client) Dial() {
  println("Waiting for available server...")
  alwaysMatch := func(txtRecords map[string]string) bool { return true }
  c.DialWhenMatch(alwaysMatch)
}

type Matcher func(txtRecords map[string]string) bool

func (c *Client) DialWhenMatch (matcher Matcher) {
  bc := make(chan *dnssd.BrowseReply)
  ctx, err := dnssd.Browse(dnssd.DNSServiceInterfaceIndexAny, c.ServiceType, bc)
  if err != nil { println(err); return }

  defer ctx.Release()
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

    defer rctx.Release()
    go dnssd.Process(rctx)
    resolveReply, _ := <-rc

    if matcher(resolveReply.TxtRecordMap) {
      c.Host = resolveReply.HostTarget
      c.Port = int(resolveReply.Port)
      break
    }
  }
}
