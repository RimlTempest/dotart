package repository

import (
	"api/domain/model"
)

// 処理のインターフェイス
type IUserRepository interface {
	FindByID(id string) (*model.User, error)
	Store(user *model.User) error
	Update(user *model.User) error
	Delete(user *model.User) error
	FindAll() ([]*model.User, error)
}
