package util

import (
	"go-rest-api/model"
	"log"
	"strconv"
)

func NextId(books []model.Book) string {
	lastId := books[len(books)-1].ID

	tempId, err := strconv.Atoi(lastId)

	if err != nil {
		log.Fatal("Conversion failed. Book id is not correct")
	}

	return strconv.Itoa(tempId + 1)
}
