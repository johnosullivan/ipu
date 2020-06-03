package main

import (
	"fmt"
  "os"
  "flag"
  "net"
  "strings"
  "gopkg.in/gookit/color.v1"
)

var (
  version = "0.1.0"
)

func main() {
  helpPtr := flag.Bool("h", false, "")
  versionPtr := flag.Bool("v", false, "Get the current version on the ipu cli.")
  listIPPtr := flag.Bool("l", false, "List all possible IP adddresses within a given a CIDR block.")
  cidrBlockPtr := flag.String("b", "", "CIDR Block (IPv4), for example: 192.0.0.0/8.")
  existIPPtr := flag.String("ip", "", "IPv4 Addresses, for example: 192.0.0.0,192.0.0.1,192.0.0.3.")

  flag.Parse()

  if *versionPtr {
    color.Style{color.FgWhite, color.OpBold}.Println("ipu/", version)
    os.Exit(0)
  }

  if *helpPtr {
    fmt.Println("")
    os.Exit(0)
  }

  required := []string{"b"}

  seen := make(map[string]bool)
  flag.Visit(func(f *flag.Flag) { seen[f.Name] = true })
  for _, req := range required {
      if !seen[req] {
          fmt.Fprintf(os.Stderr, "missing required -%s arguments/flags\n", req)
          os.Exit(2)
      }
  }

  if *existIPPtr != "" {
    CIDRBlockDetails(*cidrBlockPtr, *listIPPtr);
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
    CIDRBlockDetails(*cidrBlockPtr, *listIPPtr);
  }
}
