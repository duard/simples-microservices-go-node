package handler

import (
	"github.com/duard/simples-microservices-go-node/internal/model"
	"github.com/gin-gonic/gin"
)

func (u *UserHandler) Unfollow(c *gin.Context) {
	follow := model.User_followers{}
	follow.Follower = c.Value("username").(string)
	follow.Followee = c.Query("username")
	if follow.Followee == "" {
		c.JSON(400, gin.H{"message": "Please specify followee"})
		return
	}

	err := u.userUseCase.UnfollowUser(follow)
	if err != nil {
		c.JSON(400, gin.H{"message": err.Error()})
		return
	}
	c.JSON(201, gin.H{"message": "User Unfollowed"})
}
