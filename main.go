package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"log"
	"time"
)

func main() {
	r := gin.Default()
	r.POST("/api/user-service/login", func(c *gin.Context) {

		fmt.Println(c.Request.RequestURI)
		bs, err :=  ioutil.ReadAll(c.Request.Body)
		if err != nil {
			log.Fatalln(err)
		}
		fmt.Println("bs:", string(bs))
		time.Sleep(40*time.Second)
		c.String(200, "Hello, Geektutu")
	})
	r.Run(":5001") // listen and serve on 0.0.0.0:8080
}
