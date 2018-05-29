package main

import (
	"fmt"
	"html/template"
	"math/rand"
	"net/http"
	"time"
)

type Vehicle struct {
	Name  string
	Count int
}

type Vehicles struct {
	List []*Vehicle
}

type View struct {
	Username string
	Vehicles Vehicles
}

var joinT = template.Must(template.ParseFiles("templates/join.html"))
var playT = template.Must(template.ParseFiles("templates/play.html"))

var view = View{"", Vehicles{make([]*Vehicle, 0)}}

func inOneYear() time.Time {
	return time.Now().AddDate(1, 0, 0)
}

func join(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		joinT.Execute(w, nil)
	} else {
		r.ParseForm()
		username := r.Form.Get("username")
		cookie := http.Cookie{Name: "username", Value: username, Expires: inOneYear()}
		http.SetCookie(w, &cookie)
		view = View{"", Vehicles{make([]*Vehicle, 0)}}
		http.Redirect(w, r, "/play", 303)
	}
}

func play(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("username")
	if err != nil {
		fmt.Fprintf(w, "Could not find cookie named 'username'")
		return
	}

	view.Username = cookie.Value
	if r.Method == "GET" {
		playT.Execute(w, view)
	}
}

func add(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		playT.Execute(w, view)
	} else {
		r.ParseForm()
		vehicle := r.Form.Get("vehicle")
		speed := r.Form.Get("speed")

		name := vehicle + ": " + speed

		temp := Vehicle{name, rand.Intn(10)}
		view.Vehicles.List = append(view.Vehicles.List, &temp)

		http.Redirect(w, r, "/play", 303)
	}
}

func main() {
	http.HandleFunc("/join", join)
	http.HandleFunc("/play", play)
	http.HandleFunc("/add", add)
	http.HandleFunc("/exit", join)
	http.Handle("/public/", http.StripPrefix("/public/", http.FileServer(http.Dir("public"))))
	http.ListenAndServe("localhost:8080", nil)
}
