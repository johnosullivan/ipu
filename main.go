package main

import (
	"fmt"
  "os"
  "flag"
  "strconv"
  "strings"
  "gopkg.in/gookit/color.v1"
  "github.com/johnosullivan/ipu/cidr"
)

var (
  version = "0.4.0"
)

func main() {
  helpPtr := flag.Bool("h", false, "more info: https://github.com/johnosullivan/ipu")
  versionPtr := flag.Bool("v", false, "current version (v" + version + ")")
  listIPPtr := flag.Bool("l", false, "list all possible IP adddresses within a given a CIDR block.")
  cidrBlockPtr := flag.String("sn", "", "CIDR subnet (IPv4/IPv6), for example: 192.0.0.0/8 or 2002::1234:abcd:ffff:c0a8:101/122.")
  existIPPtr := flag.String("ip", "", "IPv4 Addresses, for example: 192.0.0.0 or 2002::1234:abcd:ffff:c0a8:100.")
  portsPtr := flag.String("p", "", "ports which each host should try to ping, for example: 80,22,5432 or 75-80")
  pingTimeoutPtr := flag.Int("pto", 3, "ping port timeout window in seconds. ")
  flag.Parse()

  if *versionPtr {
    color.Style{color.FgWhite, color.OpBold}.Println("ipu/", version)
    os.Exit(0)
  }

  if *helpPtr {
    flag.PrintDefaults()
    os.Exit(0)
  }

  required := []string{"sn"}

  seen := make(map[string]bool)
  flag.Visit(func(f *flag.Flag) { seen[f.Name] = true })
  for _, req := range required {
      if !seen[req] {
          fmt.Fprintf(os.Stderr, "Missing required -%s arguments/flags\n\n", req)
          flag.PrintDefaults()
          os.Exit(2)
      }
  }

  ports := []string{}

  if strings.Contains(*portsPtr, "-") {
    portRange := strings.Split(*portsPtr, "-")
    start, _ := strconv.Atoi(portRange[0])
    end, _ := strconv.Atoi(portRange[1])

    for i := start; i <= end; i++ {
      ports = append(ports,strconv.FormatInt(int64(i), 10))
    }
  } else {
    if *portsPtr != "" {
      ports = strings.Split(*portsPtr, ",")
    }
  }

  cidr.CIDRBlockDetails(*cidrBlockPtr, *listIPPtr, ports, *pingTimeoutPtr)

  if *existIPPtr != "" {
    cidr.InCIDRBlock(*existIPPtr, *cidrBlockPtr)
  }
}
