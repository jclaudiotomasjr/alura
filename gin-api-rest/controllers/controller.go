package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jclaudiotomasjr/alura/api-go-gin/database"
	"github.com/jclaudiotomasjr/alura/api-go-gin/models"
)

func AllUsers(c *gin.Context) {
	var users []models.User
	database.DB.Find(&users)
	c.JSON(200, users)
}

func Welcome(c *gin.Context) {
	name := c.Params.ByName("name")
	c.JSON(200, gin.H{
		"APP": "Welcome " + name + " Tudo Tranquilo?",
	})
}

func CreateUser(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"erro": err.Error()})
	}
	database.DB.Create(&user)
	c.JSON(http.StatusOK, gin.H{
		"status": "Usuário criado com sucesso!"})
}

func ReturnUser(c *gin.Context) {
	var user models.User
	id := c.Params.ByName("id")
	database.DB.First(&user, id)
	if user.ID == 0 {
		c.JSON(http.StatusFound, gin.H{
			"Status": "Usuário não encontrado!"})
		return
	}
	c.JSON(http.StatusOK, user)
}

func DeleteUser(c *gin.Context) {
	var user models.User
	id := c.Params.ByName("id")
	database.DB.First(&user, id)
	if user.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"status": "Usuário não foi deletado, pois não foi encontrado!"})
		return
	}
	database.DB.Delete(&user, id)
	c.JSON(http.StatusOK, user)

}

func EditUser(c *gin.Context) {
	var user models.User
	id := c.Params.ByName("id")
	database.DB.First(&user, id)
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
		return
	}
	database.DB.Model(&user).UpdateColumns(user)
	c.JSON(http.StatusOK, gin.H{
		"status": "Usuário editado com sucesso!"})
	c.JSON(http.StatusOK, user)
}
