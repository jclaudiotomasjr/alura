package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/jclaudiotomasjr/alura/api-go-gin/controllers"
)

func HandleRequests() {
	r := gin.Default()
	r.GET("/users/", controllers.AllUsers)
	r.GET(":name", controllers.Welcome)
	r.Run()

}
