package main

import (
	"os"

	"github.com/Pratap2018/go-lang-server/controller"
	HealthCheck "github.com/Pratap2018/go-lang-server/controller"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		panic(err)
	}

	server := gin.Default()
	server.GET("/", HealthCheck.HealthVar.Health)
	server.POST("/sign", controller.AuxController)
	server.Run(":" + (os.Getenv("PORT")))
}
