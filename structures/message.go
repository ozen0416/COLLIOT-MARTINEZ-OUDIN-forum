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
	Topic     Topic
	PubliDate string
}

func SendMess(m Message) (*sql.Rows, error) {
	db := database.ReturnDatabase()
	return db.Query("INSERT INTO `message`(`content`, `user_id`, `topic_id`) VALUES (?,?,?)", m.Content, m.Author, m.Topic.Id)
}

func GetMessByTopicId(id string) []Message {
	db := database.ReturnDatabase()
	rows, err := db.Query("SELECT message.id_mess, message.content, message.user_id, message.DatePost, message.topic_id, u.nickname, topic.id_user, topic.content, users.nickname from `message` inner join `users` as u on message.user_id = u.id inner join `topic` on message.topic_id = topic.id inner join `users` on topic.id_user = users.id  where topic_id = ? order by DatePost DESC", id)
	if err != nil {
		fmt.Println("GetMessByTopicId: ", err)
	}

	var messageArray []Message
	for rows.Next() {
		var messToAppend Message
		err := rows.Scan(&messToAppend.Id, &messToAppend.Content, &messToAppend.Author.Id, &messToAppend.PubliDate, &messToAppend.Topic.Id, &messToAppend.Author.Username, &messToAppend.Topic.Author.Id, &messToAppend.Topic.Content, &messToAppend.Topic.Author.Username)
		if err != nil {
			fmt.Println("Scan through topics", err)
		}
		messageArray = append(messageArray, messToAppend)
	}

	return messageArray
}
