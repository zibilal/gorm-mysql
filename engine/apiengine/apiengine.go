package apiengine

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

type ApiEngine struct {
	Engine *gin.Engine
}

func InitApiEngine() *ApiEngine {
	a := new(ApiEngine)

	e := gin.New()
	e.Use(gin.Logger())
	e.Use(gin.Recovery())
	e.NoRoute(func(c *gin.Context){
		c.JSON(404, gin.H{
			"code": "NOT_FOUND",
			"message": fmt.Sprintf("Endpoint %s Not Found", c.Request.URL.Path),
		})
	})

	a.Engine = e
	return a
}


