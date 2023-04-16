package repositories

import "github.com/duard/simples-microservices-go-node/auth/internal/model"

func (u *UserPostgresRepo) AddToken(username, token string) error {
	user := model.User{}
	user.Token = token
	result := u.db.Model(&user).Where("username = ?", username).Update("token", user.Token)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
