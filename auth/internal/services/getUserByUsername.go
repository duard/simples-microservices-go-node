package services

import (
	"github.com/duard/simples-microservices-go-node/auth/internal/model"
)

func (u *userUseCase) GetUserbyUsername(username string) (*model.User, error) {
	user, err := u.userRepo.GetUserbyUsername(username)
	if err != nil {
		return nil, err
	}
	return user, nil
}
