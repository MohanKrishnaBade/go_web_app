package main

import (
	"github.com/gorilla/mux"
	"net/http"
	"fmt"
	"github.com/go_web_app/tpl"
	"github.com/go_web_app/Auth"
	"github.com/go_web_app/models"
	"github.com/go_web_app/Chat"
)

func init() {
	Setenv()
	Auth.SetConfig()
}

func main() {

	r := mux.NewRouter()

	Auth.Routes(r)
	Chat.Routes(r)
	r.HandleFunc("/dashboard", index)
	r.PathPrefix("/").Handler(http.FileServer(http.Dir("./frontend/")))

	fmt.Println(http.ListenAndServe(":80", r))
}

func index(w http.ResponseWriter, req *http.Request) {
	var i models.Dashboard
	i.User = Auth.GetCurrentUser(w,req)
	tpl.Tpl.ExecuteTemplate(w, "dashboard.html", i)
}
