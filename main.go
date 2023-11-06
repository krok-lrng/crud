package main

import (
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"net/http"
)

var dsn = "root:@tcp(127.0.0.1:3306)/golang?charset=utf8mb4&parseTime=True&loc=Local"
var db, _ = gorm.Open(mysql.Open(dsn), &gorm.Config{})

type User struct {
	ID       int    `form:"id"`
	NAME     string `form:"name"`
	PASSWORD int    `form:"password"`
}

func all(c *gin.Context) {
	var users []User
	db.Find(&users)
	for _, user := range users {
		c.JSON(200, gin.H{
			"ID":       user.ID,
			"NAME":     user.NAME,
			"PASSWORD": user.PASSWORD,
		})
	}

}

func load(c *gin.Context) {
	c.HTML(http.StatusOK, "form.html", nil)
}

func create(c *gin.Context) {
	user := User{}
	c.Bind(&user)
	db.Create(&user)
}

func deleteData(c *gin.Context) {
	user := User{}
	c.Bind(&user)
	db.Delete(&user)
}

func main() {
	r := gin.Default()
	r.LoadHTMLGlob("./html/*")
	r.GET("/all", all)
	r.GET("/create", load)
	r.POST("/create", create)
	r.GET("/delete", load)
	r.POST("/delete", deleteData)
	r.Run() // listen and serve on 0.0.0.0:8080
}
