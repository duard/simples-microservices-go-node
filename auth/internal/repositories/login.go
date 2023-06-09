package repositories

import "github.com/duard/simples-microservices-go-node/auth/internal/model"

func (u *UserPostgresRepo) Login(username string) (*model.User, error) {
	user := model.User{}
	result := u.db.Model(&user).Select("id", "username", "password", "salt").Where("username = ?", username).Scan(&user)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}
