package db

import (
	_ "github.com/go-sql-driver/mysql"
	"database/sql"
	"encoding/json"
	"os/exec"
	"fmt"
	"github.com/go_web_app/models"
)

func Connection() (db *sql.DB) {
	dbDriver := "mysql"
	dbUser := "root"
	dbPass := "r00tmysql"
	dbName := "go_project"
	db, err := sql.Open(dbDriver, dbUser+":"+dbPass+"@tcp(mysql:3306)/"+dbName)
	if err != nil {
		panic(err.Error())
	}
	return db
}

// This functions accepts a byte array containing a JSON
func ParseUser(jsonBuffer []byte) (models.User, error) {

	var user models.User

	// Unmarshal the json into it. this will use the struct tag
	err := json.Unmarshal(jsonBuffer, &user)

	// the array is now filled with users
	return user, err
}

func Import() {
	//cmdStr := "docker exec docker_mysql_1 mysql -u root -pr00tmysql -e crate database auth;"
	//err := exec.Command("/bin/sh", "-c", cmdStr).Run()

	cmd := exec.Command("/bin/sh", "docker_mysql_1",
		"-uroot", "-pr00tmysql", "-Dauth",
		"-e", "source /go/src/github.com/webApp/test.sql")
	err := cmd.Run()
	if err != nil {
		fmt.Println(err.Error())
	}
}
