package main

import (
	"fmt"
  "log"
  "net"
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

func CIDRBlockDetails(cdir string, list bool) {
  p, err := ParseCIDR(cdir)
	if err != nil {
		log.Fatal(err)
	}
  color.Style{color.FgWhite, color.OpBold}.Println("CIDR Block Details")
  fmt.Println("  Subnet: ", cdir)
	fmt.Println("  First IP: ", p.FirstIP)
	fmt.Println("  Last IP: ", p.LastIP)
	if p.IsIPv4 {
		fmt.Println("  Total Host: ", p.HostCountIPv4())
	} else {
		fmt.Println("  Total Host: ", p.HostCountIPv6())
	}

  if list {
    color.Style{color.FgWhite, color.OpBold}.Println("ALL Hosts")
    ip, ipnet, err := net.ParseCIDR(cdir)
  	if err != nil {
  		log.Fatal(err)
  	}
  	for ip := ip.Mask(ipnet.Mask); ipnet.Contains(ip); inc(ip) {
  		fmt.Println(ip)
  	}
  }
}
