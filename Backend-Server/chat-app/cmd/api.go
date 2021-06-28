package main

import (
	"chat-app/internal/api/router"
	"github.com/gin-gonic/gin"
	"log"
)

func main() {

	port := "5000"

	engine := gin.New()

	r := &router.Router{
		Engine: engine,
	}

	r.InitRouter()
	r.SetupHandler()

	err := engine.Run(":" + port)
	if err != nil {
		log.Fatal(err)
	}
}
