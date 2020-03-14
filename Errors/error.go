package Errors

type Error struct {
	Login Login
	prob  string
}

func (e *Error) Error() string {
	return e.prob;
}

type Login struct {
	Email    string
	Password string
	Auth     string
}

type SignIn struct {
	Email     string
	FirstName string
	LastName  string
	Password  string
	Auth      string
}
