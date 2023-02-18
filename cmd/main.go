package main

import (
	"github.com/gin-gonic/gin"
	"github.com/rashevskiivv/go-gin-api/pkg/books"
	"github.com/rashevskiivv/go-gin-api/pkg/common/configs"
	"log"
)

func main() {
	log.Println("Server started")
	err := configs.LoadConfig()
	if err != nil {
		log.Fatalln("Failed at config .env", err)
	}

	r := gin.Default()
	log.Println("Router created")

	err = configs.ConnectDB()
	if err != nil {
		log.Fatalln(err)
	}

	books.RegisterRoutes(r)

	err = r.Run(":" + configs.GetConfig().Port)
	if err != nil {
		log.Fatalln(err)
	}
}
