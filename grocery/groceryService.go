package grocery

import (
	"log"
	"net/http"

	"github.com/carlosm27/apiwithgorm/model"
	"github.com/gin-gonic/gin"
)

type NewGrocery struct {
	Name     string `json:"name" binding:"required"`
	Quantity int    `json:"quantity" binding:"required"`
}

type GroceryUpdate struct {
	Name     string `json:"name"`
	Quantity int    `json:"quantity"`
}

func GetGroceries(c *gin.Context) {

	var groceries []model.Grocery

	db, err := model.Database()
	if err != nil {
		log.Println(err)
	}

	if err := db.Find(&groceries).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, groceries)

}

func GetGrocery(c *gin.Context) {

	var grocery model.Grocery

	db, err := model.Database()
	if err != nil {
		log.Println(err)
	}

	if err := db.Where("id= ?", c.Param("id")).First(&grocery).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Grocery not found"})
		return
	}

	c.JSON(http.StatusOK, grocery)

}
