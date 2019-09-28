package main

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/themusicman/echo/echo"
)

var log = logrus.New()

func main() {
	log.Out = os.Stdout
	router := gin.New()
	router.Use(gin.Recovery())
	router.GET("/echo", echo.HandleGet(log))
	router.Run(":8080")
}
