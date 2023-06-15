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
	return db.Query("INSERT INTO `message`(`content`, `id_user`, `id_topic`) VALUES (?,?,?)", m.Content, m.Author.Id, m.Topic.Id)
}

func GetMessByTopicId(id string) []Message {
	db := database.ReturnDatabase()
	// Get mess + author of mess + topic of mess + author of topic
	rows, err := db.Query("SELECT message.id_message, message.content, message.id_user, message.Date_Post, message.id_topic, u.username, topic.id_user, topic.content, users.username from `message` inner join `users` as u on message.id_user = u.id_user inner join `topic` on message.id_topic = topic.id_topic inner join `users` on topic.id_user = users.id_user  where message.id_topic = ? order by Date_Post DESC", id)
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
