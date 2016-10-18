package main

import (
	"github.com/gin-gonic/gin"
	"github.com/nabetama/Kiz/controllers"
	"reflect"
	"strconv"
)

const SESSION_NAME string = "Kiz"

func main() {
	router := gin.Default()
	router.GET("/:id", func(c *gin.Context) {
		// Paramを処理する
		n := c.Param("id")
		id, err := strconv.Atoi(n)
		if err != nil {
			c.JSON(400, err)
			return
		}
		if id <= 0 {
			c.JSON(400, gin.H{"status": "id should be bigger than 0"})
			return
		}
		// データを処理する
		ctrl := controllers.NewUser()
		result := ctrl.Get(id)
		if result == nil || reflect.ValueOf(result).IsNil() {
			c.JSON(404, gin.H{})
			return
		}

		c.JSON(200, result)
	})
	router.Run(":8080")
}
