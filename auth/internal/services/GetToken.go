package services

import "github.com/duard/simples-microservices-go-node/auth/internal/model"

func (u *userUseCase) GetToken(refreshToken string) (model.User, error) {
	user, err := u.userRepo.GetToken(refreshToken)

	if err != nil {
		return user, err
	}
	return user, nil
}
