package main

import (
	"fmt"
  "os"
  "flag"
  "net"
  "strings"
)

func main() {
  helpPtr := flag.Bool("h", false, "")
  cidrBlockPtr := flag.String("b", "", "CIDR Block (IPv4), for example: 192.0.0.0/8")
  existIPPtr := flag.String("ip", "", "IPv4 Addresses, for example: 192.0.0.0,192.0.0.1,192.0.0.3")

  flag.Parse()

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
    fmt.Println(*cidrBlockPtr)
    fmt.Println(*existIPPtr)
  } else {
    fmt.Println(*cidrBlockPtr)
  }

}
