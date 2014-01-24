package zeroconf

import (
  "os/exec"
  "regexp"
  "strings"
  "bufio"
)

type Client struct {
  zeroconf *exec.Cmd
}

func NewClient() (*Client) {
  //dns-sd -L MyTest _http._tcp. local.
  zeroconf := exec.Command("dns-sd", "-L", "GoPtyScreen", "_http._tcp.", "local.")
  return &Client{zeroconf: zeroconf}
}

func (c *Client) Dial() (host string) {
  zeroconf_stdout, _ := c.zeroconf.StdoutPipe()
  zeroconf_buff      := bufio.NewReader(zeroconf_stdout)
  c.zeroconf.Start()
  defer c.zeroconf.Process.Kill()

  println("Waiting for available server...")
  for {
    // read from zeroconf until we get a line matching: GoPtyScreen._http._tcp.local.
    line, zc_err := zeroconf_buff.ReadString('\n')
    if (zc_err != nil){ break; }
    match, _ := regexp.MatchString("GoPtyScreen._http._tcp.local.", line)
    if (match) {
      //GoPtyScreen._http._tcp.local. can be reached at BabyGoat.local.:3000 (interface 5)
      fields := strings.Fields(line)
      host    = strings.Split(fields[6], ":")[0]
      break
    }
  }
  return
}
