package main

import "net/http"

func main() {

	//here first argument i.e. put path that handler register on
	http.HandleFunc("/hiworld", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("hello http server working"))
	})

	http.ListenAndServe(":8080", nil) //initilize http server -step1
	//ListenAndserve launches default http server on our desired port
}
