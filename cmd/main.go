package main

import (
	"github.com/gin-gonic/gin"
	"github.com/rashevskiivv/go-gin-api/pkg/books"
	"github.com/rashevskiivv/go-gin-api/pkg/common/db"
	"github.com/spf13/viper"
	"log"
)

func main() {
	viper.SetConfigFile("./pkg/common/envs/.env")
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalln(err)
	}

	port := viper.Get("PORT").(string)
	dbUrl := viper.Get("DB_URL").(string)

	r := gin.Default()
	h := db.Init(dbUrl)

	books.RegisterRoutes(r, h)

	err = r.Run(port)
	if err != nil {
		log.Fatalln(err)
	}
}
