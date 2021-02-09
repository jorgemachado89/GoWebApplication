package main

import (
	"io/ioutil"
	"log"
	"main/viewmodel"
	"net/http"
	"os"
	"text/template"
)

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

/* sfunc main() {
	templateString := `Lemonade Stand Supply`
	t, err := template.New("title").Parse(templateString)
	if err != nil {
		fmt.Println(err)
	}
	err = t.Execute(os.Stdout, nil)
	if err != nil {
		fmt.Println(err)
	}
} */

func main() {
	templates := populateTemplates()
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		requestedFile := r.URL.Path[1:] // Slice initial slash character
		t := templates[requestedFile+".html"]
		if t != nil {
			var context interface{}
			switch requestedFile {
			case "shop":
				context = viewmodel.NewShop()
			default:
				context = viewmodel.NewBase()
			}
			err := t.Execute(w, context)
			if err != nil {
				log.Println(err)
			}
		} else {
			w.WriteHeader(http.StatusNotFound)
		}
	})
	// Handle files prefixed with the following strings.
	http.Handle("/img/", http.FileServer(http.Dir("public")))
	http.Handle("/css/", http.FileServer(http.Dir("public")))

	http.ListenAndServe(":8088", nil)
}

/*func populateTemplates() *template.Template {
	result := template.New("templates")
	const basePath = "templates"
	template.Must(result.ParseGlob(basePath + "/*.html"))
	return result
}*/

func populateTemplates() map[string]*template.Template {
	result := make(map[string]*template.Template)

	const basePath = "templates"

	layout := template.Must(template.ParseFiles(basePath + "/_layout.html"))
	template.Must(layout.ParseFiles(basePath+"/_header.html", basePath+"/_footer.html"))

	dir, err := os.Open(basePath + "/content")
	if err != nil {
		panic("Failed to open template blocks directory: " + err.Error())
	}

	fis, err := dir.Readdir(-1)
	if err != nil {
		panic("Failed to read contents of content directory: " + err.Error())
	}

	for _, fi := range fis {
		f, err := os.Open(basePath + "/content/" + fi.Name())
		if err != nil {
			panic("Failed to open template '" + fi.Name() + "'")
		}

		content, err := ioutil.ReadAll(f)
		if err != nil {
			panic("Failed to read content from file '" + fi.Name() + "'")
		}

		f.Close()

		tmpl := template.Must(layout.Clone())

		_, err = tmpl.Parse(string(content))
		if err != nil {
			panic("Failed to parse contents of '" + fi.Name() + "' as template")
		}

		result[fi.Name()] = tmpl
	}

	return result
}
