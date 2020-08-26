package persistance

import (
	"api/domain/model"
	"api/usecase/repository"

	"fmt"

	"github.com/jinzhu/gorm"
)

/* Type UserRepository */
type userRepository struct {
	db *gorm.DB
}

// NewUserRepository repository datastore.
func NewUserRepository(db *gorm.DB) repository.IUserRepository {
	fmt.Println("REPOSITORY : NewUserRepository OK")
	return &userRepository{db}
}

func (userRepo *userRepository) FindByID(id string) (*model.User, error) {
	user := model.User{UserId: id}
	err := userRepo.db.First(&user).Error

	// First ERROR
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (userRepo *userRepository) Store(user *model.User) error {
	return userRepo.db.Save(user).Error
}

func (userRepo *userRepository) Update(user *model.User) error {
	// 値を更新していなくても、SQLの更新を実行すると、すべてのフィールドが上書きされます。
	return userRepo.db.Model(&model.User{UserId: user.UserId}).Updates(user).Error
}

func (userRepo *userRepository) Delete(user *model.User) error {
	return userRepo.db.Delete(user).Error
}

func (userRepo *userRepository) FindAll() ([]*model.User, error) {
	users := []*model.User{}

	err := userRepo.db.Find(&users).Error

	// Find ERROR
	if err != nil {
		return nil, err
	}

	return users, nil
}
