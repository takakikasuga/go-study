package repository

import (
	"go-rest-api/model"

	"gorm.io/gorm"
)

// NOTE: usecaseが依存するメソッドを定義
type IUserRepository interface {
	GetUserByEmail(user *model.User, email string) error
	CreateUser(user *model.User) error
}

// NOTE: DI（依存注入）用の構造体
type userRepository struct {
	db *gorm.DB
}

// NOTE: コンストラクター関数（外部でインスタンス化される依存関係を注入）
func NewUserRepository(db *gorm.DB) IUserRepository {
	return &userRepository{db: db}
}

func (ur *userRepository) GetUserByEmail(user *model.User, email string) error {
	if err := ur.db.Where("email = ?", email).First(user).Error; err != nil {
		return err
	}
	return nil
}

func (ur *userRepository) CreateUser(user *model.User) error {
	if err := ur.db.Create(user).Error; err != nil {
		return err
	}
	return nil
}
