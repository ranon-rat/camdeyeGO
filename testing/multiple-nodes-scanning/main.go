package main

import (
	"crypto/rand"
	"fmt"
	"math/big"
	"net"
	"net/http"
	"strconv"
	"sync"

	"github.com/pion/mdns"
	"golang.org/x/net/ipv4"
)

var names = []string{}

func CreateMultipleServersIDK() {
	m := http.NewServeMux()
	bigInt, _ := rand.Int(rand.Reader, big.NewInt(999))

	m.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "hello you are in "+r.Host+" in the port "+bigInt.String())
	})

	http.ListenAndServe(":8"+bigInt.String(), m)

}
func main() {
	var wg sync.WaitGroup
	for i := 1; i <= 20; i++ {
		name := "camdeye" + strconv.Itoa(i) + ".local"
		names = append(names, name)
		wg.Add(1)
		go func() {
			defer wg.Done()
			CreateMultipleServersIDK()
		}()

	}
	fmt.Println(names)
	addr, err := net.ResolveUDPAddr("udp", mdns.DefaultAddress)
	if err != nil {
		panic(err)
	}
	l, err := net.ListenUDP("udp4", addr)
	if err != nil {
		panic(err)
	}
	server, err := mdns.Server(ipv4.NewPacketConn(l), &mdns.Config{
		LocalNames: names,
	})
	if err != nil {
		panic(err)
	}
	defer server.Close()
	wg.Wait()
}
