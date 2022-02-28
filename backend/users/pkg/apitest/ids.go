package apitest

import (
	"github.com/gocql/gocql"
)

func getUserIDs() []gocql.UUID {
	arr := []gocql.UUID{}
	for i := range TestUsers {
		arr = append(arr, TestUsers[i].UserID)
	}
	return arr
}

func SetUserIDs() {
	userIDs := getUserIDs()
	for i := range TestUsers {
		TestUsers[i].UserID = userIDs[i]
	}
}
