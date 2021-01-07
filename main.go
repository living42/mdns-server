package main

import (
	"context"
	"flag"
	"net"
	"os"

	"github.com/hashicorp/mdns"
)

func main() {
	service := flag.String("service", "", "service")
	domain := flag.String("domain", "", "domain")
	port := flag.Int("port", 8000, "port")
	ip := flag.String("ip", "", "ip")

	flag.Parse()

	var ips []net.IP

	if *ip != "" {
		ips = append(ips, net.ParseIP(*ip))
	}

	host, _ := os.Hostname()
	info := []string{*service}
	svc, err := mdns.NewMDNSService(host, *service, *domain, "", *port, ips, info)
	if err != nil {
		panic(err)
	}

	// Create the mDNS server, defer shutdown
	server, err := mdns.NewServer(&mdns.Config{Zone: svc})
	if err != nil {
		panic(err)
	}
	<-context.Background().Done()
	defer server.Shutdown()
}
