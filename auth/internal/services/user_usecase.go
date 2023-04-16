package services

import "github.com/duard/simples-microservices-go-node/internal/ports"

type userUseCase struct {
	userRepo ports.UserRepository
}

func NewUseCase(repo ports.UserRepository) ports.UserUseCase {
	return &userUseCase{userRepo: repo}
}
