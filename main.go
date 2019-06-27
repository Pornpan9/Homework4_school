package main

import (
	"github.com/Pornpan9/Homework4_school/todo"
	"github.com/Pornpan9/Homework4_school/database"
	"github.com/gin-gonic/gin"
)

func main() {

	db := database.database{}
	db.InitialDB()
	
router := gin.Default()
api := router.Group("/api")
api.GET("/todos", todo.GetHandler)
router.Run(":1234")

}
