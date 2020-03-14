package Auth

import (
	"net/http"
	"time"
	"github.com/gorilla/mux"
	"github.com/go_web_app/models"
)

const sessionLength int = 60

var emailReex = "^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$"
var profileCollection [4] string
var dbUsers = map[string]models.User{} // user ID, user
var dbSessions = map[string]session{}  // session ID, session
var dbSessionsCleaned time.Time

type session struct {
	un           string
	lastActivity time.Time
}

func init() {
	dbSessionsCleaned = time.Now()
	profileCollection[0] = "https://www.pngkey.com/png/detail/412-4127687_privileged-listen-in-go-golang-hamster.png"
	profileCollection[1] = "https://www.pngkey.com/png/detail/412-4126781_gopher-by-rene-french-and-tenntenn-gophers-golang.png"
	profileCollection[2] = "https://www.pinclipart.com/picdir/middle/57-579569_go-lang-gopher-clipart.png"
	profileCollection[3] = "https://www.pngkey.com/png/detail/71-717551_winnie-the-pooh-gopher-coming-out-of-his.png"
}

func Routes(r *mux.Router) {
	r.HandleFunc("/signup", signup).Methods(http.MethodGet, http.MethodPost)
	r.HandleFunc("/login", login)
	r.HandleFunc("/logout", authorized(logout))
	r.HandleFunc("/profile", getProfile).Methods(http.MethodGet)
	r.HandleFunc("/GLogin", HandleGoogleLogin).Methods(http.MethodGet)
	r.HandleFunc("/callback", handleGoogleCallback).Methods(http.MethodGet)
	r.HandleFunc("/updateUser", UpdateUser).Methods(http.MethodPost)

	r.Handle("/favicon.ico", http.NotFoundHandler())
}

func authorized(h http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if !alreadyLoggedIn(w, r) {
			//http.Error(w, "not logged in", http.StatusUnauthorized)
			http.Redirect(w, r, "/", http.StatusSeeOther)
			return // don't call original handler
		}
		h.ServeHTTP(w, r)
	})
}
