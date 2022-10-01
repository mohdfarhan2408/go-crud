package controllers

import (
	
	"github.com/gin/gin-gonic/gin"
	"Projects/go-crud/models"
	"Projects/go-crud/initializers"
)


func PostsCreate (c *gin.Context) {

	//Get data of req body


	//Create a post
	post := models.Post{Title: "Jinzhu", Body : "post"}

	result := initializers.DB.Create(&post)

	if result.Error != nil {
		c.Status(400)
		return
	}

	//Return it
	c.JSON(200, gin.H{
		"post": post,
	})
}