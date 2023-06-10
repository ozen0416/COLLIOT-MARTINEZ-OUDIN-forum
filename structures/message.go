package structures

import (
	"database/sql"
	"fmt"
	"forum/database"
)

type Message struct {
	Id        int
	Content   string
	AuthorId  int
	TopicId   int
	PubliDate string
}

func SendMess(m Message) (*sql.Rows, error) {
	db := database.ReturnDatabase()
	return db.Query("INSERT INTO `message`(`content`, `user_id`, `topic_id`) VALUES (?,?,?)", m.Content, m.AuthorId, m.TopicId)
}

func GetMessByTopicId(id string) []Message {
	db := database.ReturnDatabase()
	rows, err := db.Query("select * from `message` where topic_id = ? order by DatePost DESC", id)
	if err != nil {
		fmt.Println("GetMessByTopicId: ", err)
	}

	var messageArray []Message
	for rows.Next() {
		var messcToAppend Message
		err := rows.Scan(&messcToAppend.Id, &messcToAppend.Content, &messcToAppend.AuthorId, &messcToAppend.PubliDate, &messcToAppend.TopicId)
		if err != nil {
			fmt.Println("Scan through topics", err)
		}
		messageArray = append(messageArray, messcToAppend)
	}

	return messageArray
}
