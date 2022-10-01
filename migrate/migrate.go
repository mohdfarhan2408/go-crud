package main

import (
	
	"Projects/go-crud/initializers"
	"Projects/go-crud/models"
	
)


func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main(){
	initializers.DB.AutoMigrate(&models.Post{})
}