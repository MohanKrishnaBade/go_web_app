package Auth

import (
	"net/http"
	"github.com/go_web_app/db"
	"regexp"
	"database/sql"
	"golang.org/x/crypto/bcrypt"
	"github.com/go_web_app/tpl"
	"fmt"
	"strconv"
	"math/rand"
	"github.com/go_web_app/models"
	"github.com/go_web_app/Errors"
	"github.com/satori/go.uuid"
	"time"
)

func login(w http.ResponseWriter, req *http.Request) {
	if alreadyLoggedIn(w, req) {
		http.Redirect(w, req, "/", http.StatusSeeOther)
		return
	}
	//var u user
	// process form submission
	if req.Method == http.MethodPost {
		e := req.FormValue("email")
		p := req.FormValue("password")

		var l Errors.Login

		if !VLogin(req, &l) {
			tpl.Tpl.ExecuteTemplate(w, "login.html", l)
			return
		}

		u, er := db.FindByEmail(e)
		if er != nil {
			http.Redirect(w, req, "/signup", http.StatusSeeOther)
			return
		}
		// does the entered password match the stored password?
		err := bcrypt.CompareHashAndPassword([]byte(u.Password.String), []byte(p))
		if err != nil {
			l.Auth = "Username and/or password do not match"
			tpl.Tpl.ExecuteTemplate(w, "login.html", l)

			return
		}
		// create session
		helper(w, req, u)
	} else {
		showSessions() // for demonstration purposes
		http.Redirect(w, req, "/", http.StatusSeeOther)
	}
}

func logout(w http.ResponseWriter, req *http.Request) {
	c, _ := req.Cookie("session")
	// delete the session
	delete(dbSessions, c.Value)
	// remove the cookie
	c = &http.Cookie{
		Name:   "session",
		Value:  "",
		MaxAge: -1,
	}
	http.SetCookie(w, c)

	// clean up dbSessions
	if time.Now().Sub(dbSessionsCleaned) > (time.Second * 30) {
		go cleanSessions()
	}

	http.Redirect(w, req, "/login", http.StatusSeeOther)
}

func signup(w http.ResponseWriter, req *http.Request) {
	if alreadyLoggedIn(w, req) {
		http.Redirect(w, req, "/", http.StatusSeeOther)
		return
	}
	// process form submission
	if req.Method == http.MethodPost {
		// get form values
		e := req.FormValue("email")
		p := req.FormValue("password")
		f := req.FormValue("firstName")
		l := req.FormValue("lastName")

		var s Errors.SignIn
		if !VSignUp(req, &s) {
			tpl.Tpl.ExecuteTemplate(w, "register.html", s)
			return
		}

		// store user in dbUsers
		bs, err := bcrypt.GenerateFromPassword([]byte(p), bcrypt.DefaultCost)
		if err != nil {
			http.Error(w, "Internal server error", http.StatusInternalServerError)
			return
		}
		u := models.User{
			e,
			strconv.Itoa(rand.Int()),
			true,
			profileCollection[rand.Intn(len(profileCollection))],
			sql.NullString{f, true},
			sql.NullString{l, true},
			sql.NullString{string(bs), true},
		}

		//dbUsers[e] = u
		db.CreateUser(u)

		// redirect
		http.Redirect(w, req, "/", http.StatusSeeOther)
		return
	} else {
		showSessions() // for demonstration purposes
		tpl.Tpl.ExecuteTemplate(w, "register.html", nil)
	}
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {

	//c, _ := Utilities.GetCookie(r, "user");
	var dashboard models.Dashboard

	dashboard.User = db.FindByGId(GetCurrentUser(w, r).GId)

	re := regexp.MustCompile(emailReex)

	if (re.MatchString(r.FormValue("email"))) {
		dashboard.User.FirstName = sql.NullString{r.FormValue("firstName"), true}
		dashboard.User.LastName = sql.NullString{r.FormValue("lastName"), true}
		dashboard.User.Email = r.FormValue("email")
		bs, _ := bcrypt.GenerateFromPassword([]byte(r.FormValue("password")), bcrypt.DefaultCost)
		dashboard.User.Password = sql.NullString{string(bs), true}

		db.Update(dashboard.User)
		tpl.Tpl.ExecuteTemplate(w, "user-profile.html", dashboard)
	} else {
		fmt.Println("email is not valid")
		tpl.Tpl.ExecuteTemplate(w, "user-profile.html", dashboard)
	}
}

func getProfile(w http.ResponseWriter, r *http.Request) {

	var dashboard models.Dashboard
	dashboard.User = db.FindByGId(GetCurrentUser(w, r).GId);
	tpl.Tpl.ExecuteTemplate(w, "user-profile.html", dashboard)
}

func helper(w http.ResponseWriter, req *http.Request, u models.User) {
	// create session
	sID:= uuid.NewV4()
	c := &http.Cookie{
		Name:  "session",
		Value: sID.String(),
	}
	c.MaxAge = sessionLength
	http.SetCookie(w, c)
	dbSessions[c.Value] = session{u.Email, time.Now()}
	dbUsers[u.Email] = u
	http.Redirect(w, req, "/dashboard", http.StatusSeeOther)
	return
}