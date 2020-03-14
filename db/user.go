package db

import (
	"github.com/go_web_app/Errors"
	"fmt"
	"github.com/go_web_app/models"
)

func Insert(user models.User) models.User {

	fmt.Println(user.VerifiedEmail)
	if user.VerifiedEmail {

		db := Connection()

		rows, err := db.Query("SELECT * FROM users WHERE id=?", user.GId)
		Errors.PanicError(err)

		if !rows.Next() {
			insForm, err := db.Prepare("INSERT INTO  users (email, id, picture) VALUES(?,?,?)")

			_, err = insForm.Exec(user.Email, user.GId, user.Picture)

			Errors.PanicError(err)
		}

		defer db.Close();
	}
	return user
}

func CreateUser(user models.User) {

	db := Connection()

	insForm, err := db.Prepare("INSERT INTO  users (email, id, firstName, lastName, password, picture) VALUES(?,?,?,?,?,?)")

	_, err = insForm.Exec(user.Email, user.GId, user.FirstName.String, user.LastName.String, user.Password.String, user.Picture)
	Errors.PanicError(err)

	db.Close()
}
func FindByGId(gId string) models.User {
	db := Connection()

	rows, err := db.Query("SELECT email, id, picture, firstName, lastName FROM users WHERE id=?", gId)
	var user models.User
	for rows.Next() {

		err = rows.Scan(&user.Email, &user.GId, &user.Picture, &user.FirstName, &user.LastName)
		if err != nil {
			fmt.Printf("Scan: %v", err)
		}
	}
	defer db.Close();
	fmt.Println(user)
	return user
}

func Update(user models.User) {
	db := Connection()
	_, err := db.Query("update users set email=? , firstName=? ,lastName=?, picture=?,id=? where id =? or email=?", user.Email, user.FirstName.String, user.LastName.String, user.Picture, user.GId, user.GId, user.Email)
	Errors.PanicError(err)
	defer db.Close();
}

func FindByEmail(email string) (models.User, error) {
	db := Connection()

	rows, err := db.Query("SELECT email, id, picture, firstName, lastName, password FROM users WHERE email=?", email)
	var user models.User
	for rows.Next() {
		err = rows.Scan(&user.Email, &user.GId, &user.Picture, &user.FirstName, &user.LastName, &user.Password)
		if err != nil {
			return user, &Errors.Error{}
		}
	}

	if user.GId == "" {
		return user, &Errors.Error{}
	}
	defer db.Close();
	return user, nil
}

func Merge(gUser models.User, user2 models.User) (models.User) {
	var user models.User

	user.Email = gUser.Email
	user.GId = gUser.GId
	user.Picture = gUser.Picture
	user.FirstName = user2.FirstName
	user.LastName = user2.LastName
	user.Password = user2.Password

	return user
}
