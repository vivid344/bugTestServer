package main

import (
	"github.com/gin-gonic/gin"
)

type Send_json struct {
	User     string `form:"name" json:"name"`
	Password string `form:"raw_password" json:"raw_password"`
}

func main(){
	r := gin.Default()

	r.POST("/", func(c *gin.Context) {
		var json Send_json
		if c.BindJSON(&json) == nil {
			c.JSON(200, gin.H{
				"name": json.User,
				"pass": json.Password,
			})
		}
	})
	r.Run(":8080")
}

// curlコマンド　curl -v http://localhost:8080/ -H "Content-Type: application/json" -X POST -d '{"name":"hoge","raw_password":"huga"}'