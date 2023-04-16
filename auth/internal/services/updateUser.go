package services

import "github.com/duard/simples-microservices-go-node/auth/internal/model"

func (u *userUseCase) Update(user model.User, prevUsername string) error {
	err := u.userRepo.Update(user, prevUsername)
	if err != nil {
		return err
	}
	return nil
}
