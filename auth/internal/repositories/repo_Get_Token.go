package repositories

import (
	"github.com/duard/simples-microservices-go-node/auth/internal/model"
)

func (u *UserPostgresRepo) GetToken(refreshToken string) (model.User, error) {
	user := model.User{}
	u.db.Model(&user).Select("id", "token", "username").Where("token = ?", refreshToken).Scan(&user)
	return user, nil
}
