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

func GetTopicsByTime() []Topic {
	db := database.ReturnDatabase()
	rows, err := db.Query("select topic.id, topic.content, topic.publi_date, users.id, users.nickname from `topic` inner join `users` on topic.id_user = users.id order by topic.publi_date DESC")
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
	_, err := db.Exec("delete from topic where id = ?", id)
	if err != nil {
		fmt.Println("Delete Topic: ", err)
	}
}
