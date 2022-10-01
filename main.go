package main

import (
	

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"Projects/go-crud/initializers"
	
)


func init() {

	initializers.LoadEnvVariables()
	initializers.ConnectToDB()

}

func main() {
	r := gin.Default()

	r.POST("/posts", controllers.PostsCreate)

	r.Run() 
}