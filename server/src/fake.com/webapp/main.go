package main

import (
	"html/template"
	"io/ioutil"
	"net/http"
	"os"
	"log"
	"fmt"
	_ "github.com/lib/pq"
	"database/sql"
	"fake.com/webapp/controller"
	"fake.com/webapp/middleware"
	"fake.com/webapp/model"
)

func main() {
	templates := populateTemplates()

	db := connectDB()
	defer db.Close()

	controller.Startup(templates)

	http.ListenAndServe(":8000", &middleware.TimeoutMiddleware{Next: new (middleware.GzipMiddleware)})
}

func connectDB() *sql.DB {
	connStr := "postgres://lss:lss@localhost/lss?sslmode=disable"
	
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatalln(fmt.Errorf("Unable to connect to Database: %v", err))
	}

	model.SetDatabase(db)
	return db;
}

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
