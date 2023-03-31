package routes

import (
	"github.com/LittleMikle/rest_user/pkg/controller"
	"github.com/gin-gonic/gin"
)

func UserRoute(router *gin.Engine) {
	router.GET("/users", controller.GetUsers)
	router.GET("/users/:id", controller.GetUserByID)
	router.POST("/users", controller.CreateUser)
	router.PUT("/users/:id", controller.UpdateUser)
	router.DELETE("/users/:id", controller.DeleteUser)
}
