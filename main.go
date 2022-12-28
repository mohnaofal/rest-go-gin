package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/mohnaofal/rest-go-gin/app/delivery"
	"github.com/mohnaofal/rest-go-gin/config"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// load config
	cfg := config.LoadConfig()

	// init Gin
	r := gin.Default()
	// Global middleware
	// Logger middleware will write the logs to gin.DefaultWriter even if you set with GIN_MODE=release.
	// By default gin.DefaultWriter = os.Stdout
	r.Use(gin.Logger())

	// Recovery middleware recovers from any panics and writes a 500 if there was one.
	r.Use(gin.Recovery())

	articleDelivery := delivery.NewArticleDelivery(cfg)
	articleGroups := r.Group("articles")
	articleDelivery.Apply(articleGroups)

	r.Run() // listen and serve on 0.0.0.0:8080
}
