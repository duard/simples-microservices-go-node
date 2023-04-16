package services

import "github.com/duard/simples-microservices-go-node/auth/internal/model"

func (u *userUseCase) FollowUser(follow model.User_followers) error {
	err := u.userRepo.FollowUser(follow)
	return err
}
