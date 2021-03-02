package repository

import (
	"strings"

	"github.com/thospol/go-graphql/model"
	"gorm.io/gorm"
)

// UserRepository repository user interface
type UserRepository interface {
	Create(database *gorm.DB, i interface{}) error
	Update(database *gorm.DB, i interface{}) error
	Delete(database *gorm.DB, i interface{}) error
	FindByID(database *gorm.DB, id uint, i interface{}) error
	FindByName(database *gorm.DB, name string) ([]*model.User, error)
}

type userRepository struct {
	Repository
}

// NewUserRepository new user repository
func NewUserRepository() UserRepository {
	return &userRepository{
		NewRepository(),
	}
}

// FindByName find by name
func (rp *userRepository) FindByName(database *gorm.DB, name string) ([]*model.User, error) {
	users := []*model.User{}
	if err := database.
		Where("lower(name) = ?", strings.ToLower(name)).
		Find(&users).Error; err != nil {
		return nil, err
	}

	return users, nil
}
