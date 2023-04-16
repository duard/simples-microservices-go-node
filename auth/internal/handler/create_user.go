package handler

import (
	"github.com/duard/simples-microservices-go-node/internal/model"
	"github.com/gin-gonic/gin"
)

func (u *UserHandler) CreateUser(c *gin.Context) {
	user := model.User{}
	err := c.BindJSON(&user)
	if err != nil {
		c.JSON(400, gin.H{"message": "something went wrong"})
		return
	}

	err = u.userUseCase.Create(user)
	if err != nil {
		c.JSON(400, gin.H{"message": "Please check your values"})
		return
	}
	c.JSON(201, gin.H{"message": "User Created Successfully"})
}
