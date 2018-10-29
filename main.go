package main

import "github.com/gin-gonic/gin"

func main(){
	r := gin.Default()

	r.POST("/", func(c *gin.Context) {
		name := c.PostForm("name")
		pass := c.PostForm("raw_password")

		c.JSON(200, gin.H{
			"name": name,
			"pass": pass,
		})
	})

	r.Run(":8080")
}

// curlコマンド　curl -X POST http://localhost:8080 -d "name=hoge&raw_password=huga"