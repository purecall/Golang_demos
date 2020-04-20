package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// gin.H is a shortcut for map[string]interface{}
	r.GET("/someJSON", func(c *gin.Context) {
		//第一种方式，自己拼json
		c.JSON(http.StatusOK, gin.H{"message": "hey", "status": http.StatusOK})
	})

	r.GET("/moreJSON", func(c *gin.Context) {
		// You can also use a struct
		var msg struct {
			Name    string `json:"user"`
			Message string
			Number  int
		}
		msg.Name = "Lena"
		msg.Message = "hey"
		msg.Number = 123

		// Note that msg.Name becomes "user" in the JSON
		c.JSON(http.StatusOK, msg)
	})

	r.GET("/moreXML", func(c *gin.Context) {
		// You can also use a struct
		type MessageRecord struct {
			Name    string
			Message string
			Number  int
		}

		var msg MessageRecord
		msg.Name = "Lena"
		msg.Message = "hey"
		msg.Number = 123
		c.XML(http.StatusOK, msg)
	})

	// Listen and serve on 0.0.0.0:8080
	r.Run(":8080")
}
