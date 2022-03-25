package router

import (
	"encoding/json"
	"net/http"
	"os"

	"github.com/ranon-rat/camdeyeGO/api/src/controllers"
)

func setup() error {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		json.NewEncoder(w).Encode(map[string]string{
			"camera-log-in": "http://" + r.Host + "/camera-log-in",
			"add-ice-offer": "http://" + r.Host + "/add-ice-offer",
			"get-ice-offer": "http://" + r.Host + "/get-ice-offer?id=<put the camera id>",
		})
	})
	http.HandleFunc("/add-ice-offer", controllers.AddIceOffer)
	http.HandleFunc("/get-ice-offer", controllers.GetIceOffer)
	port, err := os.LookupEnv("PORT")
	if !err {
		port = "8080"
	}
	go controllers.PingCameras()
	return http.ListenAndServe(":"+port, nil)

}
