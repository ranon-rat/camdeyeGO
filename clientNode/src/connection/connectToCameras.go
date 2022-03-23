package connection

import (
	"bufio"
	"encoding/base64"
	"fmt"
	"image/jpeg"
	"io/ioutil"
	"log"
	"mime"
	"mime/multipart"
	"net/http"
	"sync"
	"time"

	"github.com/ranon-rat/camdeyeGO/clientNode/src/core"
)

func ConnectToCameras() {
	var wg sync.WaitGroup
	domains := []string{"camdeye1.local"} //database.GetCamerasDomain()

	for i, domain := range domains {

		cameras = append(cameras, core.Camera{
			Domain: domain,
			Up:     true,
		})
		wg.Add(1)

		go func(domain string, pos int) {
			defer wg.Done()
			defer func() {
				cameras[pos].LastTimeUp = time.Now().Unix()
				log.Println("camera isnt active or it doesnt exist", cameras[pos].LastTimeUp)
				cameras[pos].Up = false
			}()
			response, err := http.Get("http://" + domain)
			if err != nil {
				log.Println("unexpected error happened", err)
				return

			}

			_, params, err := mime.ParseMediaType(response.Header.Get("Content-Type"))
			if err != nil {
				log.Println("invalid body format", err)
				return
			}
			boundary, exist := params["boundary"]

			if !exist {
				log.Println("boundary should exists")
				return

			}
			buf := bufio.NewReader(response.Body)
			before := time.Now().Unix()
			for {

				part, err := multipart.NewReader(buf, boundary).NextPart()
				if err != nil {
					log.Println("unexpected error happened", err)
					break
				}
				// this is in seconds so i save this every ten seconds
				if time.Now().Unix()-before >= 10 {

					img, err := jpeg.Decode(part)
					if err != nil {

						break
					}
					f, err := ioutil.TempFile("./images", "test*.jpg")
					if err != nil {
						log.Println("unexpected error happened", err)
						break
					}
					jpeg.Encode(f, img, defaultJPEGConfig)
					before = time.Now().Unix()
					continue
				}
				partBytes, _ := ioutil.ReadAll(part)
				cameras[pos].Image = "data:image/jpeg;base64," + base64.RawStdEncoding.EncodeToString(partBytes)
				fmt.Println(cameras[pos].Up, cameras[pos].Domain)
			}

		}(domain, i)

	}
	wg.Wait()
}
