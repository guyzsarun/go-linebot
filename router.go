package main

import (
	"github.com/gin-gonic/gin"
)

func setupRouter() *gin.Engine {
	r := gin.Default()

	r.GET("/", greeting)
	r.POST("/webhook", handler)

	return r
}
