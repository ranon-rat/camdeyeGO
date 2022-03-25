package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/ranon-rat/camdeyeGO/api/src/core"
)

// i will change it to fiber later but for now its not a big dial
func AddIceOffer(w http.ResponseWriter, req *http.Request) {
	var iceOffer core.CameraIceOffer
	iceOffer.IP = req.RemoteAddr
	err := json.NewDecoder(req.Body).Decode(&iceOffer)
	if err != nil {
		http.Error(w, "http bad request ,the body request is wrong", http.StatusBadRequest)
		return
	}
	if iceOffer.Token == "" || iceOffer.IceOffer == "" {
		http.Error(w, "http bad request ,the body request is wrong", http.StatusBadRequest)
		return
	}
	/*
		idInt,_:=database.VerifyIfExists(iceOffer.Token)
		id:=strconv.Itoa(idInt)
	*/
	id := "1"
	ip := req.RemoteAddr

	if _, ok := camerasIPs[ip]; ok { // if i get the ip

		if _, ok := camerasNodesIceOffer[id]; !ok {
			// check if you are in the camera nodes ice offer
			//if you arent , you will receive an error
			http.Error(w, "ip is already here", http.StatusBadRequest)
		}
	}

}
