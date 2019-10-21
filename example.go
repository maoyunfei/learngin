package main

import "github.com/gin-gonic/gin"

type Person struct {
	Id   int64  `uri:"id" binding:"required"`
	Name string `uri:"name" binding:"required"`
}

func main() {
	//gin.DisableConsoleColor()
	//gin.SetMode(gin.DebugMode)

	r := gin.Default()

	r.GET("/ping", func(context *gin.Context) {
		context.JSON(200, gin.H{"message": "pong"})
	})

	r.GET("/person/:name/:id", func(context *gin.Context) {
		var person Person
		if err := context.ShouldBindUri(&person); err != nil {
			context.JSON(400, gin.H{"msg": err})
			return
		}
		context.JSON(200, gin.H{"name": person.Name, "id": person.Id})
	})

	r.Run(":8086")
}
