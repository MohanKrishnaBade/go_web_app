package models

import "database/sql"

// This is the type we define for deserialization.
// You can use map[string]string as well
type User struct {
	Email         string `json:"email"`
	GId           string `json:"id"`
	VerifiedEmail bool   `json:"verified_email"`
	Picture       string `json:"picture"`
	FirstName     sql.NullString
	LastName      sql.NullString
	Password      sql.NullString
}
