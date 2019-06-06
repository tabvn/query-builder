package main

import (
	"QueryBuilder/query"
	"log"
	"time"
)
import _ "github.com/go-sql-driver/mysql"

func main() {

	db := query.NewDatabase("root:root@/query_builder")

	// INSERT INTO `user` (`id`, `name`, `created_at`) VALUES (?, ?, ?);
	result, err := db.Insert("user").Fields(map[string]interface{}{
		"name":       "Tom Nguyen",
		"created_at": time.Now(),
	}).Execute()

	id, err := result.LastInsertId()

	log.Println("Last insert ID: ", id, " error:", err)

}
