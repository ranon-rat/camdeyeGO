package controllers

import "net/http"

func GetIceOffer(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	if id == "" {
		http.Error(w, "http bad request ,the body request is wrong", http.StatusBadRequest)
		return
	}
	if _, ok := camerasNodesIceOffer[id]; !ok {
		http.Error(w, "http bad request ,the body request is wrong", http.StatusBadRequest)
		return
	}
	w.Write([]byte(camerasNodesIceOffer[id].IceOffer))

}
