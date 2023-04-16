package middleware

import "github.com/duard/simples-microservices-go-node/internal/ports"

type UserMiddleware struct {
	userUseCase ports.UserUseCase
}

func NewUserMiddleware(u ports.UserUseCase) *UserMiddleware {
	handler := &UserMiddleware{userUseCase: u}
	return handler
}
