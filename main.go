package main

import "log"
func main() {

	db, err:= CreateDb()
	if err != nil {
		log.Fatal(err)
	}

	server := NewApiServer(":8000", db)
	server.Run()
}