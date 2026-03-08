package main

import (
	"redis-clone/internal/database"
	"redis-clone/internal/handler"
	"redis-clone/internal/tcp"
)



func main(){

	database := database.NewData()

	handler := handler.NewHandler(database)

	tcp := tcp.NewHandler(*handler)

	tcp.Start()
}
