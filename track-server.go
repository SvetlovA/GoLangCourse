// track-server
package main

import (
	"encoding/base64"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		decodeAddress := r.URL.Path[1:]
		encodeAddress, err := base64.StdEncoding.DecodeString(decodeAddress)
		if err != nil {
			log.Fatal(err)
		}

		http.Redirect(w, r, "http://"+string(encodeAddress), http.StatusSeeOther)
	})

	log.Fatal(http.ListenAndServe(":8085", nil))
}
