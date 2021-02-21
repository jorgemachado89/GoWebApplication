package controller

import (
	"fmt"
	"html/template"
	"log"
	"net/http"

	"fake.com/webapp/model"
	"fake.com/webapp/viewmodel"
)

type home struct {
	homeTemplate  *template.Template
	loginTemplate *template.Template
}

func (h home) registerRoutes() {
	http.HandleFunc("/", h.handleHome)
	http.HandleFunc("/home", h.handleHome)
	http.HandleFunc("/login", h.handleLogin)
}

func validateServerPush(w http.ResponseWriter, r *http.Request) {
	if pusher, ok := w.(http.Pusher); ok {
		pusher.Push("/css/app.min.css", &http.PushOptions{
			Header: http.Header{"Content-Type": []string{"text/css"}},
		})
	}
}

func (h home) handleHome(w http.ResponseWriter, r *http.Request) {
	validateServerPush(w, r)
	vm := viewmodel.NewHome()
	h.homeTemplate.Execute(w, vm)
}

func getFormData(r *http.Request) (username string, password string) {
	return r.Form.Get("email"), r.Form.Get("password")
}

func (h home) handleLogin(w http.ResponseWriter, r *http.Request) {
	vm := viewmodel.NewLogin()
	if r.Method == http.MethodPost {

		err := r.ParseForm()
		if err != nil {
			log.Println(fmt.Errorf("Error logging in %v: ", err))
		}

		username, password := getFormData(r)

		if user, err := model.Login(username, password); err == nil {
			log.Printf("User %s has logged in successfuly.\n", user.GetUsername())
			http.Redirect(w, r, "/home", http.StatusTemporaryRedirect)
			return
		} else {
			log.Printf("Error validating User: %s - Password: %s | crendentials - %v \n", username, password, err)
			vm.Email = username
			vm.Password = password
		}
		// fmt.Printf("Initial Form: %v \n", r.Form)
		// fmt.Printf("Username: %v \n", r.FormValue("email"))
		// fmt.Printf("Form: %v \n", r.Form)
	}
	w.Header().Add("Content-Type", "text/html")
	h.loginTemplate.Execute(w, vm)
}
