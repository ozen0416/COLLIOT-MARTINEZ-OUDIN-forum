package structures

import (
	"fmt"
	"forum/database"
)

type Topic struct {
	Id        int
	CatId     int
	Content   string
	PubliDate string
	Author    User
}

func SendTopic(t Topic) error {
	db := database.ReturnDatabase()
	_, err := db.Query("insert into `topic`(`content`, `id_category`, `id_user`) values (?,?,?)", t.Content, t.CatId, t.Author.Id)
	if err != nil {
		fmt.Println("Insert Topic: ", err)
	}
	var idTopic int
	_ = db.QueryRow("select id_topic from topic where content = ? and id_user = ?", t.Content, t.Author.Id).Scan(&idTopic)
	_, err = db.Exec("insert into `message` (content, id_user, id_topic) VALUES (?, ?, ?)", t.Content, t.Author.Id, idTopic)
	if err != nil {
		fmt.Println("Insert first mess of topic: ", err)
	}
	return err
}

func GetTopicsByTime() []Topic {
	db := database.ReturnDatabase()
	rows, err := db.Query("select topic.id_topic, topic.content, topic.publication_date, users.id_user, users.username from `topic` inner join `users` on topic.id_user = users.id_user order by topic.publication_date DESC")
	if err != nil {
		fmt.Println("Login: ", err)
	}
	var topicArray []Topic
	for rows.Next() {
		var topicToAppend Topic
		err := rows.Scan(&topicToAppend.Id, &topicToAppend.Content, &topicToAppend.PubliDate, &topicToAppend.Author.Id, &topicToAppend.Author.Username)
		if err != nil {
			fmt.Println("Scan through topics", err)
		}
		topicArray = append(topicArray, topicToAppend)
	}
	return topicArray
}

func DeleteTopic(id int) {
	db := database.ReturnDatabase()
	_, err := db.Exec("delete from topic where id_topic = ?", id)
	if err != nil {
		fmt.Println("Delete Topic: ", err)
	}
}
