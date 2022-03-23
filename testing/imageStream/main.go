package main

import (
	"bytes"
	"fmt"
	"image/jpeg"
	"io/ioutil"
	"math/rand"
	"net"
	"net/http"
	"time"

	"github.com/pion/mdns"
	"golang.org/x/net/ipv4"
)

var (
	imageNames = []string{
		"blackRusty.jpg", "rustyCum.jpg", "rustySmiling.jpg",
	}
)

const (
	boundaryWord = "MJPEGBOUNDARY"
	headerf      = "\r\n" +
		"--" + boundaryWord + "\r\n" +
		"Content-Type: image/png\r\n" +
		"\r\n"
)

func streamImage(res http.ResponseWriter, req *http.Request) {
	res.Header().Add("Content-Type", "multipart/x-mixed-replace; boundary="+boundaryWord)
	b, _ := ioutil.ReadFile(imageNames[rand.Intn(2)])
	go func() {
		for {
			b, _ = ioutil.ReadFile(imageNames[rand.Intn(2)])
			time.Sleep(time.Second)
		}
	}()
	for {

		img, _ := jpeg.Decode(bytes.NewReader(b))
		if _, err := res.Write([]byte(headerf)); err != nil {
			return
		}

		if err := jpeg.Encode(res, img, &jpeg.Options{
			Quality: jpeg.DefaultQuality,
		}); err != nil {
			return
		}

	}
}

func main() {
	addr, err := net.ResolveUDPAddr("udp", mdns.DefaultAddress)
	if err != nil {
		panic(err)
	}
	l, err := net.ListenUDP("udp4", addr)
	if err != nil {
		panic(err)
	}
	server, err := mdns.Server(ipv4.NewPacketConn(l), &mdns.Config{
		LocalNames: []string{
			"camdeye1.local",
		},
	})
	if err != nil {
		panic(err)
	}
	defer server.Close()
	http.HandleFunc("/", streamImage)
	fmt.Println(http.ListenAndServe(":80", nil))

}
