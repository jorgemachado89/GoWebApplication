package controller

import (
	"html/template"
	"net/http"
	"main/viewmodel"
	"fmt"
	"log"
)

type home struct {
	homeTemplate			*template.Template
	standLocatorTemplate	*template.Template
	loginTemplate			*template.Template
}

func (h home) registerRoutes() {
	http.HandleFunc("/", h.handleHome)
	http.HandleFunc("/home", h.handleHome)
	http.HandleFunc("/stand-locator", h.handleStandLocator)
	http.HandleFunc("/login", h.handleLogin)
}

func (h home) handleHome(w http.ResponseWriter, r *http.Request) {
	vm := viewmodel.NewHome()
	h.homeTemplate.Execute(w, vm)
}

func (h home) handleStandLocator(w http.ResponseWriter, r *http.Request) {
	vm := viewmodel.NewStandLocator()
	h.standLocatorTemplate.Execute(w, vm)
}

func (h home) handleLogin(w http.ResponseWriter, r *http.Request) {
	vm := viewmodel.NewLogin();
	fmt.Printf("Start")
	if r.Method == http.MethodPost {
		err := r.ParseForm();
		if err != nil {
			log.Println(fmt.Errorf("Error logging in %v: ", err));
		}
		email := r.Form.Get("email")
		password := r.Form.Get("password")
		if (email == "test@gmail.com" && password == "password") {
			http.Redirect(w, r, "/home", http.StatusTemporaryRedirect)
			return
		} else {
			vm.Email = email
			vm.Password = password
		}
		// fmt.Printf("Initial Form: %v \n", r.Form)
		// fmt.Printf("Username: %v \n", r.FormValue("email"))
		// fmt.Printf("Form: %v \n", r.Form)
	}
	h.loginTemplate.Execute(w, vm);
}
