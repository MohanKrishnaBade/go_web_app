package Auth

import (
	"net/http"
	"github.com/go_web_app/Errors"
	"regexp"
	"github.com/go_web_app/db"
)

func VLogin(req *http.Request, l *Errors.Login) (bool) {
	e := req.FormValue("email")
	p := req.FormValue("password")
	isValid := true
	if e == "" {
		l.Email = "Email should not be empty"
		isValid = false
	} else if !regexp.MustCompile(emailReex).MatchString(e) {
		l.Email = "Email is not valid"
		isValid = false
	}
	if p == "" {
		l.Password = "password should not be empty"
		isValid = false
	}
	return isValid
}

func VSignUp(req *http.Request, s *Errors.SignIn) (bool) {
	e := req.FormValue("email")
	p := req.FormValue("password")

	isValid := true
	if e == "" {
		s.Email = "Email should not be empty"
		isValid = false
	} else if !regexp.MustCompile(emailReex).MatchString(e) {
		s.Email = "Email is not valid, try with different one."
		isValid = false
	} else {
		_, er := db.FindByEmail(e)
		if er == nil {
			s.Email = "Email already taken"
			isValid = false
		}
	}
	if p == "" {
		s.Password = "password should not be empty"
		isValid = false
	}

	return isValid
}
