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
	_, err := db.Query("INSERT INTO `users`(`username`, `email`, `birth_date`, `passwd`) VALUES (?,?,?,?)", u.Username, u.Email, u.BirthDate, u.Pass)
	var id int
	_ = db.QueryRow("select id_user from users where email = ?", u.Email).Scan(&id)
	return id, err
}

func Login(email, pass string) (int, bool) {
	db := database.ReturnDatabase()
	var hashPass string
	var idUser int
	err := db.QueryRow("SELECT id_user, passwd FROM `users` WHERE users.email = ?", email).Scan(&idUser, &hashPass)
	if err == sql.ErrNoRows {
		return 0, false
	}
	return idUser, secret.CheckPasswordHash(pass, hashPass)
}

func DeleteUser(id int) {
	db := database.ReturnDatabase()
	_, err := db.Exec("delete from users where id_user = ?", id)
	if err != nil {
		fmt.Println("Delete User: ", err)
	}
}
