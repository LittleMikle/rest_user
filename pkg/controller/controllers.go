package controller

import (
	"github.com/LittleMikle/rest_user/pkg/models"
	"github.com/LittleMikle/rest_user/storage"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetUsers(c *gin.Context) {
	users := []models.User{}
	storage.DB.Find(&users)
	c.JSON(http.StatusOK, &users)
}

func GetUserByID(c *gin.Context) {
	var user models.User
	storage.DB.Where("id = ?", c.Param("id")).First(&user)
	storage.DB.Find(&user)
	c.JSON(http.StatusOK, &user)
}

func CreateUser(c *gin.Context) {
	var user models.User
	c.BindJSON(&user)
	storage.DB.Create(&user)
	c.JSON(http.StatusOK, &user)
}

func UpdateUser(c *gin.Context) {
	var user models.User
	storage.DB.Where("id = ?", c.Param("id")).First(&user)
	c.BindJSON(&user)
	storage.DB.Save(&user)
	c.JSON(http.StatusOK, &user)
}

func DeleteUser(c *gin.Context) {
	var user models.User
	storage.DB.Where("id = ?", c.Param("id")).Delete(&user)
	c.JSON(http.StatusOK, &user)
}
