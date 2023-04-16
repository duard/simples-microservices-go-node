package services

import (
	"time"

	"github.com/duard/simples-microservices-go-node/auth/internal/model"
	"github.com/duard/simples-microservices-go-node/internal/utils"
)

func (u *userUseCase) Create(user model.User) error {

	HashedPassword, salt, err := utils.GenerateHashfromPassword(user.Password)
	if err != nil {
		return err
	}
	user.Password = HashedPassword
	user.Salt = salt
	user.Id = utils.GenerateID().String()
	user.Date_created = time.Now()
	// fmt.Println(user.Date_created)
	err = u.userRepo.Create(user)
	if err != nil {
		return err
	}
	return nil
}
