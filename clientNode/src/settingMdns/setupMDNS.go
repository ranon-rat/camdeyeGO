package settingMDNS

import (
	"net"

	"github.com/pion/mdns"
	"golang.org/x/net/ipv4"
)

func Setup() *mdns.Conn {

	addr, err := net.ResolveUDPAddr("udp", mdns.DefaultAddress)
	if err != nil {
		panic(err)
	}
	l, err := net.ListenUDP("udp4", addr)
	if err != nil {
		panic(err)
	}
	server, err := mdns.Server(ipv4.NewPacketConn(l), &mdns.Config{
		LocalNames: []string{"camdeye-admin.local"},
	})
	if err != nil {
		panic(err)
	}
	return server

}
