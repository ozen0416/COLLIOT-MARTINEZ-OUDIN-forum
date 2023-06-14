package structures

import (
	"database/sql"
	"fmt"
	"forum/database"
)

type Message struct {
	Id        int
	Content   string
	Author    User
	TopicId   int
	PubliDate string
}

func SendMess(m Message) (*sql.Rows, error) {
	db := database.ReturnDatabase()
	return db.Query("INSERT INTO `message`(`content`, `user_id`, `topic_id`) VALUES (?,?,?)", m.Content, m.Author, m.TopicId)
}

func GetMessByTopicId(id string) []Message {
	db := database.ReturnDatabase()
	rows, err := db.Query("SELECT message.id, message.content, message.user_id, message.DatePost, message.topic_id, users.nickname FROM `message` INNER JOIN `users` on message.user_id = users.id WHERE message.topic_id = ? ORDER BY message.DatePost DESC", id)
	if err != nil {
		fmt.Println("GetMessByTopicId: ", err)
	}

	var messageArray []Message
	for rows.Next() {
		var messcToAppend Message
		err := rows.Scan(&messcToAppend.Id, &messcToAppend.Content, &messcToAppend.Author.Id, &messcToAppend.PubliDate, &messcToAppend.TopicId, &messcToAppend.Author.Username)
		if err != nil {
			fmt.Println("Scan through topics", err)
		}
		messageArray = append(messageArray, messcToAppend)
	}

	return messageArray
}
