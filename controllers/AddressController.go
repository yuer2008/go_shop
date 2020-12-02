package controllers

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"go-test.com/models"
)

func Add(c *gin.Context) {
	db, _ := models.InitDatabase()
	sql := "insert into person (first_name,last_name) value (?,?)"
	models.Execute(db, sql, "yang", "vincent")
	c.JSON(200, gin.H{
		"aaa": 120,
	})
}
func List(c *gin.Context) {

	db, _ := models.InitDatabase()
	sql := "select * from person"
	rows, _ := models.QueryData(db, sql)
	fmt.Printf("%v", rows)
	c.JSON(200, gin.H{
		"data": rows,
	})
}
