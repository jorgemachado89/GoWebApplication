package main

import "net/http"

/* type myHandler struct {
	greeting string
} */

/* func (myH myHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(fmt.Sprintf("%v world", myH.greeting)))
}*/

/*func main() {
	http.Handle("/", &myHandler{greeting: "Hello"})

	http.ListenAndServe(":8088", nil)
}*/

/*func handleContentType(path string) string {
	var contentType string

	switch {
	case strings.HasSuffix(path, "css"):
		contentType = "text/css"
	case strings.HasSuffix(path, "html"):
		contentType = "text/html"
	case strings.HasSuffix(path, "png"):
		contentType = "image/png"
	default:
		contentType = "text/plain"
	}

	return contentType
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// w.Write([]byte("Hello Go!"))

		f, err := os.Open("public" + r.URL.Path)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			log.Println(err)
		}

		defer f.Close()

		w.Header().Add("Content-Type", handleContentType(r.URL.Path))

		io.Copy(w, f)
	})

	http.ListenAndServe(":8088", nil)
}*/

/*func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "public"+r.URL.Path)
	})

	http.ListenAndServe(":8088", nil)
}*/

func main() {
	http.ListenAndServe(":8088", http.FileServer(http.Dir("public")))
}
