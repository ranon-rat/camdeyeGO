package router

import (
	"net/http"
	"os"
)

func setup() error {

	port, err := os.LookupEnv("PORT")
	if !err {
		port = "8080"
	}
	return http.ListenAndServe(":"+port, nil)

}
