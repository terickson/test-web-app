package main

import (
	"net/http"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

var Port = 8000

func main() {
	log := logrus.New()
	log.SetLevel(logrus.DebugLevel)

	router := gin.Default()

	router.LoadHTMLGlob("web/templates/**/*")
	baseRouter := router.Group("")
	hostname, _ := os.Hostname()

	baseRouter.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "main", gin.H{"host": hostname, "headers": c.Request.Header})
	})

	log.Infoln("Serving on http://127.0.0.1:" + strconv.Itoa(Port))
	if err := router.Run("0.0.0.0:" + strconv.Itoa(Port)); err != nil {
		log.Fatalln(err)
	}
}
