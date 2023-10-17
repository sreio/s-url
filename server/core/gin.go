package core

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/google/martian/log"
	"github.com/sreio/s-url/define"
)

func InitGin() {
	if !define.Config.Server.Debug {
		gin.SetMode(gin.ReleaseMode)
	}

	r := gin.New()
	startRouter(r)
	err := r.Run(fmt.Sprintf("%s:%d", define.Config.Server.Host, define.Config.Server.Port))
	if err != nil {
		log.Errorf("server run error: %s", err.Error())
	}
}

func startRouter(c *gin.Engine) {
	router := c.Group("/api")
	router.GET("/ping", func(c *gin.Context) {

		c.JSON(200, gin.H{
			"ping":   "pong",
			"config": define.Config,
		})

	})
}
