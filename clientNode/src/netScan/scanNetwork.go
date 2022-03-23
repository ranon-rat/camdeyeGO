package netScan

import (
	"fmt"
	"net"
	"strconv"
	"sync"

	"github.com/ranon-rat/camdeyeGO/clientNode/src/database"
)

// i know that im using mdns with this name
// camDeye and with a number ant the end so i just need to make something like this
// in case that you want to use the urls or something like that you can use this
// but by default its going to store all the urls in the database
func getCamerasURL(limit int) (urls []string) {
	// this is going to mantain the goroutines executing until they  finish
	var wg sync.WaitGroup

	for i := 1; i <= limit; i++ {
		wg.Add(1)
		// remember that i will be using mdns for  knowing the directions of the cameras
		url := "camdeye" + strconv.Itoa(i) + ".local"

		go func(url string) {
			defer wg.Done()
			// im just using this for know if the device exists
			// maybe someone is faking it so , maybe i need to do something uhhh
			if _, err := net.LookupHost(url); err != nil {
				fmt.Println(err)
				return
			}
			database.AddURL(url)
			urls = append(urls, url)

		}(url)
	}
	wg.Wait()

	return urls
}
