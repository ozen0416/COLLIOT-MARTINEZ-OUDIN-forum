package structures

import (
	"database/sql"
)

type Message struct {
	Id        int
	Content   string
	AuthorId  int
	TopicId   int
	PubliDate string
}

func SendMess(db *sql.DB, m Message) (*sql.Rows, error) {
	return db.Query("INSERT INTO `message`(`content`, `user_id`, `topic_id`) VALUES (?,?,?)", m.Content, m.AuthorId, m.TopicId)
}
