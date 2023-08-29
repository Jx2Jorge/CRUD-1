package main

import (
	"github.com/gin-gonic/gin"
	"githug.com/Jx2Jorge/CRUD-1/models"
)

func main() {
	r := gin.Default()

	models.ConnectDatabase() // new

	r.Run()
}
