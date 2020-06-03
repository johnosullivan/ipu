package cidr

import (
	"fmt"
  "log"
  "net"
  "time"
  "gopkg.in/gookit/color.v1"
)

type ParsedCIDR struct {
	FirstIP net.IP
	LastIP  net.IP
	IPNet   *net.IPNet
	IsIPv4  bool
	IsIPv6  bool
}

func ParseCIDR(s string) (*ParsedCIDR, error) {
	firstIP, ipNet, err := net.ParseCIDR(s)
	if err != nil {
		return nil, err
	}

	var v4, v6 bool

	if firstIP.To4() != nil {
		v4 = true
	} else {
		v6 = true
	}

	parsed := ParsedCIDR{
		FirstIP: firstIP,
		IPNet:   ipNet,
		IsIPv4:  v4,
		IsIPv6:  v6,
	}

	if parsed.IsIPv4 {
		parsed.LastIP = parsed.getLastIPv4()
	} else {
		parsed.LastIP = parsed.getLastIPv6()
	}

	return &parsed, nil
}

func inc(ip net.IP) {
	for j := len(ip)-1; j>=0; j-- {
		ip[j]++
		if ip[j] > 0 {
			break
		}
	}
}

func tcpScanner(ip string, ports []string, timeout int) map[string]bool {
    results := make(map[string]bool)
    for _, port := range ports {
        address := net.JoinHostPort(ip, port)
        conn, err := net.DialTimeout("tcp", address, time.Duration(timeout) * time.Second)
        if err != nil {
            results[port] = false
        } else {
            if conn != nil {
                results[port] = true
                _ = conn.Close()
            } else {
                results[port] = false
            }
        }
    }
    return results
}

func CIDRBlockDetails(cdir string, list bool, ports []string, timeout int) {
  p, err := ParseCIDR(cdir)
	if err != nil {
		log.Fatal(err)
	}

  color.Style{color.FgWhite, color.OpBold}.Println("CIDR Subnet Details")
  fmt.Println("  Subnet: ", cdir)
	fmt.Println("  First IP: ", p.FirstIP)
	fmt.Println("  Last IP: ", p.LastIP)
	if p.IsIPv4 {
		fmt.Println("  Total Host: ", p.HostCountIPv4())
	} else {
		fmt.Println("  Total Host: ", p.HostCountIPv6())
	}
  fmt.Println("")

  if list {
    if len(ports) != 0 {
      color.Style{color.FgWhite, color.OpBold}.Println("ALL Hosts w/")
      color.Style{color.FgGray, color.OpBold}.Println(timeout, " seconds timeout")
    } else {
      color.Style{color.FgWhite, color.OpBold}.Println("ALL Hosts")
    }
    ip, ipnet, err := net.ParseCIDR(cdir)
  	if err != nil {
  		log.Fatal(err)
  	}
  	for ip := ip.Mask(ipnet.Mask); ipnet.Contains(ip); inc(ip) {
      if len(ports) != 0 {
        results := tcpScanner(ip.String(), ports, timeout)
        fmt.Println(ip)
        fmt.Print("  -- ")
        for port, isOpened := range results {
          if isOpened {
            color.Green.Print(port)
          } else {
            color.Red.Print(port)
          }
          fmt.Print(" ")
        }
        fmt.Println("")
      } else {
        fmt.Println(ip)
      }
  	}
    fmt.Println("")
  }
}
