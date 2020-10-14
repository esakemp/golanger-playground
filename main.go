package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type quote struct {
	Quote string
}

func servePong(c *gin.Context) {
	resp, err := http.Get("https://api.kanye.rest")

	if err != nil {
		log.Println("Kanye is not well", err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		log.Println("What is happening", err)
	}

	var q quote
	json.Unmarshal([]byte(body), &q)
	c.JSON(200, q)
}

func main() {
	r := gin.Default()
	r.GET("/ping", servePong)
	r.Run(":3000")
}
