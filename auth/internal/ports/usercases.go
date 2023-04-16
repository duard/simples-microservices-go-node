package ports

import (
	"github.com/duard/simples-microservices-go-node/auth/internal/model"
	"github.com/golang-jwt/jwt"
)

type UserUseCase interface {
	Create(user model.User) error
	Update(user model.User, prevUsername string) error
	Delete(username string) error

	Login(username string, password string) (string, string, error)
	Logout(id string) error
	AddToken(username, token string) error
	GetToken(refreshToken string) (model.User, error)
	GetTokenClaims(token string) (*jwt.Token, jwt.MapClaims, error)
	GenerateAccessToken(name, Id string) (string, error)

	GetUserbyUsername(id string) (*model.User, error)
	FollowUser(follow model.User_followers) error
	UnfollowUser(follow model.User_followers) error
}
