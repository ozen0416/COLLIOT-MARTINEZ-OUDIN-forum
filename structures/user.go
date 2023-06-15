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

func SignUser(u User) (int, error) {
	db := database.ReturnDatabase()
	_, err := db.Query("INSERT INTO `users`(`nickname`, `email`, `Date_Birth`, `passwd`) VALUES (?,?,?,?)", u.Username, u.Email, u.BirthDate, u.Pass)
	var id int
	_ = db.QueryRow("select id from users where email = ?", u.Email).Scan(&id)
	return id, err
}

func Login(email, pass string) (int, bool) {
	db := database.ReturnDatabase()
	var hashPass string
	var idUser int
	err := db.QueryRow("SELECT id, passwd FROM `users` WHERE users.email = ?", email).Scan(&idUser, &hashPass)
	if err == sql.ErrNoRows {
		return 0, false
	}
	return idUser, secret.CheckPasswordHash(pass, hashPass)
}

func DeleteUser(id int) {
	db := database.ReturnDatabase()
	_, err := db.Exec("delete from users where id = ?", id)
	if err != nil {
		fmt.Println("Delete User: ", err)
	}
}
