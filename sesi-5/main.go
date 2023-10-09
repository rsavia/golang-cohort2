package main

import (
	"net/http"
	"html/template"
)

type Employee struct {
	Email string
	Address string
	Job string
	Reason string
}

var PORT = ":9090"
var DATA = []Employee{
	{Email:"reni@yahoo.com", Address:"bekasi", Job:"product analyst", Reason:"keren"},
	{Email:"manda@yahoo.com", Address:"cinere", Job:"head of data", Reason:"mantap"},
	{Email:"nadya@yahoo.com", Address:"deplu", Job:"ux researcher", Reason:"cepat"},
	{Email:"jilly@yahoo.com", Address:"bintaro", Job:"software engineer", Reason:"simpel"},
	{Email:"dizty@yahoo.com", Address:"nusaloka", Job:"lead qa", Reason:"ikut jilly"},
	{Email:"triisya@yahoo.com", Address:"karang tengah", Job:"qa", Reason:"ikut dizty"},
}

func main() {
	http.HandleFunc("/", redirectToLogin)
	http.HandleFunc("/login", login)
	http.HandleFunc("/users", showUser)
	http.HandleFunc("/errorLogin", showError)
	
	http.ListenAndServe(PORT, nil)
}

func redirectToLogin(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/login", http.StatusSeeOther)
}

func login(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		var data = map[string]interface{}{"data": DATA}

		renderTemplate(w, "login.html", data)
	} else if r.Method == http.MethodPost {
		r.ParseForm()
		var email, success = authenticate(r.Form.Get("email"), r.Form.Get("password"))
		if success {
			http.Redirect(w, r, "/users?email=" + email, http.StatusSeeOther)
		} else {
			http.Redirect(w, r, "/errorLogin", http.StatusSeeOther)
		}
	} else {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}
	
	return
}

func authenticate(email string, password string) (string, bool) {
	found := false
	for _, value := range DATA {
		if value.Email == email {
			found = true
			break
		}
	}

	return email, found
}

func showUser(w http.ResponseWriter, r *http.Request) {
	email := r.URL.Query().Get("email")
	var employee Employee;
	for _, value := range DATA {
                if value.Email == email {
                	employee = value;
		        break
                }
        }
	var data = map[string]interface{}{"data": employee}
	renderTemplate(w, "show_user.html", data)
}

func showError(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "show_error.html", nil)
}

func renderTemplate(w http.ResponseWriter, templateName string, data map[string]interface{}) {
	 var tmpl, err = template.ParseFiles(templateName)
         if err != nil {
         	http.Error(w, err.Error(), http.StatusInternalServerError)
               	return
         }

       	err = tmpl.Execute(w, data)
        if err != nil {
        	http.Error(w, err.Error(), http.StatusInternalServerError)
       		return
	}
}
