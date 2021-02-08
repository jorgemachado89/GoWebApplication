package main

import (
	"fmt"
	"net/http"
)

type myHandler struct {
	greeting string
}

func (myH myHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(fmt.Sprintf("%v world", myH.greeting)))
}

func main() {
	http.Handle("/", &myHandler{greeting: "Hello"})

	http.ListenAndServe(":8088", nil)
}

/*func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello Go!"))
	})

	http.ListenAndServe(":8088", nil)
}*/
