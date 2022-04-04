package main

import (
	"net/http"
)

func main() {

	http.HandleFunc("/", helloHandler)
	http.ListenAndServe(":3000", nil)
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("duckhue01"))
}

func listHandler(w http.ResponseWriter, r *http.Request) {

}
