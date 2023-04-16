package services

import "github.com/duard/simples-microservices-go-node/auth/internal/model"

func (u *userUseCase) UnfollowUser(follow model.User_followers) error {
	err := u.userRepo.UnfollowUser(follow)
	if err != nil {
		return err
	}
	return nil
}
