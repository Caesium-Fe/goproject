package webproject

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Gin() {
	r := gin.Default()
	r.GET("/ping/:name", func(c *gin.Context) {
		name := c.Param("name")
		c.String(http.StatusOK, "Hello %s", name)
		//c.JSON(200, gin.H{
		//	"message": "pong",
		//})
	})

	r.GET("/user/:name/:action", func(c *gin.Context) {
		name := c.Param("name")
		action := c.Param("action")
		message := name + " is " + action
		c.String(http.StatusOK, message)
	})

	r.POST("/user/:name/:action", func(c *gin.Context) {
		b := c.FullPath() == "/user/:name/:action" // true
		c.String(http.StatusOK, "%t", b)
	})

	r.GET("/user/groups", func(c *gin.Context) {
		c.String(http.StatusOK, "The available groups are [...]")
	})

	err := r.Run()
	if err != nil {
		return
	} // listen and serve on 127.0.0.1:8080
}

// a b c d e f g h i j k l m n o p q r s t u v w x y z
// A B C D E F G H I J K L M N O P Q R S T U V W X Y Z
