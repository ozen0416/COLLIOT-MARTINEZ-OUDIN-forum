package structures

import (
	"database/sql"
	"fmt"
	"forum/database"
	"forum/secret"
)

type User struct {
	Id        int
	Username  string
	Email     string
	Pass      string
	BirthDate string
}

func SignUser(u User) (*sql.Rows, error) {
	db := database.ReturnDatabase()
	return db.Query("INSERT INTO `users`(`nickname`, `email`, `Date_Birth`, `passwd`) VALUES (?,?,?,?)", u.Username, u.Email, u.BirthDate, u.Pass)
}

func Login(email, pass string) bool {
	db := database.ReturnDatabase()
	var hashPass string
	err := db.QueryRow("SELECT passwd FROM `users` WHERE users.email = ?", email).Scan(&hashPass)
	if err != nil {
		fmt.Println("Login: ", err)
	}
	return secret.CheckPasswordHash(pass, hashPass)
}
