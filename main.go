package main

import (
	"fmt"
  "os"
  "flag"
  "net"
  "strings"
  "gopkg.in/gookit/color.v1"
  "github.com/johnosullivan/ipu/cidr"
)

var (
  version = "0.3.0"
)

func main() {
  helpPtr := flag.Bool("h", false, "")
  versionPtr := flag.Bool("v", false, "Get the current version of the ipu cli.")
  listIPPtr := flag.Bool("l", false, "List all possible IP adddresses within a given a CIDR block.")
  cidrBlockPtr := flag.String("sn", "", "CIDR subnet (IPv4/IPv6), for example: 192.0.0.0/8 or 2002::1234:abcd:ffff:c0a8:101/122.")
  existIPPtr := flag.String("ip", "", "IPv4 Addresses, for example: 192.0.0.0 or 2002::1234:abcd:ffff:c0a8:100.")
  portsPtr := flag.String("p", "", "")
  pingTimeoutPtr := flag.Int("pto", 3, "Ping port timeout window in seconds. ")
  flag.Parse()

  if *versionPtr {
    color.Style{color.FgWhite, color.OpBold}.Println("ipu/", version)
    os.Exit(0)
  }

  if *helpPtr {
    fmt.Println("")
    os.Exit(0)
  }

  required := []string{"sn"}

  seen := make(map[string]bool)
  flag.Visit(func(f *flag.Flag) { seen[f.Name] = true })
  for _, req := range required {
      if !seen[req] {
          fmt.Fprintf(os.Stderr, "Missing required -%s arguments/flags\n", req)
          os.Exit(2)
      }
  }

  ports := []string{}

  if *portsPtr != "" {
    ports = strings.Split(*portsPtr, ",")
  }

  if *existIPPtr != "" {
    cidr.CIDRBlockDetails(*cidrBlockPtr, *listIPPtr, ports, *pingTimeoutPtr);
    color.Style{color.FgWhite, color.OpBold}.Println("Range Results")
    clientips := strings.Split(*existIPPtr, ",")
    _, subnet, _ := net.ParseCIDR(*cidrBlockPtr)
    for _, clientip := range clientips {
        ip := net.ParseIP(clientip)
        if subnet.Contains(ip) {
            color.Green.Println(clientip, " in subnet ", subnet)
        } else {
            color.Red.Println(clientip, " not in subnet ", subnet)
        }
    }
  } else {
    cidr.CIDRBlockDetails(*cidrBlockPtr, *listIPPtr, ports, *pingTimeoutPtr);
  }
}
