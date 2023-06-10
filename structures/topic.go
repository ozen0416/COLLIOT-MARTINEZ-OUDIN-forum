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
}

func GetTopicsByTime() []Topic {
	db := database.ReturnDatabase()
	rows, err := db.Query("select * from `topic` order by topic.publi_date DESC")
	if err != nil {
		fmt.Println("Login: ", err)
	}
	var topicArray []Topic
	for rows.Next() {
		var topicToAppend Topic
		err := rows.Scan(&topicToAppend.Id, &topicToAppend.CatId, &topicToAppend.Content, &topicToAppend.PubliDate)
		if err != nil {
			fmt.Println("Scan through topics", err)
		}
		topicArray = append(topicArray, topicToAppend)
	}
	return topicArray
}
