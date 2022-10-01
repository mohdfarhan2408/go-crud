package controllers

import (
	
	"github.com/gin/gin-gonic/gin"
	"Projects/go-crud/models"
	"Projects/go-crud/initializers"
)


func PostsCreate (c *gin.Context) {

	//Get data of req body
	var body struct {
		Body string
		Title string
	}

	c.Bind(&body)

	//Create a post
	post := models.Post{Title: body.Title, Body : body.Body}

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

func PostsIndex (c *gin.Context) {

	//Get the posts
	var posts []models.Post //array
	initializers.DB.Find(&posts)

	//Respond with them
	c.JSON(200, gin.H{
		"posts": posts,
	})
}

func PostsShow (c *gin.Context) {
	//Get id of url
	id := c.Param("id")

	//Get the post
	var post models.Post
	initializers.DB.First(&post, id)

	//Respond with them
	c.JSON(200, gin.H{
		"post": post,
	})
}

func PostsUpdate (c *gin.Context) {
	//Get the id of the url
	id := c.Param("id")

	//Get the data off req body
	var body struct {
		Body string
		Title string
	}

	c.Bind(&body)

	//Find the post were updating
	var post models.Post
	initializers.DB.First(&post, id)


	//Update it
	initializers.DB.Model(&post).Updates(models.Post{
		Title: body.Title,
		Body: body.Body,
	})

	//Respond it
	c.JSON(200, gin.H{
		"post": post,
	})

}

func PostsDelete (c *gin.Context) {
	//Get the id  of the url
	id := c.Param("id")

	//Delete the data in db
	initializers.DB.Delete(&models.Post{}, id)

	//Respond
	c.Status(200, gin.H{
		"status" : "success",
	})

}
